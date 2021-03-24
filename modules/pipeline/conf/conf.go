// Package conf 定义了 pipeline 所需要的配置选项，这些配置选项都是通过环境变量加载.
package conf

import (
	"path/filepath"
	"time"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/pkg/envconf"
)

// Conf 定义配置对象.
type Conf struct {
	ListenAddr  string `env:"LISTEN_ADDR" default:":3081"`
	Debug       bool   `env:"DEBUG" default:"false"`
	DiceCluster string `env:"DICE_CLUSTER" default:"local"` // 服务所在集群

	// task level
	TaskDefaultCPU      float64       `env:"TASK_DEFAULT_CPU" default:"0.5"`
	TaskDefaultMEM      float64       `env:"TASK_DEFAULT_MEM" default:"2048"`
	TaskDefaultTimeout  time.Duration `env:"TASK_DEFAULT_TIMEOUT" default:"1h"`
	TaskRunWaitInterval time.Duration `env:"TASK_RUN_WAIT_INTERVAL" default:"5s"`
	TaskQueueAlertTime  time.Duration `env:"TASK_QUEUE_ALERT_TIME" default:"10m"`

	// agent
	AgentAccessibleCacheTTL int64  `env:"AGENT_ACCESSIBLE_CACHE_TTL" default:"43200"` // 默认 12 小时
	AgentPreFetcherDestDir  string `env:"AGENT_PRE_FETCHER_DEST_DIR" default:"/opt/emptydir"`

	// build cache
	BuildCacheCleanJobCron string        `env:"BUILD_CACHE_CLEAN_JOB_CRON" default:"0 0 0 * * ?"`
	BuildCacheExpireIn     time.Duration `env:"BUILD_CACHE_EXPIRE_IN" default:"168h"`

	// bundle
	GittarAddr    string `env:"GITTAR_ADDR" required:"true"`
	OpenAPIAddr   string `env:"OPENAPI_ADDR" required:"true"`
	EventboxAddr  string `env:"EVENTBOX_ADDR" required:"true"`
	DiceHubAddr   string `env:"DICEHUB_ADDR" required:"true"`
	SchedulerAddr string `env:"SCHEDULER_ADDR" required:"true"`
	HepaAddr      string `env:"HEPA_ADDR" required:"true"`
	CollectorAddr string `env:"COLLECTOR_ADDR" required:"false"`

	// public url
	GittarPublicURL    string `env:"GITTAR_PUBLIC_URL" required:"true"`
	OpenAPIPublicURL   string `env:"OPENAPI_PUBLIC_URL" required:"true"`
	CollectorPublicURL string `env:"COLLECTOR_PUBLIC_URL" required:"false"`

	// oss/nfs storage
	PipelineStorageURL string `env:"PIPELINE_STORAGE_URL" required:"true"`

	// action type mapping
	ActionTypeMappingStr string `env:"ACTION_TYPE_MAPPING"` // git:git-checkout,dicehub:release
	ActionTypeMapping    map[string]string

	// 流水线自动清理策略，JSON 格式配置。
	// 1. 提供默认策略
	// 2. 自定义 source 级别策略只需填写需要覆盖的字段即可，未填写的自动使用默认值。
	// e.g.
	// {
	//    "default": {
	//        "maxFinishedStoreTime": "720h", # 已完成的 默认保留 30 天
	//        "minFinishedStoreCount": 100,   # 已完成的 至少保留 100 条
	//        "maxAnalyzedStoreTime": "24h",  # 未开始的 默认保留 1 天
	//        "minAnalyzedStoreCount": 10     # 未开始的 至少保留 10 条
	//    },
	//    "ops": {                            # ops 场景自定义策略
	//        "minFinishedStoreCount": 5      # ops 场景下，已完成的 至少保留 5 条
	//    }                                   # 其他未声明字段使用默认值
	//}
	PipelineAutoCleanupStrategy PipelineAutoCleanupStrategy `env:"PIPELINE_AUTO_CLEANUP_STRATEGY"`
	PipelineAutoCleanupJobCron  string                      `env:"PIPELINE_AUTO_CLEANUP_JOB_CRON" default:"0 30 5 * * ?"`

	// 默认用户 ID，用于鉴权
	InternalUserID string `env:"INTERNAL_USER_ID" default:"1103"`

	// cms
	CmsBase64EncodedRsaPublicKey  string `env:"CMS_BASE64_ENCODED_RSA_PUBLIC_KEY" default:"LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUFrOCtVK3QyeHhoM1hpREJnRjM2dApxWU5UZmN2NDA4aTdsZnFZRG9TRHMxbDA5bitsLzFOZTQ5b0xxZ0h1ZTQ5MmJHNFI0T0ZHZW1IMktIZmUya3BnCjZpd2tFM0xrZW5KMm56NFdPQWNnOUhiWlA0TFpReGxoeUVwNlE2aHQyekgxZ25Uc2p0QUlzMEZxbXJXZmlVVkQKdFdib1lmSDMvNWZReSs3V00yWkU3bzdnWWxIM1RLR2M5amEvWmgwOTBUZXdULzV3TVhPb1llcFRsWVBmTDVoTwo0em9GeGFpbzltanhpQmVveDNrUkM5RlZsSFM4ZDVlYWRHNkttR2cydjlTaE96SThDaGErRkJHSm83b3E4UEZEClRFMUFuZnBjZml5ckVxVVpzbDZTckl1TjVZUTREM3h1clZnY1RkcG9MV1dpallJbVZ0bytJU3FScW9QemxqVWQKTzdDa2NVRXUvVno2UCt2Vjc4b1JWRktYM0E0aG9vYlFFSkphNlFISmlzN1JQRW5TTjZXS2k4RXkzSlFhT3hXWAppejR3aDk3VmIyZDU4c3l1M0pJSTFOWVlyemtqTitEd1RLV1dqcjVYaVhHSGVCRDFtMmpaMytxV1RCTW1oNC9QCmtWc2M0T29lOG40ZXFoYVc1d2QyaU5jUlRHUS9sUmY4ekNSRlhCN1lvbWJrVlQwc1hVcllXQWFkWURFUEFmazUKTncvUjJaTXkyNGVhd0ZCcTVmYVB6VVJWRUY4WC9uUm5kL1YwUFZBSGgySG9CeFJaZzFkSGJrSWQ3SUo5R2cxbwpKVzJZOTlobzRpK0QvTDl2cWNPOVRyOXN0dStWcG1UQ1BRdFZqWHlpY0FuZmN4MWxhOEI0Q2Y4azhWN1RBSmJWCm14SjdaUTJEbGs3TTdBYzNTamVEUmJrQ0F3RUFBUT09Ci0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo="`
	CmsBase64EncodedRsaPrivateKey string `env:"CMS_BASE64_ENCODED_RSA_PRIVATE_KEY" default:"LS0tLS1CRUdJTiBwcml2YXRlIGtleS0tLS0tCk1JSUpLUUlCQUFLQ0FnRUFrOCtVK3QyeHhoM1hpREJnRjM2dHFZTlRmY3Y0MDhpN2xmcVlEb1NEczFsMDluK2wKLzFOZTQ5b0xxZ0h1ZTQ5MmJHNFI0T0ZHZW1IMktIZmUya3BnNml3a0UzTGtlbkoybno0V09BY2c5SGJaUDRMWgpReGxoeUVwNlE2aHQyekgxZ25Uc2p0QUlzMEZxbXJXZmlVVkR0V2JvWWZIMy81ZlF5KzdXTTJaRTdvN2dZbEgzClRLR2M5amEvWmgwOTBUZXdULzV3TVhPb1llcFRsWVBmTDVoTzR6b0Z4YWlvOW1qeGlCZW94M2tSQzlGVmxIUzgKZDVlYWRHNkttR2cydjlTaE96SThDaGErRkJHSm83b3E4UEZEVEUxQW5mcGNmaXlyRXFVWnNsNlNySXVONVlRNApEM3h1clZnY1RkcG9MV1dpallJbVZ0bytJU3FScW9QemxqVWRPN0NrY1VFdS9WejZQK3ZWNzhvUlZGS1gzQTRoCm9vYlFFSkphNlFISmlzN1JQRW5TTjZXS2k4RXkzSlFhT3hXWGl6NHdoOTdWYjJkNThzeXUzSklJMU5ZWXJ6a2oKTitEd1RLV1dqcjVYaVhHSGVCRDFtMmpaMytxV1RCTW1oNC9Qa1ZzYzRPb2U4bjRlcWhhVzV3ZDJpTmNSVEdRLwpsUmY4ekNSRlhCN1lvbWJrVlQwc1hVcllXQWFkWURFUEFmazVOdy9SMlpNeTI0ZWF3RkJxNWZhUHpVUlZFRjhYCi9uUm5kL1YwUFZBSGgySG9CeFJaZzFkSGJrSWQ3SUo5R2cxb0pXMlk5OWhvNGkrRC9MOXZxY085VHI5c3R1K1YKcG1UQ1BRdFZqWHlpY0FuZmN4MWxhOEI0Q2Y4azhWN1RBSmJWbXhKN1pRMkRsazdNN0FjM1NqZURSYmtDQXdFQQpBUUtDQWdCSlhxbngyS2ZNMHJWUTJjcG8veTJPeml4Y2Jpb21YaWFYTE52Ym9QV0t5aVhmMGI4QlBVNEZ4Zzh5CkpXRk9uZ2pIaTk5K0EvU3EvUU5tVlJJZXd2cldZbkRKNHFiOURPSks2MU8ySGZ2Q3ZWZmJTY1UwcEYzQVFRL3QKazZac1BxRkNUMjI0K2hUSGZmby9yMVh3bXB3Z2FHT0Rjc3VLYUw1dzdDNFJOM3VSK3dQd2FnVmFXWUtEU091Ngo4VnJsQmtLVGdwWUlSZ1BZRHF2TXRMZk5kVW43U3FyZzBYYUZVZFJLbkl2ZjcvMkJJempheHhOaVBiT2loZGh3CkRKTFlwK0FjZFRRT1FmbTZGblorK2dNa3RHMldhMlpleEk2eTV0TklId0hoWTBabE5hU0t3Qlhmd2dGaU5ERmcKaDhCY2dHMnUxbUxYaTk5NU1SczdTK0pXdGlpNjNqYUxsMmN0eXIxYVJPNlJRdmtYMlgzbTU4MFRKODVwZ0dBbgpYY3hENW1HNTRsRFlqdTllbVlRUUZlRDNrQXQzV3pGWmhkWXFtUEU4VEk1clBCbGtxNzkwVzF3K3BveUpHc0VoCjJIZXlMekQ5WGNSdXBoVE1QaEFXUTlBMzVHbnBOM0NrUXA0c1lBSnlNck5FQWhHQ285cXVGZW5rRDMxQmFPSjMKQWN1MDBISVpyUy9JTkRZNGJ4MmExd3RFbnd4dDkxSnJ4R1I5eFFtZE5DejFLZkR2Z2laSldTR2ZPYXhkbmtoMgp0WHkrYkl1WER0ZVhsK3dDN0wySG9PTmZKUjM5RDNPTjVLdzRvbTdlMjVSM3dKcEN2UGFFR01ELzZ0YWt6N1dCCnVEbmVBTHRSb1ZkSzhsK1RuaGVKcmtQRG1tUnkwL1BPRzgxWXp3VVlEVCs1VmtuTnNRS0NBUUVBd3paQ0JWTk4KaGlza3psTmdRTTRZTkVLT3lLUkcwL05LUlBIM1lzVUZyeDJEQjhESzVITWJCSVNqQ0ZrQUhJRXI1R1VlVVUwRgpCMndURFFScnQ0KzZYcG92aE5SZmdOQWZodXZDdXZNajk2eG05N3hpa01GeDI3ZW9VcVQ0amlKSU5NTG12MXh1CnY1UTNGbnhhQVgxdE05UHpEbHJQN3ZXaWsrWWprZFRZZlFLL0dHdU90d2RwdC96V2xXREtaSmp6RmozamtVNUoKK3lNcjVlbDBSdVJFQlZkaHdHTkRqWHRGNXg0UkQ2YWpXRXFtNEJ4VXBHeGM3UEZ2NU96d0NMNm50NkxTUGwzMwo0OGpOVVhsYWhESk0rUStQeHBBR2UxTXVUM2dTdmNaVk5hMG83eXNuZHE2b3VaSUx5VFZ3S0xoZG1kbnpZY1BWCmpKL0lFOTdZR0RBblZRS0NBUUVBd2RhbENoWDlacFVGalFOWG1KcnUyaGwxck9UYmE1V3FCWGh1ME91bnNHaG4KUkN2YjgyckQ4dnR4VVNLV1drSTd3cFFEakVxOVQyVmRRbmswR3ZVK3lIUW1WVE9MenNTcVlBRk0rdnUrWnJ0UApMcWxSa05lYll4UGx0TkwyVWJQYUMrOVpWblNOYUJDeEpOKzBwcTJnZG5IRWd5YWRJQ1FUVE41OGFDSVN4SGg5ClRqQ2N6Njd1RDMrU3FTVHVEZHFnQUhmeDlNTytjb3JQaW1Zc2c4d2FwZ3JtR0paMkpRWC9QdllLUlkrVXFoSngKd2VBaGlYVTFDY2Y3eWpMUEdWZURrZTQwTktEcGFNZjJnYW9qakgxei8yaERuU3A2enc4Wnc1VWg3dlA2azBkeQpvYzVzQ0FxYURWNkZEMG4xNWF5Q1RVWG1EQU1VTWU1Qlh5SjVCQkJjMVFLQ0FRRUFwNVcrMjkrRjREYk5wQ3RECnFKN0ZmS2ZlK0RTL2NWbWRXczcyOTkzNFlUdE9yNnM5QXg0bUJaendjVXdtb2xIcUltc0V1ZnNLNURKTnNKRXAKQUM3dGFpV253YnFvT21keGlWeUFrZ29GeUt4Q3dVOEN0dzY2OWtzV3Y4eE1iWWpVd0NiSi9XSVcyWFVlVGJsMwpjMndBQWN4bER0KzdQb08xakk2MzNvd0JSbURET08ydFdVZU41SnUwaEF6Ujg4YXllVmVzTTZRb011Y2cyb0d1Cmh1V1QxNW9LbXlVY2F5dDIrVkNBaVJVZmliNmN3Q3pTSlUyNkFOZk1uWlVqQS83WThQZGcwcFhOSjhuTktiS3EKbUc2dVVlcWdIWENyZjlnTEc4SVRKTVJOaG9VZmJTTjQvNVExMlFtZUFLQlZweitQYTNNR1U5blJUS1luRjVmcApuK3BHK1FLQ0FRRUF3UTMrb2VUMDFFNW5rT0piUStwTEtYMWg3aWloUUsxM0FLdko4dHBCMFRpcVlRTXR0V29JCmJ1QnZJOWZHMTI1UUJxTlVSVTNLN21DT1diNU5YdXdTODZKNjZ6RERkZFA1dkZTUFR3bWJ3TVdkUDJQemtNYXMKUkNsMUJudDJTRGxRV2NLd3Y2S2xrNWZNVm1WWGp3b3VYc2xBWno3MkR5VGU5QmhDMzVQUURVM1R2eVE3aWIwMwo3TWVxVWp3dHZDNmFYTjBaWmlYdWNEWkFMaDlGQnA4cGkyWWZkUzJsellvRGhibVcwV0VITjd2WEFMa3hyYTNHCmZVOW9QeUlMa2JuUG1IQWVIcXlFeTQ4Y3ZGZXZ3Q1RTZXZabElRdEY5U09kRFdaaXZaTFJaZzRxNVd5cHUvaVQKSmUyVnFIeUpJNDZFMkdGZGxXa2JtLzhuckpDdzVwTkZZUUtDQVFCWTNjKzVnY0dkcGJoYkRjenlHNngrcHJVbwpWclVlanNJTUlIZEh1MUUzWW5rWVdrL2hMWEh4anYrbW14RTVOb0hpMmJSTG53VXZhelhReW1CV2l2cW9FRGtmClZsMGR1Tk5xZmpWcG1JY092U1Rkb0wwRHlsZVExT2JEUjBZMEhMZmNvUnJLSUhzOE13b3BWWGtIZjFzamdxVEIKMjNycHd6aDVBTHdaMXJNVDJGNC9CL3RRdlBBMzRTeGlhanQ0RlBHU3pwUE50b0trU1ZOTnM3TS96SXFmY1JtZAppNjNhcW1Tc2pJWEdrZWY1cGpZZTV2WHM2ODZkVldMVGZ4SllMa0N6TkZ4c1dZdXBFaWhZQVJQbnViUGpNd21QCkVWcDBiQTdJNlpnQUQ3Vmw1TkR3UGg4bFZYUlYzWHVyV2dFVXUrcytqMHNTMEZyS2tUeldsRmF1aUhmawotLS0tLUVORCBwcml2YXRlIGtleS0tLS0tCg=="`

	// openapi oauth2 token client
	OpenapiOAuth2TokenClientID     string `env:"OPENAPI_OAUTH2_TOKEN_CLIENT_ID" default:"pipeline"`
	OpenapiOAuth2TokenClientSecret string `env:"OPENAPI_OAUTH2_TOKEN_CLIENT_SECRET" default:"devops/pipeline"`

	// DisablePipelineVolume default is false, means enable context volumes
	DisablePipelineVolume bool `env:"DISABLE_PIPELINE_VOLUME" default:"false"`
}

var cfg Conf

// Load 从环境变量加载配置选项.
func Load() {
	envconf.MustLoad(&cfg)

	// actionTypeMapping
	checkActionTypeMapping(&cfg)

	// pipelineAutoCleanupStrategy
	handlePipelineAutoCleanupStrategy(&cfg)
}

// ListenAddr 返回 pipeline 服务监听地址.
func ListenAddr() string {
	return cfg.ListenAddr
}

// Debug 返回 Debug 选项.
func Debug() bool {
	return cfg.Debug
}

// DiceCluster 返回 pipeline 服务所在集群，用于判断 task 运行在 中心 or SaaS 集群.
func DiceCluster() string {
	return cfg.DiceCluster
}

// TaskDefaultCPU 返回 task 默认的 cpu 限制.
func TaskDefaultCPU() float64 {
	return cfg.TaskDefaultCPU
}

// TaskDefaultMEM 返回 task 默认的 memory 限制.
func TaskDefaultMEM() float64 {
	return cfg.TaskDefaultMEM
}

// TaskDefaultTimeout 返回 task 默认超时时间.
func TaskDefaultTimeout() time.Duration {
	return cfg.TaskDefaultTimeout
}

// TaskRunWaitInterval 返回 task run 在 wait 过程中的时间间隔.
// 时间越短越精确，但对 scheduler 压力会变大，建议使用默认值 5s.
func TaskRunWaitInterval() time.Duration {
	return cfg.TaskRunWaitInterval
}

// TaskQueueAlertTime 返回 task 排队超过多少时间告警.
func TaskQueueAlertTime() time.Duration {
	return cfg.TaskQueueAlertTime
}

// AgentTTL 返回 agent 下载成功后 accessible 状态缓存的有效时间，默认为 12 小时。
func AgentAccessibleCacheTTL() int64 {
	return cfg.AgentAccessibleCacheTTL
}

// AgentContainerPathWhenExecute 返回 agent 在 action 运行时需要被挂载的路径, 例如: /opt/emptydir/agent .
func AgentContainerPathWhenExecute() string {
	return filepath.Join(cfg.AgentPreFetcherDestDir, "agent")
}

// AgentPreFetcherDestDir 返回 agent 在 PreFetcher 中需要被 entrypoint 拷贝到的目录.
// 注意：不能是 /opt/action(agent 在镜像中的目录为 /opt/action/agent)，
// 这样会导致 k8s initContainer emptyDir 与 mountDir 相同，entrypoint 执行失败。
func AgentPreFetcherDestDir() string {
	return cfg.AgentPreFetcherDestDir
}

// BuildCacheCleanJobCron 返回清理 构建缓存镜像 任务的 定时配置.
func BuildCacheCleanJobCron() string {
	return cfg.BuildCacheCleanJobCron
}

// BuildCacheExpireIn 返回 构建缓存镜像 的失效时间.
func BuildCacheExpireIn() time.Duration {
	return cfg.BuildCacheExpireIn
}

// GittarAddr 返回 gittar 的集群内部地址.
func GittarAddr() string {
	return cfg.GittarAddr
}

// OpenAPIAddr 返回 openapi 的集群内部地址.
func OpenAPIAddr() string {
	return cfg.OpenAPIAddr
}

// HepaAddr 返回 hepa 的集群内部地址.
func HepaAddr() string {
	return cfg.HepaAddr
}

// EventboxAddr 返回 eventbox 的集群内部地址.
func EventboxAddr() string {
	return cfg.EventboxAddr
}

// DiceHubAddr 返回 dicehub 的集群内部地址.
func DiceHubAddr() string {
	return cfg.DiceHubAddr
}

// SchedulerAddr 返回 scheduler 的集群内部地址.
func SchedulerAddr() string {
	return cfg.SchedulerAddr
}

// CollectorAddr 返回 collector 的集群内部地址.
func CollectorAddr() string {
	return cfg.CollectorAddr
}

// GittarPublicURL 返回 gittar 的公网地址.
func GittarPublicURL() string {
	return cfg.GittarPublicURL
}

// OpenAPIPublicURL 返回 openapi 的公网地址，用于 SaaS 集群 task 回调中心集群 openapi.
func OpenAPIPublicURL() string {
	return cfg.OpenAPIPublicURL
}

// CollectorPublicURL 返回 collector 的公网地址
func CollectorPublicURL() string {
	return cfg.CollectorPublicURL
}

// StorageURL 返回 storage url.
func StorageURL() string {
	return cfg.PipelineStorageURL
}

// ActionTypeMapping 返回新老 action type 映射关系.
func ActionTypeMapping() map[string]string {
	return cfg.ActionTypeMapping
}

// GetPipelineAutoCleanupSourceStrategy 获取 source 级别流水线自动清理策略
func GetPipelineAutoCleanupSourceStrategy(source apistructs.PipelineSource) PipelineAutoCleanupSourceStrategyItem {
	return cfg.PipelineAutoCleanupStrategy.GetSourceStrategy(source)
}

// PipelineAutoCleanupJobCron 返回 流水线自动清理 定时任务 的定时设置.
func PipelineAutoCleanupJobCron() string {
	return cfg.PipelineAutoCleanupJobCron
}

// InternalUserID 返回 pipeline 组件在内部调用时默认分配的 用户 ID
func InternalUserID() string {
	return cfg.InternalUserID
}

// CmsBase64EncodedRsaPublicKey 返回 配置管理 用于加密 value 的 rsa public key.
func CmsBase64EncodedRsaPublicKey() string {
	return cfg.CmsBase64EncodedRsaPublicKey
}

// CmsBase64EncodedRsaPrivateKey 返回 配置管理 用于解密 value 的 rsa private key.
func CmsBase64EncodedRsaPrivateKey() string {
	return cfg.CmsBase64EncodedRsaPrivateKey
}

// OpenapiOAuth2TokenClientID 返回 用于申请 openapi oauth2 token 的客户端 id.
func OpenapiOAuth2TokenClientID() string {
	return cfg.OpenapiOAuth2TokenClientID
}

// OpenapiOAuth2TokenClientID 返回 用于申请 openapi oauth2 token 的客户端 secret.
func OpenapiOAuth2TokenClientSecret() string {
	return cfg.OpenapiOAuth2TokenClientSecret
}

// DisablePipelineVolume 返回 是否关闭 pipeline volume，只有值引用.
func DisablePipelineVolume() bool {
	return cfg.DisablePipelineVolume
}
