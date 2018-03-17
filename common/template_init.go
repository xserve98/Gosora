package common

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"./templates"
)

var Ctemplates []string
var Templates = template.New("")
var PrebuildTmplList []func(User, *HeaderVars) CTmpl

type CTmpl struct {
	Name       string
	Filename   string
	Path       string
	StructName string
	Data       interface{}
	Imports    []string
}

// TODO: Stop duplicating these bits of code
// nolint
func interpreted_topic_template(pi TopicPage, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["topic"]
	if !ok {
		mapping = "topic"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// nolint
var Template_topic_handle = interpreted_topic_template
var Template_topic_alt_handle = interpreted_topic_template

// nolint
var Template_topics_handle = func(pi TopicsPage, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["topics"]
	if !ok {
		mapping = "topics"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// nolint
var Template_forum_handle = func(pi ForumPage, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["forum"]
	if !ok {
		mapping = "forum"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// nolint
var Template_forums_handle = func(pi ForumsPage, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["forums"]
	if !ok {
		mapping = "forums"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// nolint
var Template_profile_handle = func(pi ProfilePage, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["profile"]
	if !ok {
		mapping = "profile"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// nolint
var Template_create_topic_handle = func(pi CreateTopicPage, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["create_topic"]
	if !ok {
		mapping = "create_topic"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// nolint
var Template_login_handle = func(pi Page, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["login"]
	if !ok {
		mapping = "login"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// nolint
var Template_register_handle = func(pi Page, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["register"]
	if !ok {
		mapping = "register"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// nolint
var Template_error_handle = func(pi Page, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["error"]
	if !ok {
		mapping = "error"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// nolint
var Template_ip_search_handle = func(pi IPSearchPage, w http.ResponseWriter) error {
	mapping, ok := Themes[DefaultThemeBox.Load().(string)].TemplatesMap["ip_search"]
	if !ok {
		mapping = "ip_search"
	}
	return Templates.ExecuteTemplate(w, mapping+".html", pi)
}

// ? - Add template hooks?
func compileTemplates() error {
	var config tmpl.CTemplateConfig
	config.Minify = Config.MinifyTemplates
	config.SuperDebug = Dev.TemplateDebug

	var c tmpl.CTemplateSet
	c.SetConfig(config)

	// Schemas to train the template compiler on what to expect
	// TODO: Add support for interface{}s
	user := User{62, BuildProfileURL("fake-user", 62), "Fake User", "compiler@localhost", 0, false, false, false, false, false, false, GuestPerms, make(map[string]bool), "", false, BuildAvatar(62, ""), "", "", "", "", 0, 0, "0.0.0.0.0", 0}
	// TODO: Do a more accurate level calculation for this?
	user2 := User{1, BuildProfileURL("admin-alice", 1), "Admin Alice", "alice@localhost", 1, true, true, true, true, false, false, AllPerms, make(map[string]bool), "", true, BuildAvatar(1, ""), "", "", "", "", 58, 1000, "127.0.0.1", 0}
	user3 := User{2, BuildProfileURL("admin-fred", 62), "Admin Fred", "fred@localhost", 1, true, true, true, true, false, false, AllPerms, make(map[string]bool), "", true, BuildAvatar(2, ""), "", "", "", "", 42, 900, "::1", 0}
	headerVars := &HeaderVars{
		Site:        Site,
		Settings:    SettingBox.Load().(SettingMap),
		Themes:      Themes,
		Theme:       Themes[DefaultThemeBox.Load().(string)],
		NoticeList:  []string{"test"},
		Stylesheets: []string{"panel"},
		Scripts:     []string{"whatever"},
		Widgets: PageWidgets{
			LeftSidebar: template.HTML("lalala"),
		},
	}

	log.Print("Compiling the templates")

	var now = time.Now()
	poll := Poll{ID: 1, Type: 0, Options: map[int]string{0: "Nothing", 1: "Something"}, Results: map[int]int{0: 5, 1: 2}, QuickOptions: []PollOption{
		PollOption{0, "Nothing"},
		PollOption{1, "Something"},
	}, VoteCount: 7}
	topic := TopicUser{1, "blah", "Blah", "Hey there!", 0, false, false, now, RelativeTime(now), now, RelativeTime(now), 0, "", "127.0.0.1", 0, 1, "classname", poll.ID, "weird-data", BuildProfileURL("fake-user", 62), "Fake User", Config.DefaultGroup, "", 0, "", "", "", "", "", 58, false}
	var replyList []ReplyUser
	replyList = append(replyList, ReplyUser{0, 0, "Yo!", "Yo!", 0, "alice", "Alice", Config.DefaultGroup, now, RelativeTime(now), 0, 0, "", "", 0, "", "", "", "", 0, "127.0.0.1", false, 1, "", ""})

	var varList = make(map[string]tmpl.VarItem)
	tpage := TopicPage{"Title", user, headerVars, replyList, topic, poll, 1, 1}
	topicIDTmpl, err := c.Compile("topic.html", "templates/", "common.TopicPage", tpage, varList)
	if err != nil {
		return err
	}
	topicIDAltTmpl, err := c.Compile("topic_alt.html", "templates/", "common.TopicPage", tpage, varList)
	if err != nil {
		return err
	}

	varList = make(map[string]tmpl.VarItem)
	ppage := ProfilePage{"User 526", user, headerVars, replyList, user}
	profileTmpl, err := c.Compile("profile.html", "templates/", "common.ProfilePage", ppage, varList)
	if err != nil {
		return err
	}

	// TODO: Use a dummy forum list to avoid o(n) problems
	var forumList []Forum
	forums, err := Forums.GetAll()
	if err != nil {
		return err
	}

	for _, forum := range forums {
		forumList = append(forumList, *forum)
	}
	varList = make(map[string]tmpl.VarItem)
	forumsPage := ForumsPage{"Forum List", user, headerVars, forumList}
	forumsTmpl, err := c.Compile("forums.html", "templates/", "common.ForumsPage", forumsPage, varList)
	if err != nil {
		return err
	}

	var topicsList []*TopicsRow
	topicsList = append(topicsList, &TopicsRow{1, "topic-title", "Topic Title", "The topic content.", 1, false, false, "Date", time.Now(), "Date", user3.ID, 1, "", "127.0.0.1", 0, 1, "classname", "", &user2, "", 0, &user3, "General", "/forum/general.2"})
	topicsPage := TopicsPage{"Topic List", user, headerVars, topicsList, forumList, Config.DefaultForum, []int{1}, 1, 1}
	topicsTmpl, err := c.Compile("topics.html", "templates/", "common.TopicsPage", topicsPage, varList)
	if err != nil {
		return err
	}

	//var topicList []TopicUser
	//topicList = append(topicList,TopicUser{1,"topic-title","Topic Title","The topic content.",1,false,false,"Date","Date",1,"","127.0.0.1",0,1,"classname","","admin-fred","Admin Fred",config.DefaultGroup,"",0,"","","","",58,false})
	forumItem := BlankForum(1, "general-forum.1", "General Forum", "Where the general stuff happens", true, "all", 0, "", 0)
	forumPage := ForumPage{"General Forum", user, headerVars, topicsList, forumItem, []int{1}, 1, 1}
	forumTmpl, err := c.Compile("forum.html", "templates/", "common.ForumPage", forumPage, varList)
	if err != nil {
		return err
	}

	loginPage := Page{"Login Page", user, headerVars, tList, nil}
	loginTmpl, err := c.Compile("login.html", "templates/", "common.Page", loginPage, varList)
	if err != nil {
		return err
	}

	registerPage := Page{"Registration Page", user, headerVars, tList, nil}
	registerTmpl, err := c.Compile("register.html", "templates/", "common.Page", registerPage, varList)
	if err != nil {
		return err
	}

	errorPage := Page{"Error", user, headerVars, tList, "A problem has occurred in the system."}
	errorTmpl, err := c.Compile("error.html", "templates/", "common.Page", errorPage, varList)
	if err != nil {
		return err
	}

	var ipUserList = make(map[int]*User)
	ipUserList[1] = &user2
	ipSearchPage := IPSearchPage{"IP Search", user2, headerVars, ipUserList, "::1"}
	ipSearchTmpl, err := c.Compile("ip_search.html", "templates/", "common.IPSearchPage", ipSearchPage, varList)
	if err != nil {
		return err
	}

	// Let plugins register their own templates
	DebugLog("Registering the templates for the plugins")
	config = c.GetConfig()
	config.SkipHandles = true
	c.SetConfig(config)
	for _, tmplfunc := range PrebuildTmplList {
		tmplItem := tmplfunc(user, headerVars)
		varList = make(map[string]tmpl.VarItem)
		compiledTmpl, err := c.Compile(tmplItem.Filename, tmplItem.Path, tmplItem.StructName, tmplItem.Data, varList, tmplItem.Imports...)
		if err != nil {
			return err
		}
		go writeTemplate(tmplItem.Name, compiledTmpl)
	}

	log.Print("Writing the templates")
	go writeTemplate("topic", topicIDTmpl)
	go writeTemplate("topic_alt", topicIDAltTmpl)
	go writeTemplate("profile", profileTmpl)
	go writeTemplate("forums", forumsTmpl)
	go writeTemplate("topics", topicsTmpl)
	go writeTemplate("forum", forumTmpl)
	go writeTemplate("login", loginTmpl)
	go writeTemplate("register", registerTmpl)
	go writeTemplate("ip_search", ipSearchTmpl)
	go writeTemplate("error", errorTmpl)
	go func() {
		out := "package main\n\n"
		for templateName, count := range c.TemplateFragmentCount {
			out += "var " + templateName + "_frags = make([][]byte," + strconv.Itoa(count) + ")\n"
		}
		out += "\n// nolint\nfunc init() {\n" + c.FragOut + "}\n"
		err := writeFile("./template_list.go", out)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func writeTemplate(name string, content string) {
	err := writeFile("./template_"+name+".go", content)
	if err != nil {
		log.Fatal(err)
	}
}

func InitTemplates() error {
	if Dev.DebugMode {
		log.Print("Initialising the template system")
	}
	err := compileTemplates()
	if err != nil {
		return err
	}

	// TODO: Add support for 64-bit integers
	// TODO: Add support for floats
	fmap := make(map[string]interface{})
	fmap["add"] = func(left interface{}, right interface{}) interface{} {
		var leftInt, rightInt int
		switch left := left.(type) {
		case uint, uint8, uint16, int, int32:
			leftInt = left.(int)
		}
		switch right := right.(type) {
		case uint, uint8, uint16, int, int32:
			rightInt = right.(int)
		}
		return leftInt + rightInt
	}

	fmap["subtract"] = func(left interface{}, right interface{}) interface{} {
		var leftInt, rightInt int
		switch left := left.(type) {
		case uint, uint8, uint16, int, int32:
			leftInt = left.(int)
		}
		switch right := right.(type) {
		case uint, uint8, uint16, int, int32:
			rightInt = right.(int)
		}
		return leftInt - rightInt
	}

	fmap["multiply"] = func(left interface{}, right interface{}) interface{} {
		var leftInt, rightInt int
		switch left := left.(type) {
		case uint, uint8, uint16, int, int32:
			leftInt = left.(int)
		}
		switch right := right.(type) {
		case uint, uint8, uint16, int, int32:
			rightInt = right.(int)
		}
		return leftInt * rightInt
	}

	fmap["divide"] = func(left interface{}, right interface{}) interface{} {
		var leftInt, rightInt int
		switch left := left.(type) {
		case uint, uint8, uint16, int, int32:
			leftInt = left.(int)
		}
		switch right := right.(type) {
		case uint, uint8, uint16, int, int32:
			rightInt = right.(int)
		}
		if leftInt == 0 || rightInt == 0 {
			return 0
		}
		return leftInt / rightInt
	}

	fmap["dock"] = func(dock interface{}, headerVarInt interface{}) interface{} {
		return template.HTML(BuildWidget(dock.(string), headerVarInt.(*HeaderVars)))
	}

	fmap["lang"] = func(phraseNameInt interface{}) interface{} {
		phraseName, ok := phraseNameInt.(string)
		if !ok {
			panic("phraseNameInt is not a string")
		}
		return GetTmplPhrase(phraseName) // TODO: Log non-existent phrases?
	}

	// The interpreted templates...
	DebugLog("Loading the template files...")
	Templates.Funcs(fmap)
	template.Must(Templates.ParseGlob("templates/*"))
	template.Must(Templates.ParseGlob("pages/*"))

	return nil
}
