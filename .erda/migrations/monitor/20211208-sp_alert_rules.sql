UPDATE `sp_alert_rules` SET enable='0' WHERE alert_index IN ('dice_component_container_mem','kubernetes_instance_mem','kubernetes_node','machine_disk_util','machine_netdisk','machine_netdisk_used') AND enable='1';
INSERT `sp_alert_rules`(`alert_index`,`alert_scope`,`alert_type`,`attributes`,`name`,`template`) VALUES("dice_component_container_mem","org","dice_component","{\"alert_group\":\"{{cluster_name}}-{{container_id}}\",\"level\":\"Critical\",\"recover\":\"true\",\"tickets_metric_key\":\"{{container_id}}\"}","平台组件实例内存状态","{\"filters\":[{\"operator\":\"eq\",\"tag\":\"org_name\",\"value\":\"$org_name\"},{\"operator\":\"eq\",\"tag\":\"pod_namespace\",\"value\":\"default\"},{\"operator\":\"any\",\"tag\":\"component_name\"}],\"functions\":[{\"aggregator\":\"avg\",\"field\":\"mem_usage_percent\",\"operator\":\"gt\",\"value\":70},{\"aggregator\":\"max\",\"alias\":\"mem_usage_value\",\"field\":\"mem_usage\"},{\"aggregator\":\"max\",\"alias\":\"mem_limit_value\",\"field\":\"mem_limit\"},{\"aggregator\":\"max\",\"alias\":\"mem_max_usage_value\",\"field\":\"mem_max_usage\"},{\"aggregator\":\"max\",\"alias\":\"mem_allocation_value\",\"field\":\"mem_allocation\"}],\"group\":[\"cluster_name\",\"container_id\"],\"metrics\":[\"docker_container_summary\"],\"outputs\":[\"alert\"],\"select\":{\"_meta\":\"#_meta\",\"_metric_scope\":\"#_metric_scope\",\"_metric_scope_id\":\"#_metric_scope_id\",\"cluster_name\":\"#cluster_name\",\"component_name\":\"#component_name\",\"container_id\":\"#container_id\",\"host_ip\":\"#host_ip\",\"org_name\":\"#org_name\",\"pod_ip\":\"#pod_ip\",\"pod_name\":\"#pod_name\",\"pod_namespace\":\"#pod_namespace\"},\"window\":2}");
INSERT `sp_alert_rules`(`alert_index`,`alert_scope`,`alert_type`,`attributes`,`name`,`template`) VALUES("kubernetes_instance_mem","org","kubernetes","{\"alert_group\":\"{{cluster_name}}-{{container_id}}\",\"level\":\"Critical\",\"recover\":\"true\",\"tickets_metric_key\":\"{{container_id}}\"}","kubernetes组件实例内存状态","{\"filters\":[{\"operator\":\"eq\",\"tag\":\"org_name\",\"value\":\"$org_name\"},{\"operator\":\"eq\",\"tag\":\"pod_namespace\",\"value\":\"kube-system\"}],\"functions\":[{\"aggregator\":\"avg\",\"field\":\"mem_usage_percent\",\"operator\":\"gt\",\"value\":70},{\"aggregator\":\"max\",\"alias\":\"mem_usage_value\",\"field\":\"mem_usage\"},{\"aggregator\":\"max\",\"alias\":\"mem_limit_value\",\"field\":\"mem_limit\"},{\"aggregator\":\"max\",\"alias\":\"mem_max_usage_value\",\"field\":\"mem_max_usage\"},{\"aggregator\":\"max\",\"alias\":\"mem_allocation_value\",\"field\":\"mem_allocation\"}],\"group\":[\"cluster_name\",\"container_id\"],\"metrics\":[\"docker_container_summary\"],\"outputs\":[\"alert\"],\"select\":{\"_meta\":\"#_meta\",\"_metric_scope\":\"#_metric_scope\",\"_metric_scope_id\":\"#_metric_scope_id\",\"cluster_name\":\"#cluster_name\",\"component_name\":\"#pod_name\",\"container_id\":\"#container_id\",\"host_ip\":\"#host_ip\",\"org_name\":\"#org_name\",\"pod_ip\":\"#pod_ip\",\"pod_name\":\"#pod_name\",\"pod_namespace\":\"#pod_namespace\"},\"window\":2}");
INSERT `sp_alert_rules`(`alert_index`,`alert_scope`,`alert_type`,`attributes`,`name`,`template`) VALUES("kubernetes_node","org","kubernetes","{\"alert_group\":\"{{cluster_name}}-{{node_name}}\",\"level\":\"Fatal\",\"recover\":\"true\",\"tickets_metric_key\":\"{{node_name}}\"}","kubernetes节点异常","{\"filters\":[{\"operator\":\"eq\",\"tag\":\"org_name\",\"value\":\"$org_name\"},{\"operator\":\"neq\",\"tag\":\"offline\",\"value\":\"true\"}],\"functions\":[{\"aggregator\":\"values\",\"field\":\"ready\",\"operator\":\"all\",\"value\":false},{\"aggregator\":\"max\",\"alias\":\"allocatable_pods_value\",\"field\":\"allocatable_pods\"},{\"aggregator\":\"max\",\"alias\":\"capacity_pods_value\",\"field\":\"capacity_pods\"},{\"aggregator\":\"max\",\"alias\":\"allocatable_cpu_cores_value\",\"field\":\"allocatable_cpu_cores\"},{\"aggregator\":\"max\",\"alias\":\"capacity_cpu_cores_value\",\"field\":\"capacity_cpu_cores\"},{\"aggregator\":\"max\",\"alias\":\"allocatable_memory_bytes_value\",\"field\":\"allocatable_memory_bytes\"},{\"aggregator\":\"max\",\"alias\":\"capacity_memory_bytes_value\",\"field\":\"capacity_memory_bytes\"}],\"group\":[\"cluster_name\",\"node_name\"],\"metrics\":[\"kubernetes_node\"],\"outputs\":[\"alert\"],\"select\":{\"_meta\":\"#_meta\",\"_metric_scope\":\"#_metric_scope\",\"_metric_scope_id\":\"#_metric_scope_id\",\"cluster_name\":\"#cluster_name\",\"component_name\":\"#node_name\",\"node_name\":\"#node_name\",\"org_name\":\"#org_name\"},\"window\":1}");
INSERT `sp_alert_rules`(`alert_index`,`alert_scope`,`alert_type`,`attributes`,`name`,`template`) VALUES("machine_disk_util","org","machine","{\"alert_group\":\"{{cluster_name}}-{{host_ip}}\",\"display_url_id\":\"machine_disk_io\",\"level\":\"Critical\",\"recover\":\"true\",\"tickets_metric_key\":\"{{host_ip}}\"}","机器磁盘IO","{\"filters\":[{\"operator\":\"eq\",\"tag\":\"org_name\",\"value\":\"$org_name\"},{\"operator\":\"eq\",\"tag\":\"path\",\"value\":null}],\"functions\":[{\"aggregator\":\"p75\",\"field\":\"pct_util\",\"operator\":\"gte\",\"value\":95}],\"group\":[\"host_ip\"],\"metric\":\"disk\",\"outputs\":[\"alert\"],\"select\":{\"_meta\":\"#_meta\",\"_metric_scope\":\"#_metric_scope\",\"_metric_scope_id\":\"#_metric_scope_id\",\"cluster_name\":\"#cluster_name\",\"host_ip\":\"#host_ip\",\"labels\":\"#labels\",\"org_name\":\"#org_name\"},\"window\":5}");
INSERT `sp_alert_rules`(`alert_index`,`alert_scope`,`alert_type`,`attributes`,`name`,`template`) VALUES("machine_netdisk","org","machine","{\"alert_group\":\"{{cluster_name}}\",\"level\":\"Notice\",\"recover\":\"true\",\"tickets_metric_key\":\"{{cluster_name}}\"}","网盘","{\"filters\":[{\"operator\":\"eq\",\"tag\":\"org_name\",\"value\":\"$org_name\"},{\"operator\":\"eq\",\"tag\":\"type\",\"value\":\"netdata\"}],\"functions\":[{\"aggregator\":\"max\",\"alias\":\"used_percent_value\",\"field\":\"used_percent\",\"operator\":\"gte\",\"value\":75},{\"aggregator\":\"max\",\"alias\":\"total_value\",\"field\":\"total\"}],\"group\":null,\"metric\":\"disk\",\"outputs\":[\"alert\"],\"select\":{\"_meta\":\"#_meta\",\"_metric_scope\":\"#_metric_scope\",\"_metric_scope_id\":\"#_metric_scope_id\",\"cluster_name\":\"#cluster_name\",\"labels\":\"#labels\",\"org_name\":\"#org_name\"},\"window\":5}");
INSERT `sp_alert_rules`(`alert_index`,`alert_scope`,`alert_type`,`attributes`,`name`,`template`) VALUES("machine_netdisk_used","org","machine","{\"alert_group\":\"{{cluster_name}}\",\"level\":\"Fatal\",\"recover\":\"true\",\"tickets_metric_key\":\"{{cluster_name}}\"}","网盘容量使用量异常告警","{\"filters\":[{\"operator\":\"eq\",\"tag\":\"org_name\",\"value\":\"$org_name\"},{\"operator\":\"eq\",\"tag\":\"type\",\"value\":\"netdata\"}],\"functions\":[{\"aggregator\":\"max\",\"alias\":\"used\",\"field\":\"used\",\"field_script\":\"function invoke(field, tag){ return field.used/1024/1024/1024 }\",\"operator\":\"gte\",\"value\":300},{\"aggregator\":\"max\",\"alias\":\"total\",\"field_script\":\"function invoke(field, tag){ return field.total/1024/1024/1024 }\"}],\"group\":null,\"metric\":\"disk\",\"outputs\":[\"alert\"],\"select\":{\"_meta\":\"#_meta\",\"_metric_scope\":\"#_metric_scope\",\"_metric_scope_id\":\"#_metric_scope_id\",\"cluster_name\":\"#cluster_name\",\"labels\":\"#labels\",\"org_name\":\"#org_name\"},\"window\":5}");