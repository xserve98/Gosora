// +build pgsql

// This file was generated by Gosora's Query Generator. Please try to avoid modifying this file, as it might change at any time.
package main

import "log"
import "database/sql"

// nolint
var addRepliesToTopicStmt *sql.Stmt
var removeRepliesFromTopicStmt *sql.Stmt
var addTopicsToForumStmt *sql.Stmt
var removeTopicsFromForumStmt *sql.Stmt
var updateForumCacheStmt *sql.Stmt
var addLikesToTopicStmt *sql.Stmt
var addLikesToReplyStmt *sql.Stmt
var editTopicStmt *sql.Stmt
var editReplyStmt *sql.Stmt
var stickTopicStmt *sql.Stmt
var unstickTopicStmt *sql.Stmt
var lockTopicStmt *sql.Stmt
var unlockTopicStmt *sql.Stmt
var updateLastIPStmt *sql.Stmt
var updateSessionStmt *sql.Stmt
var setPasswordStmt *sql.Stmt
var setAvatarStmt *sql.Stmt
var setUsernameStmt *sql.Stmt
var changeGroupStmt *sql.Stmt
var activateUserStmt *sql.Stmt
var updateUserLevelStmt *sql.Stmt
var incrementUserScoreStmt *sql.Stmt
var incrementUserPostsStmt *sql.Stmt
var incrementUserBigpostsStmt *sql.Stmt
var incrementUserMegapostsStmt *sql.Stmt
var incrementUserTopicsStmt *sql.Stmt
var editProfileReplyStmt *sql.Stmt
var updateForumStmt *sql.Stmt
var updateSettingStmt *sql.Stmt
var updatePluginStmt *sql.Stmt
var updatePluginInstallStmt *sql.Stmt
var updateThemeStmt *sql.Stmt
var updateUserStmt *sql.Stmt
var updateGroupPermsStmt *sql.Stmt
var updateGroupRankStmt *sql.Stmt
var updateGroupStmt *sql.Stmt
var updateEmailStmt *sql.Stmt
var verifyEmailStmt *sql.Stmt
var setTempGroupStmt *sql.Stmt
var updateWordFilterStmt *sql.Stmt
var bumpSyncStmt *sql.Stmt

// nolint
func _gen_pgsql() (err error) {
	if dev.DebugMode {
		log.Print("Building the generated statements")
	}
	
	log.Print("Preparing addRepliesToTopic statement.")
	addRepliesToTopicStmt, err = db.Prepare("UPDATE `topics` SET `postCount` = `postCount` + ?,`lastReplyBy` = ?,`lastReplyAt` = LOCALTIMESTAMP() WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing removeRepliesFromTopic statement.")
	removeRepliesFromTopicStmt, err = db.Prepare("UPDATE `topics` SET `postCount` = `postCount` - ? WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing addTopicsToForum statement.")
	addTopicsToForumStmt, err = db.Prepare("UPDATE `forums` SET `topicCount` = `topicCount` + ? WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing removeTopicsFromForum statement.")
	removeTopicsFromForumStmt, err = db.Prepare("UPDATE `forums` SET `topicCount` = `topicCount` - ? WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateForumCache statement.")
	updateForumCacheStmt, err = db.Prepare("UPDATE `forums` SET `lastTopicID` = ?,`lastReplyerID` = ? WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing addLikesToTopic statement.")
	addLikesToTopicStmt, err = db.Prepare("UPDATE `topics` SET `likeCount` = `likeCount` + ? WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing addLikesToReply statement.")
	addLikesToReplyStmt, err = db.Prepare("UPDATE `replies` SET `likeCount` = `likeCount` + ? WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing editTopic statement.")
	editTopicStmt, err = db.Prepare("UPDATE `topics` SET `title` = ?,`content` = ?,`parsed_content` = ? WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing editReply statement.")
	editReplyStmt, err = db.Prepare("UPDATE `replies` SET `content` = ?,`parsed_content` = ? WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing stickTopic statement.")
	stickTopicStmt, err = db.Prepare("UPDATE `topics` SET `sticky` = 1 WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing unstickTopic statement.")
	unstickTopicStmt, err = db.Prepare("UPDATE `topics` SET `sticky` = 0 WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing lockTopic statement.")
	lockTopicStmt, err = db.Prepare("UPDATE `topics` SET `is_closed` = 1 WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing unlockTopic statement.")
	unlockTopicStmt, err = db.Prepare("UPDATE `topics` SET `is_closed` = 0 WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateLastIP statement.")
	updateLastIPStmt, err = db.Prepare("UPDATE `users` SET `last_ip` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateSession statement.")
	updateSessionStmt, err = db.Prepare("UPDATE `users` SET `session` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing setPassword statement.")
	setPasswordStmt, err = db.Prepare("UPDATE `users` SET `password` = ?,`salt` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing setAvatar statement.")
	setAvatarStmt, err = db.Prepare("UPDATE `users` SET `avatar` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing setUsername statement.")
	setUsernameStmt, err = db.Prepare("UPDATE `users` SET `name` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing changeGroup statement.")
	changeGroupStmt, err = db.Prepare("UPDATE `users` SET `group` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing activateUser statement.")
	activateUserStmt, err = db.Prepare("UPDATE `users` SET `active` = 1 WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateUserLevel statement.")
	updateUserLevelStmt, err = db.Prepare("UPDATE `users` SET `level` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing incrementUserScore statement.")
	incrementUserScoreStmt, err = db.Prepare("UPDATE `users` SET `score` = `score` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing incrementUserPosts statement.")
	incrementUserPostsStmt, err = db.Prepare("UPDATE `users` SET `posts` = `posts` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing incrementUserBigposts statement.")
	incrementUserBigpostsStmt, err = db.Prepare("UPDATE `users` SET `posts` = `posts` + ?,`bigposts` = `bigposts` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing incrementUserMegaposts statement.")
	incrementUserMegapostsStmt, err = db.Prepare("UPDATE `users` SET `posts` = `posts` + ?,`bigposts` = `bigposts` + ?,`megaposts` = `megaposts` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing incrementUserTopics statement.")
	incrementUserTopicsStmt, err = db.Prepare("UPDATE `users` SET `topics` = `topics` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing editProfileReply statement.")
	editProfileReplyStmt, err = db.Prepare("UPDATE `users_replies` SET `content` = ?,`parsed_content` = ? WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateForum statement.")
	updateForumStmt, err = db.Prepare("UPDATE `forums` SET `name` = ?,`desc` = ?,`active` = ?,`preset` = ? WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateSetting statement.")
	updateSettingStmt, err = db.Prepare("UPDATE `settings` SET `content` = ? WHERE `name` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updatePlugin statement.")
	updatePluginStmt, err = db.Prepare("UPDATE `plugins` SET `active` = ? WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updatePluginInstall statement.")
	updatePluginInstallStmt, err = db.Prepare("UPDATE `plugins` SET `installed` = ? WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateTheme statement.")
	updateThemeStmt, err = db.Prepare("UPDATE `themes` SET `default` = ? WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateUser statement.")
	updateUserStmt, err = db.Prepare("UPDATE `users` SET `name` = ?,`email` = ?,`group` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateGroupPerms statement.")
	updateGroupPermsStmt, err = db.Prepare("UPDATE `users_groups` SET `permissions` = ? WHERE `gid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateGroupRank statement.")
	updateGroupRankStmt, err = db.Prepare("UPDATE `users_groups` SET `is_admin` = ?,`is_mod` = ?,`is_banned` = ? WHERE `gid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateGroup statement.")
	updateGroupStmt, err = db.Prepare("UPDATE `users_groups` SET `name` = ?,`tag` = ? WHERE `gid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateEmail statement.")
	updateEmailStmt, err = db.Prepare("UPDATE `emails` SET `email` = ?,`uid` = ?,`validated` = ?,`token` = ? WHERE `email` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing verifyEmail statement.")
	verifyEmailStmt, err = db.Prepare("UPDATE `emails` SET `validated` = 1,`token` = '' WHERE `email` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing setTempGroup statement.")
	setTempGroupStmt, err = db.Prepare("UPDATE `users` SET `temp_group` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing updateWordFilter statement.")
	updateWordFilterStmt, err = db.Prepare("UPDATE `word_filters` SET `find` = ?,`replacement` = ? WHERE `wfid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing bumpSync statement.")
	bumpSyncStmt, err = db.Prepare("UPDATE `sync` SET `last_update` = LOCALTIMESTAMP()")
	if err != nil {
		return err
	}
	
	return nil
}
