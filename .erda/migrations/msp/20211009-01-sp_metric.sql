ALTER TABLE `sp_metric` ADD COLUMN `config` MEDIUMTEXT NOT NULL COMMENT '入参' AFTER `mode`;