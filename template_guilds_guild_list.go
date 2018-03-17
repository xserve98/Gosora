// +build !no_templategen

// Code generated by Gosora. More below:
/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */
package main
import "strconv"
import "net/http"
import "./common"
import "./extend/guilds/lib"

var guilds_guild_list_tmpl_phrase_id int

// nolint
func init() {
	common.TmplPtrMap["o_guilds_guild_list"] = Template_guilds_guild_list
	guilds_guild_list_tmpl_phrase_id = common.RegisterTmplPhraseNames([]string{
		"menu_forums_aria",
		"menu_forums_tooltip",
		"menu_topics_aria",
		"menu_topics_tooltip",
		"menu_alert_counter_aria",
		"menu_alert_list_aria",
		"menu_account_aria",
		"menu_account_tooltip",
		"menu_profile_aria",
		"menu_profile_tooltip",
		"menu_panel_aria",
		"menu_panel_tooltip",
		"menu_logout_aria",
		"menu_logout_tooltip",
		"menu_register_aria",
		"menu_register_tooltip",
		"menu_login_aria",
		"menu_login_tooltip",
		"menu_hamburger_tooltip",
		"footer_powered_by",
		"footer_made_with_love",
		"footer_theme_selector_aria",
	})
}

// nolint
func Template_guilds_guild_list(tmpl_guilds_guild_list_vars guilds.ListPage, w http.ResponseWriter) error {
var phrases = common.GetTmplPhrasesBytes(guilds_guild_list_tmpl_phrase_id)
w.Write(header_frags[0])
w.Write([]byte(tmpl_guilds_guild_list_vars.Title))
w.Write(header_frags[1])
w.Write([]byte(tmpl_guilds_guild_list_vars.Header.Site.Name))
w.Write(header_frags[2])
w.Write([]byte(tmpl_guilds_guild_list_vars.Header.Theme.Name))
w.Write(header_frags[3])
if len(tmpl_guilds_guild_list_vars.Header.Stylesheets) != 0 {
for _, item := range tmpl_guilds_guild_list_vars.Header.Stylesheets {
w.Write(header_frags[4])
w.Write([]byte(item))
w.Write(header_frags[5])
}
}
w.Write(header_frags[6])
if len(tmpl_guilds_guild_list_vars.Header.Scripts) != 0 {
for _, item := range tmpl_guilds_guild_list_vars.Header.Scripts {
w.Write(header_frags[7])
w.Write([]byte(item))
w.Write(header_frags[8])
}
}
w.Write(header_frags[9])
w.Write([]byte(tmpl_guilds_guild_list_vars.CurrentUser.Session))
w.Write(header_frags[10])
w.Write([]byte(tmpl_guilds_guild_list_vars.Header.Site.URL))
w.Write(header_frags[11])
if tmpl_guilds_guild_list_vars.Header.MetaDesc != "" {
w.Write(header_frags[12])
w.Write([]byte(tmpl_guilds_guild_list_vars.Header.MetaDesc))
w.Write(header_frags[13])
}
w.Write(header_frags[14])
if !tmpl_guilds_guild_list_vars.CurrentUser.IsSuperMod {
w.Write(header_frags[15])
}
w.Write(header_frags[16])
w.Write(menu_frags[0])
w.Write([]byte(common.BuildWidget("leftOfNav",tmpl_guilds_guild_list_vars.Header)))
w.Write(menu_frags[1])
w.Write(menu_frags[2])
w.Write([]byte(tmpl_guilds_guild_list_vars.Header.Site.ShortName))
w.Write(menu_frags[3])
w.Write(phrases[0])
w.Write(menu_frags[4])
w.Write(phrases[1])
w.Write(menu_frags[5])
w.Write(phrases[2])
w.Write(menu_frags[6])
w.Write(phrases[3])
w.Write(menu_frags[7])
w.Write(phrases[4])
w.Write(menu_frags[8])
w.Write(phrases[5])
w.Write(menu_frags[9])
if tmpl_guilds_guild_list_vars.CurrentUser.Loggedin {
w.Write(menu_frags[10])
w.Write(phrases[6])
w.Write(menu_frags[11])
w.Write(phrases[7])
w.Write(menu_frags[12])
w.Write([]byte(tmpl_guilds_guild_list_vars.CurrentUser.Link))
w.Write(menu_frags[13])
w.Write(phrases[8])
w.Write(menu_frags[14])
w.Write(phrases[9])
w.Write(menu_frags[15])
w.Write(phrases[10])
w.Write(menu_frags[16])
w.Write(phrases[11])
w.Write(menu_frags[17])
w.Write([]byte(tmpl_guilds_guild_list_vars.CurrentUser.Session))
w.Write(menu_frags[18])
w.Write(phrases[12])
w.Write(menu_frags[19])
w.Write(phrases[13])
w.Write(menu_frags[20])
} else {
w.Write(menu_frags[21])
w.Write(phrases[14])
w.Write(menu_frags[22])
w.Write(phrases[15])
w.Write(menu_frags[23])
w.Write(phrases[16])
w.Write(menu_frags[24])
w.Write(phrases[17])
w.Write(menu_frags[25])
}
w.Write(menu_frags[26])
w.Write(phrases[18])
w.Write(menu_frags[27])
w.Write([]byte(common.BuildWidget("rightOfNav",tmpl_guilds_guild_list_vars.Header)))
w.Write(menu_frags[28])
w.Write(header_frags[17])
if tmpl_guilds_guild_list_vars.Header.Widgets.RightSidebar != "" {
w.Write(header_frags[18])
}
w.Write(header_frags[19])
if len(tmpl_guilds_guild_list_vars.Header.NoticeList) != 0 {
for _, item := range tmpl_guilds_guild_list_vars.Header.NoticeList {
w.Write(header_frags[20])
w.Write([]byte(item))
w.Write(header_frags[21])
}
}
w.Write(header_frags[22])
w.Write(guilds_guild_list_frags[0])
if len(tmpl_guilds_guild_list_vars.GuildList) != 0 {
for _, item := range tmpl_guilds_guild_list_vars.GuildList {
w.Write(guilds_guild_list_frags[1])
w.Write([]byte(item.Link))
w.Write(guilds_guild_list_frags[2])
w.Write([]byte(item.Name))
w.Write(guilds_guild_list_frags[3])
w.Write([]byte(item.Desc))
w.Write(guilds_guild_list_frags[4])
w.Write([]byte(strconv.Itoa(item.MemberCount)))
w.Write(guilds_guild_list_frags[5])
w.Write([]byte(item.LastUpdateTime))
w.Write(guilds_guild_list_frags[6])
}
} else {
w.Write(guilds_guild_list_frags[7])
}
w.Write(guilds_guild_list_frags[8])
w.Write(footer_frags[0])
w.Write([]byte(common.BuildWidget("footer",tmpl_guilds_guild_list_vars.Header)))
w.Write(footer_frags[1])
w.Write(phrases[19])
w.Write(footer_frags[2])
w.Write(phrases[20])
w.Write(footer_frags[3])
w.Write(phrases[21])
w.Write(footer_frags[4])
if len(tmpl_guilds_guild_list_vars.Header.Themes) != 0 {
for _, item := range tmpl_guilds_guild_list_vars.Header.Themes {
if !item.HideFromThemes {
w.Write(footer_frags[5])
w.Write([]byte(item.Name))
w.Write(footer_frags[6])
if tmpl_guilds_guild_list_vars.Header.Theme.Name == item.Name {
w.Write(footer_frags[7])
}
w.Write(footer_frags[8])
w.Write([]byte(item.FriendlyName))
w.Write(footer_frags[9])
}
}
}
w.Write(footer_frags[10])
w.Write([]byte(common.BuildWidget("rightSidebar",tmpl_guilds_guild_list_vars.Header)))
w.Write(footer_frags[11])
	return nil
}
