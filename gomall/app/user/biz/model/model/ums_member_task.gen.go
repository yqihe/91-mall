// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUmsMemberTask = "ums_member_task"

// UmsMemberTask 会员任务表
type UmsMemberTask struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string `gorm:"column:name" json:"name"`
	Growth       int32  `gorm:"column:growth;comment:赠送成长值" json:"growth"`            // 赠送成长值
	Intergration int32  `gorm:"column:intergration;comment:赠送积分" json:"intergration"` // 赠送积分
	Type         int32  `gorm:"column:type;comment:任务类型：0->新手任务；1->日常任务" json:"type"` // 任务类型：0->新手任务；1->日常任务
}

// TableName UmsMemberTask's table name
func (*UmsMemberTask) TableName() string {
	return TableNameUmsMemberTask
}
