package enum

type BasicInformation string

func (b BasicInformation) String() string {
	return string(b)
}

const (
	RequestID    BasicInformation = "request_id"
	AuthModel    BasicInformation = "authModel"
	AuthMqlModel BasicInformation = "authMqlModel"
)

type TaskStatus string

const (
	TaskStatusFailed  TaskStatus = "failed"  //　失败的
	TaskStatusWait    TaskStatus = "wait"    // 等待执行
	TaskStatusRunning TaskStatus = "running" // 运行中
	TaskStatusPassed  TaskStatus = "passed"  // 通过
)

type TaskType string

const (
	TaskTypeDeploy TaskType = "deploy" //　部署任务
	TaskTypeStop   TaskType = "stop"   //　停止任务
	TaskTypeReboot TaskType = "reboot" //　重启任务
)

type TaskStage string

const (
	TaskStageBuild  TaskStage = "build"
	TaskStageTest   TaskStage = "test"
	TaskStageDeploy TaskStage = "deploy"
	TaskStageStop   TaskStage = "stop"
	TaskStageReboot TaskStage = "reboot"
)

type TaskAction string

const (
	TaskActionPush        TaskAction = "push"
	TaskActionTag         TaskAction = "tag_push"
	TaskActionMerge       TaskAction = "merge_request"
	TaskActionPushOrMerge TaskAction = "push_or_merge"
)
