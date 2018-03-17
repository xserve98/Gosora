// +build !no_templategen

// Code generated by Gosora. More below:
/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */
package main
import "net/http"
import "./common"
import "strconv"

var topic_alt_tmpl_phrase_id int

// nolint
func init() {
	common.Template_topic_alt_handle = Template_topic_alt
	common.Ctemplates = append(common.Ctemplates,"topic_alt")
	common.TmplPtrMap["topic_alt"] = &common.Template_topic_alt_handle
	common.TmplPtrMap["o_topic_alt"] = Template_topic_alt
	topic_alt_tmpl_phrase_id = common.RegisterTmplPhraseNames([]string{
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
		"paginator_prev_page_aria",
		"paginator_less_than",
		"paginator_next_page_aria",
		"paginator_greater_than",
		"topic_opening_post_aria",
		"status_closed_tooltip",
		"topic_status_closed_aria",
		"topic_title_input_aria",
		"topic_update_button",
		"topic_userinfo_aria",
		"topic_level_prefix",
		"topic_poll_vote",
		"topic_poll_results",
		"topic_poll_cancel",
		"topic_opening_post_aria",
		"topic_userinfo_aria",
		"topic_level_prefix",
		"topic_like_aria",
		"topic_edit_aria",
		"topic_delete_aria",
		"topic_unlock_aria",
		"topic_lock_aria",
		"topic_unpin_aria",
		"topic_pin_aria",
		"topic_ip_full_tooltip",
		"topic_ip_full_aria",
		"topic_report_aria",
		"topic_like_count_aria",
		"topic_ip_full_tooltip",
		"topic_userinfo_aria",
		"topic_level_prefix",
		"topic_post_like_aria",
		"topic_post_edit_aria",
		"topic_post_delete_aria",
		"topic_ip_full_tooltip",
		"topic_ip_full_aria",
		"topic_report_aria",
		"topic_post_like_count_tooltip",
		"topic_your_information",
		"topic_level_prefix",
		"topic_reply_aria",
		"topic_reply_content_alt",
		"topic_reply_add_poll_option",
		"topic_reply_button",
		"topic_reply_add_poll_button",
		"topic_reply_add_file_button",
		"footer_powered_by",
		"footer_made_with_love",
		"footer_theme_selector_aria",
	})
}

// nolint
func Template_topic_alt(tmpl_topic_alt_vars common.TopicPage, w http.ResponseWriter) error {
var phrases = common.GetTmplPhrasesBytes(topic_alt_tmpl_phrase_id)
w.Write(header_frags[0])
w.Write([]byte(tmpl_topic_alt_vars.Title))
w.Write(header_frags[1])
w.Write([]byte(tmpl_topic_alt_vars.Header.Site.Name))
w.Write(header_frags[2])
w.Write([]byte(tmpl_topic_alt_vars.Header.Theme.Name))
w.Write(header_frags[3])
if len(tmpl_topic_alt_vars.Header.Stylesheets) != 0 {
for _, item := range tmpl_topic_alt_vars.Header.Stylesheets {
w.Write(header_frags[4])
w.Write([]byte(item))
w.Write(header_frags[5])
}
}
w.Write(header_frags[6])
if len(tmpl_topic_alt_vars.Header.Scripts) != 0 {
for _, item := range tmpl_topic_alt_vars.Header.Scripts {
w.Write(header_frags[7])
w.Write([]byte(item))
w.Write(header_frags[8])
}
}
w.Write(header_frags[9])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(header_frags[10])
w.Write([]byte(tmpl_topic_alt_vars.Header.Site.URL))
w.Write(header_frags[11])
if tmpl_topic_alt_vars.Header.MetaDesc != "" {
w.Write(header_frags[12])
w.Write([]byte(tmpl_topic_alt_vars.Header.MetaDesc))
w.Write(header_frags[13])
}
w.Write(header_frags[14])
if !tmpl_topic_alt_vars.CurrentUser.IsSuperMod {
w.Write(header_frags[15])
}
w.Write(header_frags[16])
w.Write(menu_frags[0])
w.Write([]byte(common.BuildWidget("leftOfNav",tmpl_topic_alt_vars.Header)))
w.Write(menu_frags[1])
w.Write(menu_frags[2])
w.Write([]byte(tmpl_topic_alt_vars.Header.Site.ShortName))
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
if tmpl_topic_alt_vars.CurrentUser.Loggedin {
w.Write(menu_frags[10])
w.Write(phrases[6])
w.Write(menu_frags[11])
w.Write(phrases[7])
w.Write(menu_frags[12])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Link))
w.Write(menu_frags[13])
w.Write(phrases[8])
w.Write(menu_frags[14])
w.Write(phrases[9])
w.Write(menu_frags[15])
w.Write(phrases[10])
w.Write(menu_frags[16])
w.Write(phrases[11])
w.Write(menu_frags[17])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
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
w.Write([]byte(common.BuildWidget("rightOfNav",tmpl_topic_alt_vars.Header)))
w.Write(menu_frags[28])
w.Write(header_frags[17])
if tmpl_topic_alt_vars.Header.Widgets.RightSidebar != "" {
w.Write(header_frags[18])
}
w.Write(header_frags[19])
if len(tmpl_topic_alt_vars.Header.NoticeList) != 0 {
for _, item := range tmpl_topic_alt_vars.Header.NoticeList {
w.Write(header_frags[20])
w.Write([]byte(item))
w.Write(header_frags[21])
}
}
w.Write(header_frags[22])
if tmpl_topic_alt_vars.Page > 1 {
w.Write(topic_alt_frags[0])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[1])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Page - 1)))
w.Write(topic_alt_frags[2])
w.Write(phrases[19])
w.Write(topic_alt_frags[3])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[4])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Page - 1)))
w.Write(topic_alt_frags[5])
w.Write(phrases[20])
w.Write(topic_alt_frags[6])
}
if tmpl_topic_alt_vars.LastPage != tmpl_topic_alt_vars.Page {
w.Write(topic_alt_frags[7])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[8])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Page + 1)))
w.Write(topic_alt_frags[9])
w.Write(phrases[21])
w.Write(topic_alt_frags[10])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[11])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Page + 1)))
w.Write(topic_alt_frags[12])
w.Write(phrases[22])
w.Write(topic_alt_frags[13])
}
w.Write(topic_alt_frags[14])
w.Write(phrases[23])
w.Write(topic_alt_frags[15])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[16])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[17])
if tmpl_topic_alt_vars.Topic.Sticky {
w.Write(topic_alt_frags[18])
} else {
if tmpl_topic_alt_vars.Topic.IsClosed {
w.Write(topic_alt_frags[19])
}
}
w.Write(topic_alt_frags[20])
w.Write([]byte(tmpl_topic_alt_vars.Topic.Title))
w.Write(topic_alt_frags[21])
if tmpl_topic_alt_vars.Topic.IsClosed {
w.Write(topic_alt_frags[22])
w.Write(phrases[24])
w.Write(topic_alt_frags[23])
w.Write(phrases[25])
w.Write(topic_alt_frags[24])
}
if tmpl_topic_alt_vars.CurrentUser.Perms.EditTopic {
w.Write(topic_alt_frags[25])
w.Write([]byte(tmpl_topic_alt_vars.Topic.Title))
w.Write(topic_alt_frags[26])
w.Write(phrases[26])
w.Write(topic_alt_frags[27])
w.Write(phrases[27])
w.Write(topic_alt_frags[28])
}
w.Write(topic_alt_frags[29])
if tmpl_topic_alt_vars.Poll.ID > 0 {
w.Write(topic_alt_frags[30])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_frags[31])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_frags[32])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[33])
w.Write(phrases[28])
w.Write(topic_alt_frags[34])
w.Write([]byte(tmpl_topic_alt_vars.Topic.Avatar))
w.Write(topic_alt_frags[35])
w.Write([]byte(tmpl_topic_alt_vars.Topic.UserLink))
w.Write(topic_alt_frags[36])
w.Write([]byte(tmpl_topic_alt_vars.Topic.CreatedByName))
w.Write(topic_alt_frags[37])
if tmpl_topic_alt_vars.Topic.Tag != "" {
w.Write(topic_alt_frags[38])
w.Write([]byte(tmpl_topic_alt_vars.Topic.Tag))
w.Write(topic_alt_frags[39])
} else {
w.Write(topic_alt_frags[40])
w.Write(phrases[29])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.Level)))
w.Write(topic_alt_frags[41])
}
w.Write(topic_alt_frags[42])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_frags[43])
if len(tmpl_topic_alt_vars.Poll.QuickOptions) != 0 {
for _, item := range tmpl_topic_alt_vars.Poll.QuickOptions {
w.Write(topic_alt_frags[44])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_frags[45])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_frags[46])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_frags[47])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_frags[48])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_frags[49])
w.Write([]byte(item.Value))
w.Write(topic_alt_frags[50])
}
}
w.Write(topic_alt_frags[51])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_frags[52])
w.Write(phrases[30])
w.Write(topic_alt_frags[53])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_frags[54])
w.Write(phrases[31])
w.Write(topic_alt_frags[55])
w.Write(phrases[32])
w.Write(topic_alt_frags[56])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_frags[57])
}
w.Write(topic_alt_frags[58])
w.Write(phrases[33])
w.Write(topic_alt_frags[59])
w.Write(phrases[34])
w.Write(topic_alt_frags[60])
w.Write([]byte(tmpl_topic_alt_vars.Topic.Avatar))
w.Write(topic_alt_frags[61])
w.Write([]byte(tmpl_topic_alt_vars.Topic.UserLink))
w.Write(topic_alt_frags[62])
w.Write([]byte(tmpl_topic_alt_vars.Topic.CreatedByName))
w.Write(topic_alt_frags[63])
if tmpl_topic_alt_vars.Topic.Tag != "" {
w.Write(topic_alt_frags[64])
w.Write([]byte(tmpl_topic_alt_vars.Topic.Tag))
w.Write(topic_alt_frags[65])
} else {
w.Write(topic_alt_frags[66])
w.Write(phrases[35])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.Level)))
w.Write(topic_alt_frags[67])
}
w.Write(topic_alt_frags[68])
w.Write([]byte(tmpl_topic_alt_vars.Topic.ContentHTML))
w.Write(topic_alt_frags[69])
w.Write([]byte(tmpl_topic_alt_vars.Topic.Content))
w.Write(topic_alt_frags[70])
if tmpl_topic_alt_vars.CurrentUser.Loggedin {
if tmpl_topic_alt_vars.CurrentUser.Perms.LikeItem {
w.Write(topic_alt_frags[71])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[72])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[73])
w.Write(phrases[36])
w.Write(topic_alt_frags[74])
}
if tmpl_topic_alt_vars.CurrentUser.Perms.EditTopic {
w.Write(topic_alt_frags[75])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[76])
w.Write(phrases[37])
w.Write(topic_alt_frags[77])
}
if tmpl_topic_alt_vars.CurrentUser.Perms.DeleteTopic {
w.Write(topic_alt_frags[78])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[79])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[80])
w.Write(phrases[38])
w.Write(topic_alt_frags[81])
}
if tmpl_topic_alt_vars.CurrentUser.Perms.CloseTopic {
if tmpl_topic_alt_vars.Topic.IsClosed {
w.Write(topic_alt_frags[82])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[83])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[84])
w.Write(phrases[39])
w.Write(topic_alt_frags[85])
} else {
w.Write(topic_alt_frags[86])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[87])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[88])
w.Write(phrases[40])
w.Write(topic_alt_frags[89])
}
}
if tmpl_topic_alt_vars.CurrentUser.Perms.PinTopic {
if tmpl_topic_alt_vars.Topic.Sticky {
w.Write(topic_alt_frags[90])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[91])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[92])
w.Write(phrases[41])
w.Write(topic_alt_frags[93])
} else {
w.Write(topic_alt_frags[94])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[95])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[96])
w.Write(phrases[42])
w.Write(topic_alt_frags[97])
}
}
if tmpl_topic_alt_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_alt_frags[98])
w.Write([]byte(tmpl_topic_alt_vars.Topic.IPAddress))
w.Write(topic_alt_frags[99])
w.Write(phrases[43])
w.Write(topic_alt_frags[100])
w.Write(phrases[44])
w.Write(topic_alt_frags[101])
}
w.Write(topic_alt_frags[102])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[103])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[104])
w.Write(phrases[45])
w.Write(topic_alt_frags[105])
}
w.Write(topic_alt_frags[106])
if tmpl_topic_alt_vars.Topic.LikeCount > 0 {
w.Write(topic_alt_frags[107])
}
w.Write(topic_alt_frags[108])
if tmpl_topic_alt_vars.Topic.LikeCount > 0 {
w.Write(topic_alt_frags[109])
w.Write(phrases[46])
w.Write(topic_alt_frags[110])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.LikeCount)))
w.Write(topic_alt_frags[111])
}
w.Write(topic_alt_frags[112])
w.Write([]byte(tmpl_topic_alt_vars.Topic.RelativeCreatedAt))
w.Write(topic_alt_frags[113])
if tmpl_topic_alt_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_alt_frags[114])
w.Write([]byte(tmpl_topic_alt_vars.Topic.IPAddress))
w.Write(topic_alt_frags[115])
w.Write(phrases[47])
w.Write(topic_alt_frags[116])
w.Write([]byte(tmpl_topic_alt_vars.Topic.IPAddress))
w.Write(topic_alt_frags[117])
}
w.Write(topic_alt_frags[118])
if len(tmpl_topic_alt_vars.ItemList) != 0 {
for _, item := range tmpl_topic_alt_vars.ItemList {
w.Write(topic_alt_frags[119])
if item.ActionType != "" {
w.Write(topic_alt_frags[120])
}
w.Write(topic_alt_frags[121])
w.Write(phrases[48])
w.Write(topic_alt_frags[122])
w.Write([]byte(item.Avatar))
w.Write(topic_alt_frags[123])
w.Write([]byte(item.UserLink))
w.Write(topic_alt_frags[124])
w.Write([]byte(item.CreatedByName))
w.Write(topic_alt_frags[125])
if item.Tag != "" {
w.Write(topic_alt_frags[126])
w.Write([]byte(item.Tag))
w.Write(topic_alt_frags[127])
} else {
w.Write(topic_alt_frags[128])
w.Write(phrases[49])
w.Write([]byte(strconv.Itoa(item.Level)))
w.Write(topic_alt_frags[129])
}
w.Write(topic_alt_frags[130])
if item.ActionType != "" {
w.Write(topic_alt_frags[131])
}
w.Write(topic_alt_frags[132])
if item.ActionType != "" {
w.Write(topic_alt_frags[133])
w.Write([]byte(item.ActionIcon))
w.Write(topic_alt_frags[134])
w.Write([]byte(item.ActionType))
w.Write(topic_alt_frags[135])
} else {
w.Write(topic_alt_frags[136])
w.Write([]byte(item.ContentHtml))
w.Write(topic_alt_frags[137])
if tmpl_topic_alt_vars.CurrentUser.Loggedin {
if tmpl_topic_alt_vars.CurrentUser.Perms.LikeItem {
w.Write(topic_alt_frags[138])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_frags[139])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[140])
w.Write(phrases[50])
w.Write(topic_alt_frags[141])
}
if tmpl_topic_alt_vars.CurrentUser.Perms.EditReply {
w.Write(topic_alt_frags[142])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_frags[143])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[144])
w.Write(phrases[51])
w.Write(topic_alt_frags[145])
}
if tmpl_topic_alt_vars.CurrentUser.Perms.DeleteReply {
w.Write(topic_alt_frags[146])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_frags[147])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[148])
w.Write(phrases[52])
w.Write(topic_alt_frags[149])
}
if tmpl_topic_alt_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_alt_frags[150])
w.Write([]byte(item.IPAddress))
w.Write(topic_alt_frags[151])
w.Write(phrases[53])
w.Write(topic_alt_frags[152])
w.Write(phrases[54])
w.Write(topic_alt_frags[153])
}
w.Write(topic_alt_frags[154])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_frags[155])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[156])
w.Write(phrases[55])
w.Write(topic_alt_frags[157])
}
w.Write(topic_alt_frags[158])
if item.LikeCount > 0 {
w.Write(topic_alt_frags[159])
}
w.Write(topic_alt_frags[160])
if item.LikeCount > 0 {
w.Write(topic_alt_frags[161])
w.Write(phrases[56])
w.Write(topic_alt_frags[162])
w.Write([]byte(strconv.Itoa(item.LikeCount)))
w.Write(topic_alt_frags[163])
}
w.Write(topic_alt_frags[164])
w.Write([]byte(item.RelativeCreatedAt))
w.Write(topic_alt_frags[165])
if tmpl_topic_alt_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_alt_frags[166])
w.Write([]byte(item.IPAddress))
w.Write(topic_alt_frags[167])
w.Write([]byte(item.IPAddress))
w.Write(topic_alt_frags[168])
}
w.Write(topic_alt_frags[169])
}
w.Write(topic_alt_frags[170])
}
}
w.Write(topic_alt_frags[171])
if tmpl_topic_alt_vars.CurrentUser.Perms.CreateReply {
w.Write(topic_alt_frags[172])
w.Write(phrases[57])
w.Write(topic_alt_frags[173])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Avatar))
w.Write(topic_alt_frags[174])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Link))
w.Write(topic_alt_frags[175])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Name))
w.Write(topic_alt_frags[176])
if tmpl_topic_alt_vars.CurrentUser.Tag != "" {
w.Write(topic_alt_frags[177])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Tag))
w.Write(topic_alt_frags[178])
} else {
w.Write(topic_alt_frags[179])
w.Write(phrases[58])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.CurrentUser.Level)))
w.Write(topic_alt_frags[180])
}
w.Write(topic_alt_frags[181])
w.Write(phrases[59])
w.Write(topic_alt_frags[182])
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_frags[183])
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_frags[184])
w.Write(phrases[60])
w.Write(topic_alt_frags[185])
w.Write(phrases[61])
w.Write(topic_alt_frags[186])
w.Write(phrases[62])
w.Write(topic_alt_frags[187])
w.Write(phrases[63])
w.Write(topic_alt_frags[188])
if tmpl_topic_alt_vars.CurrentUser.Perms.UploadFiles {
w.Write(topic_alt_frags[189])
w.Write(phrases[64])
w.Write(topic_alt_frags[190])
}
w.Write(topic_alt_frags[191])
}
w.Write(topic_alt_frags[192])
w.Write(footer_frags[0])
w.Write([]byte(common.BuildWidget("footer",tmpl_topic_alt_vars.Header)))
w.Write(footer_frags[1])
w.Write(phrases[65])
w.Write(footer_frags[2])
w.Write(phrases[66])
w.Write(footer_frags[3])
w.Write(phrases[67])
w.Write(footer_frags[4])
if len(tmpl_topic_alt_vars.Header.Themes) != 0 {
for _, item := range tmpl_topic_alt_vars.Header.Themes {
if !item.HideFromThemes {
w.Write(footer_frags[5])
w.Write([]byte(item.Name))
w.Write(footer_frags[6])
if tmpl_topic_alt_vars.Header.Theme.Name == item.Name {
w.Write(footer_frags[7])
}
w.Write(footer_frags[8])
w.Write([]byte(item.FriendlyName))
w.Write(footer_frags[9])
}
}
}
w.Write(footer_frags[10])
w.Write([]byte(common.BuildWidget("rightSidebar",tmpl_topic_alt_vars.Header)))
w.Write(footer_frags[11])
	return nil
}
