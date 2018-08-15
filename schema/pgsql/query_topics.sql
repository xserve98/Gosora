CREATE TABLE "topics" (
	`tid` serial not null,
	`title` varchar (100) not null,
	`content` text not null,
	`parsed_content` text not null,
	`createdAt` timestamp not null,
	`lastReplyAt` timestamp not null,
	`lastReplyBy` int not null,
	`createdBy` int not null,
	`is_closed` boolean DEFAULT 0 not null,
	`sticky` boolean DEFAULT 0 not null,
	`parentID` int DEFAULT 2 not null,
	`ipaddress` varchar (200) DEFAULT '0.0.0.0.0' not null,
	`postCount` int DEFAULT 1 not null,
	`likeCount` int DEFAULT 0 not null,
	`words` int DEFAULT 0 not null,
	`views` int DEFAULT 0 not null,
	`css_class` varchar (100) DEFAULT '' not null,
	`poll` int DEFAULT 0 not null,
	`data` varchar (200) DEFAULT '' not null,
	primary key(`tid`)
);