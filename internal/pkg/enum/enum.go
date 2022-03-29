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

type TaskStage string

const (
	TaskStageBuild  TaskStage = "build"
	TaskStageTest   TaskStage = "test"
	TaskStageDeploy TaskStage = "deploy"
)

type TaskAction string

const (
	TaskActionPush  TaskAction = "push"
	TaskActionTag   TaskStage  = "tag_push"
	TaskActionMerge TaskStage  = "merge_request"
)
