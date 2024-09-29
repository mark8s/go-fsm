package bug

// Condition 条件
type Condition struct {
	Fields    []FieldCondition    `json:"fields"`    // 字段条件
	TimeSheet *TimeSheetCondition `json:"timeSheet"` // 工时花费
	Commend   *CommendCondition   `json:"commend"`   // 评论
}

type FieldCondition struct {
	Field     string      `json:"field"`     // 字段
	FiledType string      `json:"filedType"` // 字段类型 number、string、bool ......
	Default   interface{} `json:"default"`   // 默认值
	Required  bool        `json:"required"`  // 是否必填
}

type TimeSheetCondition struct {
	Required bool `json:"required"` // 是否必填
}

type CommendCondition struct {
	Required bool   `json:"required"` // 是否必填
	Default  string `json:"default"`  // 默认值
}

type Args struct {
	Src      State       `json:"src"`      // 当前状态
	Dst      State       `json:"dst"`      // 目标状态
	WorkItem WorkItem    `json:"workItem"` // 工作项
	Fields   []FieldArgs `json:"fields"`   // 字段
}

type FieldArgs struct {
	Field     string      `json:"field"`     // 字段
	FiledType string      `json:"filedType"` // 字段类型 number、string、bool ......
	Value     interface{} `json:"value"`     // 值
}

const (
	FieldTypeString = "string"
	FieldTypeNumber = "number"
	FieldTypeBool   = "bool"
)
