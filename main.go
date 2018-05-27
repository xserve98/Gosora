/*
*
*	Gosora Main File
*	Copyright Azareal 2016 - 2018
*
 */
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"./common"
	"./common/counters"
	"./config"
	"./query_gen/lib"
	"./routes"
	"github.com/fsnotify/fsnotify"
)

var version = common.Version{Major: 0, Minor: 1, Patch: 0, Tag: "dev"}
var router *GenRouter
var logWriter = io.MultiWriter(os.Stderr)

// TODO: Wrap the globals in here so we can pass pointers to them to subpackages
var globs *Globs

type Globs struct {
	stmts *Stmts
}

func afterDBInit() (err error) {
	acc := qgen.Builder.Accumulator()
	common.Rstore, err = common.NewSQLReplyStore(acc)
	if err != nil {
		return err
	}
	common.Prstore, err = common.NewSQLProfileReplyStore(acc)
	if err != nil {
		return err
	}

	err = common.InitTemplates()
	if err != nil {
		return err
	}
	err = common.InitPhrases()
	if err != nil {
		return err
	}

	log.Print("Loading the static files.")
	err = common.Themes.LoadStaticFiles()
	if err != nil {
		return err
	}
	err = common.StaticFiles.Init()
	if err != nil {
		return err
	}
	err = common.StaticFiles.JSTmplInit()
	if err != nil {
		return err
	}

	log.Print("Initialising the widgets")
	err = common.InitWidgets()
	if err != nil {
		return err
	}

	log.Print("Initialising the menu item list")
	common.Menus = common.NewDefaultMenuStore()
	err = common.Menus.Load(1) // 1 = the default menu
	if err != nil {
		return err
	}
	menuHold, err := common.Menus.Get(1)
	if err != nil {
		return err
	}
	fmt.Printf("menuHold: %+v\n", menuHold)
	var b bytes.Buffer
	menuHold.Build(&b, &common.GuestUser)
	fmt.Println("menuHold output: ", string(b.Bytes()))

	log.Print("Initialising the authentication system")
	common.Auth, err = common.NewDefaultAuth()
	if err != nil {
		return err
	}

	log.Print("Loading the word filters")
	err = common.LoadWordFilters()
	if err != nil {
		return err
	}

	log.Print("Initialising the stores")
	common.Reports, err = common.NewDefaultReportStore(acc)
	if err != nil {
		return err
	}
	common.Emails, err = common.NewDefaultEmailStore(acc)
	if err != nil {
		return err
	}
	common.RegLogs, err = common.NewRegLogStore(acc)
	if err != nil {
		return err
	}
	common.ModLogs, err = common.NewModLogStore(acc)
	if err != nil {
		return err
	}
	common.AdminLogs, err = common.NewAdminLogStore(acc)
	if err != nil {
		return err
	}
	common.IPSearch, err = common.NewDefaultIPSearcher()
	if err != nil {
		return err
	}
	common.Subscriptions, err = common.NewDefaultSubscriptionStore()
	if err != nil {
		return err
	}
	common.Attachments, err = common.NewDefaultAttachmentStore()
	if err != nil {
		return err
	}
	common.Polls, err = common.NewDefaultPollStore(common.NewMemoryPollCache(100)) // TODO: Max number of polls held in cache, make this a config item
	if err != nil {
		return err
	}
	common.TopicList, err = common.NewDefaultTopicList()
	if err != nil {
		return err
	}

	log.Print("Initialising the view counters")
	counters.GlobalViewCounter, err = counters.NewGlobalViewCounter(acc)
	if err != nil {
		return err
	}
	counters.AgentViewCounter, err = counters.NewDefaultAgentViewCounter()
	if err != nil {
		return err
	}
	counters.OSViewCounter, err = counters.NewDefaultOSViewCounter()
	if err != nil {
		return err
	}
	counters.LangViewCounter, err = counters.NewDefaultLangViewCounter()
	if err != nil {
		return err
	}
	counters.RouteViewCounter, err = counters.NewDefaultRouteViewCounter()
	if err != nil {
		return err
	}
	counters.PostCounter, err = counters.NewPostCounter()
	if err != nil {
		return err
	}
	counters.TopicCounter, err = counters.NewTopicCounter()
	if err != nil {
		return err
	}
	counters.TopicViewCounter, err = counters.NewDefaultTopicViewCounter()
	if err != nil {
		return err
	}
	counters.ForumViewCounter, err = counters.NewDefaultForumViewCounter()
	if err != nil {
		return err
	}
	counters.ReferrerTracker, err = counters.NewDefaultReferrerTracker()
	if err != nil {
		return err
	}

	return nil
}

// TODO: Split this function up
func main() {
	// TODO: Recover from panics
	/*defer func() {
		r := recover()
		if r != nil {
			log.Print(r)
			debug.PrintStack()
			return
		}
	}()*/
	config.Config()

	// TODO: Have a file for each run with the time/date the server started as the file name?
	// TODO: Log panics with recover()
	f, err := os.OpenFile("./logs/ops.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	logWriter = io.MultiWriter(os.Stderr, f)
	log.SetOutput(logWriter)
	log.Print("Running Gosora v" + version.String())
	fmt.Println("")
	common.StartTime = time.Now()

	log.Print("Processing configuration data")
	err = common.ProcessConfig()
	if err != nil {
		log.Fatal(err)
	}

	common.Themes, err = common.NewThemeList()
	if err != nil {
		log.Fatal(err)
	}

	err = InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	buildTemplates := flag.Bool("build-templates", false, "build the templates")
	flag.Parse()
	if *buildTemplates {
		err = common.CompileTemplates()
		if err != nil {
			log.Fatal(err)
		}
		err = common.CompileJSTemplates()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = afterDBInit()
	if err != nil {
		log.Fatal(err)
	}

	err = common.VerifyConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Initialising the file watcher")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		var modifiedFileEvent = func(path string) error {
			var pathBits = strings.Split(path, "\\")
			if len(pathBits) == 0 {
				return nil
			}
			if pathBits[0] == "themes" {
				var themeName string
				if len(pathBits) >= 2 {
					themeName = pathBits[1]
				}
				if len(pathBits) >= 3 && pathBits[2] == "public" {
					// TODO: Handle new themes freshly plopped into the folder?
					theme, ok := common.Themes[themeName]
					if ok {
						return theme.LoadStaticFiles()
					}
				}
			}
			return nil
		}

		var err error
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				// TODO: Handle file deletes (and renames more graciously by removing the old version of it)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					err = modifiedFileEvent(event.Name)
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("new file:", event.Name)
					err = modifiedFileEvent(event.Name)
				}
				if err != nil {
					common.LogError(err)
				}
			case err = <-watcher.Errors:
				common.LogError(err)
			}
		}
	}()

	// TODO: Keep tabs on the (non-resource) theme stuff, and the langpacks
	err = watcher.Add("./public")
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add("./templates")
	if err != nil {
		log.Fatal(err)
	}
	for _, theme := range common.Themes {
		err = watcher.Add("./themes/" + theme.Name + "/public")
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("Initialising the task system")
	var runTasks = func(tasks []func() error) {
		for _, task := range tasks {
			if task() != nil {
				common.LogError(err)
			}
		}
	}

	// Run this goroutine once every half second
	halfSecondTicker := time.NewTicker(time.Second / 2)
	secondTicker := time.NewTicker(time.Second)
	fifteenMinuteTicker := time.NewTicker(15 * time.Minute)
	//hourTicker := time.NewTicker(time.Hour)
	go func() {
		var runHook = func(name string) {
			err := common.RunTaskHook(name)
			if err != nil {
				common.LogError(err)
			}
		}
		for {
			select {
			case <-halfSecondTicker.C:
				runHook("before_half_second_tick")
				runTasks(common.ScheduledHalfSecondTasks)
				runHook("after_half_second_tick")
			case <-secondTicker.C:
				runHook("before_second_tick")
				runTasks(common.ScheduledSecondTasks)

				// TODO: Stop hard-coding this
				err := common.HandleExpiredScheduledGroups()
				if err != nil {
					common.LogError(err)
				}

				// TODO: Handle delayed moderation tasks

				// Sync with the database, if there are any changes
				err = common.HandleServerSync()
				if err != nil {
					common.LogError(err)
				}

				// TODO: Manage the TopicStore, UserStore, and ForumStore
				// TODO: Alert the admin, if CPU usage, RAM usage, or the number of posts in the past second are too high
				// TODO: Clean-up alerts with no unread matches which are over two weeks old. Move this to a 24 hour task?
				// TODO: Rescan the static files for changes
				runHook("after_second_tick")
			case <-fifteenMinuteTicker.C:
				runHook("before_fifteen_minute_tick")
				runTasks(common.ScheduledFifteenMinuteTasks)

				// TODO: Automatically lock topics, if they're really old, and the associated setting is enabled.
				// TODO: Publish scheduled posts.
				runHook("after_fifteen_minute_tick")
			}

			// TODO: Handle the daily clean-up.
		}
	}()

	log.Print("Initialising the router")
	router, err = NewGenRouter(http.FileServer(http.Dir("./uploads")))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Initialising the plugins")
	common.InitPlugins()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		// TODO: Gracefully shutdown the HTTP server
		runTasks(common.ShutdownTasks)
		log.Fatal("Received a signal to shutdown: ", sig)
	}()

	//if profiling {
	//	pprof.StopCPUProfile()
	//}

	// We might not need the timeouts, if we're behind a reverse-proxy like Nginx
	var newServer = func(addr string, handler http.Handler) *http.Server {
		return &http.Server{
			Addr:    addr,
			Handler: handler,

			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		}
	}

	// TODO: Let users run *both* HTTP and HTTPS
	log.Print("Initialising the HTTP server")
	if !common.Site.EnableSsl {
		if common.Site.Port == "" {
			common.Site.Port = "80"
		}
		log.Print("Listening on port " + common.Site.Port)
		err = newServer(":"+common.Site.Port, router).ListenAndServe()
	} else {
		if common.Site.Port == "" {
			common.Site.Port = "443"
		}
		if common.Site.Port == "80" || common.Site.Port == "443" {
			// We should also run the server on port 80
			// TODO: Redirect to port 443
			go func() {
				log.Print("Listening on port 80")
				err = newServer(":80", &routes.HTTPSRedirect{}).ListenAndServe()
				if err != nil {
					log.Fatal(err)
				}
			}()
		}
		log.Printf("Listening on port %s", common.Site.Port)
		err = newServer(":"+common.Site.Port, router).ListenAndServeTLS(common.Config.SslFullchain, common.Config.SslPrivkey)
	}

	// Why did the server stop?
	if err != nil {
		log.Fatal(err)
	}
}
