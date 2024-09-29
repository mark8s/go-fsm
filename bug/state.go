package bug

// 定义缺陷状态
const (
	StatusBugNew        = State("new")
	StatusBugInProgress = State("in_progress")
	StatusBugResolved   = State("resolved")
	StatusBugReopened   = State("reopened")
	StatusBugRejected   = State("rejected")
	StatusBugClosed     = State("closed")
	StatusBugSuspended  = State("suspended")
)

// statusText 定义缺陷状态文案
var statusText = map[State]string{
	StatusBugNew:        "新",
	StatusBugInProgress: "接收/处理",
	StatusBugResolved:   "已解决",
	StatusBugReopened:   "重新打开",
	StatusBugRejected:   "拒绝",
	StatusBugClosed:     "关闭",
	StatusBugSuspended:  "挂起",
}

func StatusText(status State) string {
	return statusText[status]
}

func WorkItemText(item WorkItem) string {
	return workItemText[item]
}

// statusEvent 定义缺陷状态对应的可操作事件
var statusEvent = map[State][]Event{
	StatusBugNew:        {EventBugInProgress, EventBugRejected},
	StatusBugInProgress: {EventBugResolved, EventBugRejected, EventBugSuspended},
	StatusBugResolved:   {EventBugReopened, EventBugClosed, EventBugSuspended},
	StatusBugReopened:   {EventBugInProgress, EventBugRejected, EventBugSuspended},
	StatusBugRejected:   {EventBugReopened, EventBugClosed, EventBugSuspended},
	StatusBugClosed:     {EventBugReopened},
	StatusBugSuspended:  {EventBugResolved, EventBugReopened, EventBugClosed},
}
