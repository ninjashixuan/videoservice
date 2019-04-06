DROP TABLE IF EXISTS `users`;
CREATE TABLE users (
	id int unsigned not null auto_increment,
	username varchar(64) not null unique,
	pwd varchar(64) not null,
	primary key (id)
);

DROP TABLE IF EXISTS `video`;
CREATE TABLE video (
	id varchar(64) not null,
	author_id int(20) not null,
	info text,
	display_ctime text,
	create_time datetime default current_timestamp,
	primary key (id)
);

DROP TABLE IF EXISTS `comments`;
CREATE TABLE comments (
	id varchar(64) not null,
	author_id int(20) not null,
	video_id int(20),
	content text,
	create_time datetime default current_timestamp,
	primary key (id)
);

DROP TABLE IF EXISTS `sessions`;
CREATE TABLE sessions (
	session_id tinytext not null,
	TTL tinytext,
	username  text
);
alter table sessions add primary key (session_id(20));


DROP TABLE IF EXISTS `video_del_rec`;
CREATE TABLE video_del_rec (
    video_id varchar(64) not null,
    primary key (video_id)
);
