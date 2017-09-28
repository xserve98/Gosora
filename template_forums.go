// +build !no_templategen

// Code generated by Gosora. More below:
/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */
package main
import "net/http"

// nolint
func init() {
	template_forums_handle = template_forums
	//o_template_forums_handle = template_forums
	ctemplates = append(ctemplates,"forums")
	tmplPtrMap["forums"] = &template_forums_handle
	tmplPtrMap["o_forums"] = template_forums
}

// nolint
func template_forums(tmpl_forums_vars ForumsPage, w http.ResponseWriter) {
w.Write(header_0)
w.Write([]byte(tmpl_forums_vars.Title))
w.Write(header_1)
w.Write([]byte(tmpl_forums_vars.Header.ThemeName))
w.Write(header_2)
if len(tmpl_forums_vars.Header.Stylesheets) != 0 {
for _, item := range tmpl_forums_vars.Header.Stylesheets {
w.Write(header_3)
w.Write([]byte(item))
w.Write(header_4)
}
}
w.Write(header_5)
if len(tmpl_forums_vars.Header.Scripts) != 0 {
for _, item := range tmpl_forums_vars.Header.Scripts {
w.Write(header_6)
w.Write([]byte(item))
w.Write(header_7)
}
}
w.Write(header_8)
w.Write([]byte(tmpl_forums_vars.CurrentUser.Session))
w.Write(header_9)
if !tmpl_forums_vars.CurrentUser.IsSuperMod {
w.Write(header_10)
}
w.Write(header_11)
w.Write(menu_0)
w.Write(menu_1)
w.Write([]byte(tmpl_forums_vars.Header.Site.Name))
w.Write(menu_2)
if tmpl_forums_vars.CurrentUser.Loggedin {
w.Write(menu_3)
w.Write([]byte(tmpl_forums_vars.CurrentUser.Link))
w.Write(menu_4)
w.Write([]byte(tmpl_forums_vars.CurrentUser.Session))
w.Write(menu_5)
} else {
w.Write(menu_6)
}
w.Write(menu_7)
w.Write(header_12)
if tmpl_forums_vars.Header.Widgets.RightSidebar != "" {
w.Write(header_13)
}
w.Write(header_14)
if len(tmpl_forums_vars.Header.NoticeList) != 0 {
for _, item := range tmpl_forums_vars.Header.NoticeList {
w.Write(header_15)
w.Write([]byte(item))
w.Write(header_16)
}
}
w.Write(forums_0)
if len(tmpl_forums_vars.ItemList) != 0 {
for _, item := range tmpl_forums_vars.ItemList {
w.Write(forums_1)
if item.Desc != "" || item.LastTopic.Title != "" {
w.Write(forums_2)
}
w.Write(forums_3)
if item.Desc != "" {
w.Write(forums_4)
w.Write([]byte(item.Link))
w.Write(forums_5)
w.Write([]byte(item.Name))
w.Write(forums_6)
w.Write([]byte(item.Desc))
w.Write(forums_7)
} else {
w.Write(forums_8)
w.Write([]byte(item.Link))
w.Write(forums_9)
w.Write([]byte(item.Name))
w.Write(forums_10)
}
w.Write(forums_11)
w.Write([]byte(item.LastTopic.Link))
w.Write(forums_12)
if item.LastTopic.Title != "" {
w.Write([]byte(item.LastTopic.Title))
} else {
w.Write(forums_13)
}
w.Write(forums_14)
if item.LastTopicTime != "" {
w.Write(forums_15)
w.Write([]byte(item.LastTopicTime))
w.Write(forums_16)
}
w.Write(forums_17)
}
} else {
w.Write(forums_18)
}
w.Write(forums_19)
w.Write(footer_0)
if len(tmpl_forums_vars.Header.Themes) != 0 {
for _, item := range tmpl_forums_vars.Header.Themes {
if !item.HideFromThemes {
w.Write(footer_1)
w.Write([]byte(item.Name))
w.Write(footer_2)
if tmpl_forums_vars.Header.ThemeName == item.Name {
w.Write(footer_3)
}
w.Write(footer_4)
w.Write([]byte(item.FriendlyName))
w.Write(footer_5)
}
}
}
w.Write(footer_6)
if tmpl_forums_vars.Header.Widgets.RightSidebar != "" {
w.Write(footer_7)
w.Write([]byte(string(tmpl_forums_vars.Header.Widgets.RightSidebar)))
w.Write(footer_8)
}
w.Write(footer_9)
}
