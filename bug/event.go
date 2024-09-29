package bug

type Event string

// 定义缺陷事件
const (
	EventBugNew        = Event("创建缺陷")
	EventBugInProgress = Event("接收/处理缺陷")
	EventBugResolved   = Event("已解决缺陷")
	EventBugReopened   = Event("重新打开缺陷")
	EventBugRejected   = Event("拒绝缺陷")
	EventBugClosed     = Event("关闭缺陷")
	EventBugSuspended  = Event("挂起缺陷")
)

// 定义缺陷事件对应的处理方法
var eventHandler = map[Event]Handler{
	EventBugNew:        handlerBugNew,
	EventBugInProgress: handlerBugInProgress,
	EventBugResolved:   handlerBugResolved,
	EventBugReopened:   handlerBugReopened,
	EventBugRejected:   handlerBugRejected,
	EventBugClosed:     handlerBugClosed,
	EventBugSuspended:  handlerBugSuspended,
}

type EventDesc struct {
	Name      Event     // 事件名称
	Src       State     // 初始状态
	Dst       State     // 目的状态
	WorkItem  WorkItem  // 工作项
	Condition Condition // 事件触发条件
}
