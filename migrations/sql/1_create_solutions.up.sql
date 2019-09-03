CREATE TABLE `solutions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int(10) unsigned NOT NULL,
  `job_id` int(10) unsigned NOT NULL,
  `data` json NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `task_id` (`task_id`),
  KEY `job_id` (`job_id`)
) 

