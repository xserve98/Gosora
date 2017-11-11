package common

import (
	"database/sql"
	"sync/atomic"

	"../query_gen/lib"
)

type WordFilter struct {
	ID          int
	Find        string
	Replacement string
}
type WordFilterMap map[int]WordFilter

var WordFilterBox atomic.Value // An atomic value holding a WordFilterBox

type FilterStmts struct {
	getWordFilters *sql.Stmt
}

var filterStmts FilterStmts

func init() {
	WordFilterBox.Store(WordFilterMap(make(map[int]WordFilter)))
	DbInits.Add(func() error {
		acc := qgen.Builder.Accumulator()
		filterStmts = FilterStmts{
			getWordFilters: acc.Select("word_filters").Columns("wfid, find, replacement").Prepare(),
		}
		return acc.FirstError()
	})
}

func LoadWordFilters() error {
	rows, err := filterStmts.getWordFilters.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	var wordFilters = WordFilterMap(make(map[int]WordFilter))
	var wfid int
	var find string
	var replacement string

	for rows.Next() {
		err := rows.Scan(&wfid, &find, &replacement)
		if err != nil {
			return err
		}
		wordFilters[wfid] = WordFilter{ID: wfid, Find: find, Replacement: replacement}
	}
	WordFilterBox.Store(wordFilters)
	return rows.Err()
}

func AddWordFilter(id int, find string, replacement string) {
	wordFilters := WordFilterBox.Load().(WordFilterMap)
	wordFilters[id] = WordFilter{ID: id, Find: find, Replacement: replacement}
	WordFilterBox.Store(wordFilters)
}