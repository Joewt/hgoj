set names utf8; 
create database jol;
use jol;

CREATE TABLE  `compileinfo` (
  `solution_id` int(11) NOT NULL DEFAULT 0,
  `error` text,
  PRIMARY KEY (`solution_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE  `contest` (
  `contest_id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `defunct` char(1) NOT NULL DEFAULT 'N',
  `description` text,
  `private` tinyint(4) NOT NULL DEFAULT '0',
  `langmask` int NOT NULL DEFAULT '0' COMMENT 'bits for LANG to mask',
  `password` CHAR( 16 ) NOT NULL DEFAULT '',
  `user_id` varchar(48) NOT NULL DEFAULT 'admin',
  PRIMARY KEY (`contest_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `contest_problem` (
  `problem_id` int(11) NOT NULL DEFAULT '0',
  `contest_id` int(11) DEFAULT NULL,
  `title` char(200) NOT NULL DEFAULT '',
  `num` int(11) NOT NULL DEFAULT '0',
  `c_accepted` int(11) NOT NULL DEFAULT '0',
  `c_submit` int(11) NOT NULL DEFAULT '0',
  KEY `Index_contest_id` (`contest_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `loginlog` (
  `user_id` varchar(48) NOT NULL DEFAULT '',
  `password` varchar(40) DEFAULT NULL,
  `ip` varchar(46) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  KEY `user_log_index` (`user_id`,`time`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE  `mail` (
  `mail_id` int(11) NOT NULL AUTO_INCREMENT,
  `to_user` varchar(48) NOT NULL DEFAULT '',
  `from_user` varchar(48) NOT NULL DEFAULT '',
  `title` varchar(200) NOT NULL DEFAULT '',
  `content` text,
  `new_mail` tinyint(1) NOT NULL DEFAULT '1',
  `reply` tinyint(4) DEFAULT '0',
  `in_date` datetime DEFAULT NULL,
  `defunct` char(1) NOT NULL DEFAULT 'N',
  PRIMARY KEY (`mail_id`),
  KEY `uid` (`to_user`)
) ENGINE=MyISAM AUTO_INCREMENT=1013 DEFAULT CHARSET=utf8;

CREATE TABLE  `news` (
  `news_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(48) NOT NULL DEFAULT '',
  `title` varchar(200) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `time` datetime NOT NULL DEFAULT '2016-05-13 19:24:00',
  `importance` tinyint(4) NOT NULL DEFAULT '0',
  `defunct` char(1) NOT NULL DEFAULT 'N',
  PRIMARY KEY (`news_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1004 DEFAULT CHARSET=utf8;


CREATE TABLE  `privilege` (
  `user_id` char(48) NOT NULL DEFAULT '',
  `rightstr` char(30) NOT NULL DEFAULT '',
  `defunct` char(1) NOT NULL DEFAULT 'N'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE  `problem` (
  `problem_id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL DEFAULT '',
  `description` text,
  `input` text,
  `output` text,
  `sample_input` text,
  `sample_output` text,
  `spj` char(1) NOT NULL DEFAULT '0',
  `hint` text,
  `source` varchar(100) DEFAULT NULL,
  `in_date` datetime DEFAULT NULL,
  `time_limit` int(11) NOT NULL DEFAULT 0,
  `memory_limit` int(11) NOT NULL DEFAULT 0,
  `defunct` char(1) NOT NULL DEFAULT 'N',
  `accepted` int(11) DEFAULT '0',
  `submit` int(11) DEFAULT '0',
  `solved` int(11) DEFAULT '0',
  PRIMARY KEY (`problem_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8;

CREATE TABLE  `reply` (
  `rid` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` varchar(48) NOT NULL,
  `time` datetime NOT NULL DEFAULT '2016-05-13 19:24:00',
  `content` text NOT NULL,
  `topic_id` int(11) NOT NULL,
  `status` int(2) NOT NULL DEFAULT '0',
  `ip` varchar(46) NOT NULL,
  PRIMARY KEY (`rid`),
  KEY `author_id` (`author_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE  `sim` (
  `s_id` int(11) NOT NULL,
  `sim_s_id` int(11) DEFAULT NULL,
  `sim` int(11) DEFAULT NULL,
  PRIMARY KEY (`s_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE  `solution` (
  `solution_id` int(11) NOT NULL AUTO_INCREMENT,
  `problem_id` int(11) NOT NULL DEFAULT 0,
  `user_id` char(48) NOT NULL,
  `time` int(11) NOT NULL DEFAULT 0,
  `memory` int(11) NOT NULL DEFAULT 0,
  `in_date` datetime NOT NULL DEFAULT '2016-05-13 19:24:00',
  `result` smallint(6) NOT NULL DEFAULT '0',
  `language` INT UNSIGNED NOT NULL DEFAULT '0',
  `ip` char(46) NOT NULL,
  `contest_id` int(11) DEFAULT 0,
  `valid` tinyint(4) NOT NULL DEFAULT '1',
  `num` tinyint(4) NOT NULL DEFAULT '-1',
  `code_length` int(11) NOT NULL DEFAULT 0,
  `judgetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `pass_rate` DECIMAL(3,2) UNSIGNED NOT NULL DEFAULT 0,
  `lint_error` int UNSIGNED NOT NULL DEFAULT 0,
  `judger` CHAR(16) NOT NULL DEFAULT 'LOCAL',
  PRIMARY KEY (`solution_id`),
  KEY `uid` (`user_id`),
  KEY `pid` (`problem_id`),
  KEY `res` (`result`),
  KEY `cid` (`contest_id`)
) ENGINE=MyISAM row_format=fixed AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8;

CREATE TABLE  `source_code` (
  `solution_id` int(11) NOT NULL,
  `source` text NOT NULL,
  PRIMARY KEY (`solution_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
create table source_code_user like source_code;

CREATE TABLE  `topic` (
  `tid` int(11) NOT NULL AUTO_INCREMENT,
  `title` varbinary(60) NOT NULL,
  `status` int(2) NOT NULL DEFAULT '0',
  `top_level` int(2) NOT NULL DEFAULT '0',
  `cid` int(11) DEFAULT NULL,
  `pid` int(11) NOT NULL,
  `author_id` varchar(48) NOT NULL,
  PRIMARY KEY (`tid`),
  KEY `cid` (`cid`,`pid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE  `users` (
  `user_id` varchar(48) NOT NULL DEFAULT '',
  `email` varchar(100) DEFAULT NULL,
  `submit` int(11) DEFAULT '0',
  `solved` int(11) DEFAULT '0',
  `defunct` char(1) NOT NULL DEFAULT 'N',
  `ip` varchar(46) NOT NULL DEFAULT '',
  `accesstime` datetime DEFAULT NULL,
  `volume` int(11) NOT NULL DEFAULT '1',
  `language` int(11) NOT NULL DEFAULT '1',
  `password` varchar(32) DEFAULT NULL,
  `reg_time` datetime DEFAULT NULL,
  `nick` varchar(20) NOT NULL DEFAULT '',
  `school` varchar(20) NOT NULL DEFAULT '',
  PRIMARY KEY (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `online` (
  `hash` varchar(32) collate utf8_unicode_ci NOT NULL,
  `ip` varchar(46) character set utf8 NOT NULL default '',
  `ua` varchar(255) character set utf8 NOT NULL default '',
  `refer` varchar(255) collate utf8_unicode_ci default NULL,
  `lastmove` int(10) NOT NULL,
  `firsttime` int(10) default NULL,
  `uri` varchar(255) collate utf8_unicode_ci default NULL,
  PRIMARY KEY  (`hash`),
  UNIQUE KEY `hash` (`hash`)
) ENGINE=MEMORY DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE  `runtimeinfo` (
  `solution_id` int(11) NOT NULL DEFAULT 0,
  `error` text,
  PRIMARY KEY (`solution_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE  `custominput` (
  `solution_id` int(11) NOT NULL DEFAULT 0,
  `input_text` text,
  PRIMARY KEY (`solution_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE  `printer` (
  `printer_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` char(48) NOT NULL,
  `in_date` datetime NOT NULL DEFAULT '2018-03-13 19:38:00',
  `status` smallint(6) NOT NULL DEFAULT '0',
  `worktime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `printer` CHAR(16) NOT NULL DEFAULT 'LOCAL',
  `content` text NOT NULL ,
  PRIMARY KEY (`printer_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE  `balloon` (
  `balloon_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` char(48) NOT NULL,
  `sid` int(11) NOT NULL ,
  `cid` int(11) NOT NULL ,
  `pid` int(11) NOT NULL ,
  `status` smallint(6) NOT NULL DEFAULT '0',
  PRIMARY KEY (`balloon_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `share_code` (
  `share_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(48) COLLATE utf8_unicode_ci DEFAULT NULL,
  `share_code` text COLLATE utf8_unicode_ci,
  `language` varchar(32) COLLATE utf8_unicode_ci DEFAULT NULL,
  `share_time` datetime DEFAULT NULL,
  PRIMARY KEY (`share_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8;

delimiter //
drop trigger if exists simfilter//
create trigger simfilter
before insert on sim
for each row
begin
 declare new_user_id varchar(64);
 declare old_user_id varchar(64);
 select user_id from solution where solution_id=new.s_id into new_user_id;
 select user_id from solution where solution_id=new.sim_s_id into old_user_id;
 if old_user_id=new_user_id then
	set new.s_id=0;
 end if;
 
end;//
delimiter ;
