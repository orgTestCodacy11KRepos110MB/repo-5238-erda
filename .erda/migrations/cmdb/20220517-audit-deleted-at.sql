ALTER TABLE `dice_audit` CHANGE `deleted` `soft_deleted_at` bigint(20) NOT NULL DEFAULT 0 COMMENT 'soft deleted at';
ALTER TABLE `dice_audit` ADD INDEX `idx_orgid_deleted_scopetype_starttime` (`org_id`, `soft_deleted_at`, `scope_type`, `start_time`);