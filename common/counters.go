package common

import (
	"database/sql"
	"sync"
	"sync/atomic"

	"../query_gen/lib"
)

var GlobalViewCounter *ChunkedViewCounter
var RouteViewCounter *DefaultRouteViewCounter
var TopicViewCounter *DefaultTopicViewCounter

type ChunkedViewCounter struct {
	buckets       [2]int64
	currentBucket int64

	insert *sql.Stmt
}

func NewChunkedViewCounter() (*ChunkedViewCounter, error) {
	acc := qgen.Builder.Accumulator()
	counter := &ChunkedViewCounter{
		currentBucket: 0,
		insert:        acc.Insert("viewchunks").Columns("count, createdAt").Fields("?,UTC_TIMESTAMP()").Prepare(),
	}
	AddScheduledFifteenMinuteTask(counter.Tick) // This is run once every fifteen minutes to match the frequency of the RouteViewCounter
	//AddScheduledSecondTask(counter.Tick)
	AddShutdownTask(counter.Tick)
	return counter, acc.FirstError()
}

func (counter *ChunkedViewCounter) Tick() (err error) {
	var oldBucket = counter.currentBucket
	var nextBucket int64 // 0
	if counter.currentBucket == 0 {
		nextBucket = 1
	}
	atomic.AddInt64(&counter.buckets[oldBucket], counter.buckets[nextBucket])
	atomic.StoreInt64(&counter.buckets[nextBucket], 0)
	atomic.StoreInt64(&counter.currentBucket, nextBucket)

	var previousViewChunk = counter.buckets[oldBucket]
	atomic.AddInt64(&counter.buckets[oldBucket], -previousViewChunk)
	return counter.insertChunk(previousViewChunk)
}

func (counter *ChunkedViewCounter) Bump() {
	atomic.AddInt64(&counter.buckets[counter.currentBucket], 1)
}

func (counter *ChunkedViewCounter) insertChunk(count int64) error {
	if count == 0 {
		return nil
	}
	debugLogf("Inserting a viewchunk with a count of %d", count)
	_, err := counter.insert.Exec(count)
	return err
}

type RWMutexCounterBucket struct {
	counter int
	sync.RWMutex
}

// The name of the struct clashes with the name of the variable, so we're adding Impl to the end
type DefaultRouteViewCounter struct {
	routeBuckets []*RWMutexCounterBucket //[RouteID]count
	insert       *sql.Stmt
}

func NewDefaultRouteViewCounter() (*DefaultRouteViewCounter, error) {
	acc := qgen.Builder.Accumulator()
	var routeBuckets = make([]*RWMutexCounterBucket, len(routeMapEnum))
	for bucketID, _ := range routeBuckets {
		routeBuckets[bucketID] = &RWMutexCounterBucket{counter: 0}
	}
	counter := &DefaultRouteViewCounter{
		routeBuckets: routeBuckets,
		insert:       acc.Insert("viewchunks").Columns("count, createdAt, route").Fields("?,UTC_TIMESTAMP(),?").Prepare(),
	}
	AddScheduledFifteenMinuteTask(counter.Tick) // There could be a lot of routes, so we don't want to be running this every second
	//AddScheduledSecondTask(counter.Tick)
	AddShutdownTask(counter.Tick)
	return counter, acc.FirstError()
}

func (counter *DefaultRouteViewCounter) Tick() error {
	for routeID, routeBucket := range counter.routeBuckets {
		var count int
		routeBucket.RLock()
		count = routeBucket.counter
		routeBucket.counter = 0
		routeBucket.RUnlock()

		err := counter.insertChunk(count, routeID) // TODO: Bulk insert for speed?
		if err != nil {
			return err
		}
	}
	return nil
}

func (counter *DefaultRouteViewCounter) insertChunk(count int, route int) error {
	if count == 0 {
		return nil
	}
	var routeName = reverseRouteMapEnum[route]
	debugLogf("Inserting a viewchunk with a count of %d for route %s (%d)", count, routeName, route)
	_, err := counter.insert.Exec(count, routeName)
	return err
}

func (counter *DefaultRouteViewCounter) Bump(route int) {
	// TODO: Test this check
	debugLog("counter.routeBuckets[", route, "]: ", counter.routeBuckets[route])
	if len(counter.routeBuckets) <= route {
		return
	}
	counter.routeBuckets[route].Lock()
	counter.routeBuckets[route].counter++
	counter.routeBuckets[route].Unlock()
}

// TODO: The ForumViewCounter and TopicViewCounter

// TODO: Unload forum counters without any views over the past 15 minutes, if the admin has configured the forumstore with a cap and it's been hit?
// Forums can be reloaded from the database at any time, so we want to keep the counters separate from them
type ForumViewCounter struct {
	buckets       [2]int64
	currentBucket int64
}

/*func (counter *ForumViewCounter) insertChunk(count int, forum int) error {
	if count == 0 {
		return nil
	}
	debugLogf("Inserting a viewchunk with a count of %d for forum %d", count, forum)
	_, err := counter.insert.Exec(count, forum)
	return err
}*/

// TODO: Use two odd-even maps for now, and move to something more concurrent later, maybe a sharded map?
type DefaultTopicViewCounter struct {
	oddTopics  map[int]*RWMutexCounterBucket // map[tid]struct{counter,sync.RWMutex}
	evenTopics map[int]*RWMutexCounterBucket
	oddLock    sync.RWMutex
	evenLock   sync.RWMutex

	update *sql.Stmt
}

func NewDefaultTopicViewCounter() (*DefaultTopicViewCounter, error) {
	acc := qgen.Builder.Accumulator()
	counter := &DefaultTopicViewCounter{
		oddTopics:  make(map[int]*RWMutexCounterBucket),
		evenTopics: make(map[int]*RWMutexCounterBucket),
		update:     acc.Update("topics").Set("views = views + ?").Where("tid = ?").Prepare(),
	}
	AddScheduledFifteenMinuteTask(counter.Tick) // Who knows how many topics we have queued up, we probably don't want this running too frequently
	//AddScheduledSecondTask(counter.Tick)
	AddShutdownTask(counter.Tick)
	return counter, acc.FirstError()
}

func (counter *DefaultTopicViewCounter) Tick() error {
	counter.oddLock.RLock()
	oddTopics := counter.oddTopics
	counter.oddLock.RUnlock()
	for topicID, topic := range oddTopics {
		var count int
		topic.RLock()
		count = topic.counter
		topic.RUnlock()
		// TODO: Only delete the bucket when it's zero to avoid hitting popular topics?
		counter.oddLock.Lock()
		delete(counter.oddTopics, topicID)
		counter.oddLock.Unlock()
		err := counter.insertChunk(count, topicID)
		if err != nil {
			return err
		}
	}

	counter.evenLock.RLock()
	evenTopics := counter.evenTopics
	counter.evenLock.RUnlock()
	for topicID, topic := range evenTopics {
		var count int
		topic.RLock()
		count = topic.counter
		topic.RUnlock()
		// TODO: Only delete the bucket when it's zero to avoid hitting popular topics?
		counter.evenLock.Lock()
		delete(counter.evenTopics, topicID)
		counter.evenLock.Unlock()
		err := counter.insertChunk(count, topicID)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: Optimise this further. E.g. Using IN() on every one view topic. Rinse and repeat for two views, three views, four views and five views.
func (counter *DefaultTopicViewCounter) insertChunk(count int, topicID int) error {
	if count == 0 {
		return nil
	}
	debugLogf("Inserting %d views into topic %d", count, topicID)
	_, err := counter.update.Exec(count, topicID)
	return err
}

func (counter *DefaultTopicViewCounter) Bump(topicID int) {
	// Is the ID even?
	if topicID%2 == 0 {
		counter.evenLock.RLock()
		topic, ok := counter.evenTopics[topicID]
		counter.evenLock.RUnlock()
		if ok {
			topic.Lock()
			topic.counter++
			topic.Unlock()
		} else {
			counter.evenLock.Lock()
			counter.evenTopics[topicID] = &RWMutexCounterBucket{counter: 1}
			counter.evenLock.Unlock()
		}
		return
	}

	counter.oddLock.RLock()
	topic, ok := counter.oddTopics[topicID]
	counter.oddLock.RUnlock()
	if ok {
		topic.Lock()
		topic.counter++
		topic.Unlock()
	} else {
		counter.oddLock.Lock()
		counter.oddTopics[topicID] = &RWMutexCounterBucket{counter: 1}
		counter.oddLock.Unlock()
	}
}
