package main

import (
	"fmt"
	"github.com/mark8s/go-fsm/bug"
)

func main() {
	fsmEngine, err := bug.NewFSM()
	if err != nil {
		fmt.Printf("FSM实例化失败：%v", err)
		return
	}

	err = fsmEngine.Event(bug.Args{
		WorkItem: bug.Bug,
		Src:      bug.StatusBugInProgress,
		Dst:      bug.StatusBugResolved,
		Fields: []bug.FieldArgs{
			{Field: "reason", FiledType: bug.FieldTypeString, Value: "后端bug"},
			{Field: "impact", FiledType: bug.FieldTypeString, Value: "仅当前界面"},
		},
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func sendDingTalk(method, content string) error {
	fmt.Printf("发送钉钉通知")
	return nil
}
