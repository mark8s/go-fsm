package bug

// workItemText 定义工作项文案
var workItemText = map[WorkItem]string{
	Bug:   "缺陷",
	Story: "需求",
	Task:  "任务",
}

// 工作项类型
const (
	Bug   = WorkItem("bug")
	Story = WorkItem("story")
	Task  = WorkItem("task")
)
