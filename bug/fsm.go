package bug

import (
	"errors"
	"fmt"
)

type WorkItem string               // 工作项
type State string                  // 状态
type Handler func(arg *Args) error // 处理方法，并返回新的状态

// FSM 有限状态机
type FSM struct {
	events []EventDesc
}

// NewFSM 实例化 FSM
func NewFSM() (fsm *FSM, err error) {
	fsm = new(FSM)
	fsm.events = []EventDesc{
		{
			Name:     EventBugInProgress,
			WorkItem: Bug,
			Src:      StatusBugNew, // new -> in_progress
			Dst:      StatusBugInProgress,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "developer", FiledType: FieldTypeString, Required: false},
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugRejected,
			WorkItem: Bug,
			Src:      StatusBugNew, // new -> rejected
			Dst:      StatusBugRejected,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "developer", FiledType: FieldTypeString, Required: false},
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugResolved,
			WorkItem: Bug,
			Src:      StatusBugInProgress, // in_progress -> resolved
			Dst:      StatusBugResolved,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "developer", FiledType: FieldTypeString, Required: false},
					{Field: "handler", FiledType: FieldTypeString, Required: false},
					{Field: "reason", FiledType: FieldTypeString, Required: true},
					{Field: "impact", FiledType: FieldTypeString, Required: true},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugRejected,
			WorkItem: Bug,
			Src:      StatusBugInProgress, // in_progress -> rejected
			Dst:      StatusBugRejected,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "developer", FiledType: FieldTypeString, Required: false},
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugSuspended,
			WorkItem: Bug,
			Src:      StatusBugInProgress, // in_progress -> suspended
			Dst:      StatusBugSuspended,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugReopened,
			WorkItem: Bug,
			Src:      StatusBugResolved, // resolved -> reopened
			Dst:      StatusBugReopened,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugClosed,
			WorkItem: Bug,
			Src:      StatusBugResolved, // resolved -> closed
			Dst:      StatusBugClosed,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugSuspended,
			WorkItem: Bug,
			Src:      StatusBugResolved, // resolved -> suspended
			Dst:      StatusBugSuspended,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugInProgress,
			WorkItem: Bug,
			Src:      StatusBugReopened, // reopened -> in_progress
			Dst:      StatusBugInProgress,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "developer", FiledType: FieldTypeString, Required: false},
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		}, {
			Name:     EventBugRejected,
			WorkItem: Bug,
			Src:      StatusBugReopened, // reopened -> rejected
			Dst:      StatusBugRejected,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "developer", FiledType: FieldTypeString, Required: false},
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		}, {
			Name:     EventBugSuspended,
			WorkItem: Bug,
			Src:      StatusBugReopened, // reopened -> suspended
			Dst:      StatusBugSuspended,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		}, {
			Name:     EventBugReopened,
			WorkItem: Bug,
			Src:      StatusBugRejected, // rejected -> reopened
			Dst:      StatusBugReopened,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		}, {
			Name:     EventBugClosed,
			WorkItem: Bug,
			Src:      StatusBugRejected, // rejected -> closed
			Dst:      StatusBugClosed,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugSuspended,
			WorkItem: Bug,
			Src:      StatusBugRejected, // rejected -> suspended
			Dst:      StatusBugSuspended,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugSuspended,
			WorkItem: Bug,
			Src:      StatusBugRejected, // closed -> reopened
			Dst:      StatusBugSuspended,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugResolved,
			WorkItem: Bug,
			Src:      StatusBugSuspended, // suspended -> resolved
			Dst:      StatusBugResolved,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugReopened,
			WorkItem: Bug,
			Src:      StatusBugSuspended, // suspended -> reopened
			Dst:      StatusBugReopened,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
		{
			Name:     EventBugClosed,
			WorkItem: Bug,
			Src:      StatusBugSuspended, // suspended -> closed
			Dst:      StatusBugClosed,
			Condition: Condition{
				Fields: []FieldCondition{
					{Field: "handler", FiledType: FieldTypeString, Required: false},
				},
				TimeSheet: &TimeSheetCondition{Required: false},
				Commend:   &CommendCondition{Required: false}},
		},
	}

	return
}

func (f *FSM) Event(args Args) error {

	var isExist bool
	var condition Condition
	var event EventDesc
	for _, e := range f.events {
		if e.Src == args.Src && e.Dst == args.Dst {
			isExist = true
			event = e
		}
	}

	if !isExist {
		return errors.New(fmt.Sprintf("[%s]状态不能从[%s]流转到[%s]", WorkItemText(args.WorkItem), StatusText(args.Src), StatusText(args.Dst)))
	}

	condition = event.Condition
	if condition.Fields != nil {
		for _, field := range condition.Fields {
			if field.Required {
				var isFieldExist bool
				var existFiledArgs FieldArgs
				for _, arg := range args.Fields {
					if field.Field == arg.Field {
						isFieldExist = true
						existFiledArgs = arg
					}
				}

				if !isFieldExist {
					return errors.New(fmt.Sprintf("[%s]状态从[%s]流转到[%s]时，字段%s不能为空", WorkItemText(args.WorkItem), StatusText(args.Src), StatusText(args.Dst), field.Field))
				}

				if field.FiledType == FieldTypeString {
					if existFiledArgs.Value == "" {
						return errors.New(fmt.Sprintf("[%s]状态从[%s]流转到[%s]时，字段%s不能为空", WorkItemText(args.WorkItem), StatusText(args.Src), StatusText(args.Dst), field.Field))
					}
				}
			}
		}
	}

	fmt.Printf("[%s]状态从[%s]流转到[%s]前置校验成功\n", WorkItemText(args.WorkItem), StatusText(args.Src), StatusText(args.Dst))

	fn, ok := eventHandler[event.Name]
	if !ok {
		return errors.New(fmt.Sprintf("[警告] 事件(%s)未配置任何处理器", event.Name))
	}

	err := fn(&args)
	if err != nil {
		return err
	}

	return nil
}
