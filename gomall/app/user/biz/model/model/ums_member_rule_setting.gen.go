// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUmsMemberRuleSetting = "ums_member_rule_setting"

// UmsMemberRuleSetting 会员积分成长规则表
type UmsMemberRuleSetting struct {
	ID                int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ContinueSignDay   int32   `gorm:"column:continue_sign_day;comment:连续签到天数" json:"continue_sign_day"`         // 连续签到天数
	ContinueSignPoint int32   `gorm:"column:continue_sign_point;comment:连续签到赠送数量" json:"continue_sign_point"`   // 连续签到赠送数量
	ConsumePerPoint   float64 `gorm:"column:consume_per_point;comment:每消费多少元获取1个点" json:"consume_per_point"`    // 每消费多少元获取1个点
	LowOrderAmount    float64 `gorm:"column:low_order_amount;comment:最低获取点数的订单金额" json:"low_order_amount"`      // 最低获取点数的订单金额
	MaxPointPerOrder  int32   `gorm:"column:max_point_per_order;comment:每笔订单最高获取点数" json:"max_point_per_order"` // 每笔订单最高获取点数
	Type              int32   `gorm:"column:type;comment:类型：0->积分规则；1->成长值规则" json:"type"`                      // 类型：0->积分规则；1->成长值规则
}

// TableName UmsMemberRuleSetting's table name
func (*UmsMemberRuleSetting) TableName() string {
	return TableNameUmsMemberRuleSetting
}
