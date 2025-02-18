// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUmsMemberLevel = "ums_member_level"

// UmsMemberLevel 会员等级表
type UmsMemberLevel struct {
	ID                    int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name                  string  `gorm:"column:name" json:"name"`
	GrowthPoint           int32   `gorm:"column:growth_point" json:"growth_point"`
	DefaultStatus         int32   `gorm:"column:default_status;comment:是否为默认等级：0->不是；1->是" json:"default_status"`          // 是否为默认等级：0->不是；1->是
	FreeFreightPoint      float64 `gorm:"column:free_freight_point;comment:免运费标准" json:"free_freight_point"`               // 免运费标准
	CommentGrowthPoint    int32   `gorm:"column:comment_growth_point;comment:每次评价获取的成长值" json:"comment_growth_point"`      // 每次评价获取的成长值
	PriviledgeFreeFreight int32   `gorm:"column:priviledge_free_freight;comment:是否有免邮特权" json:"priviledge_free_freight"`   // 是否有免邮特权
	PriviledgeSignIn      int32   `gorm:"column:priviledge_sign_in;comment:是否有签到特权" json:"priviledge_sign_in"`             // 是否有签到特权
	PriviledgeComment     int32   `gorm:"column:priviledge_comment;comment:是否有评论获奖励特权" json:"priviledge_comment"`          // 是否有评论获奖励特权
	PriviledgePromotion   int32   `gorm:"column:priviledge_promotion;comment:是否有专享活动特权" json:"priviledge_promotion"`       // 是否有专享活动特权
	PriviledgeMemberPrice int32   `gorm:"column:priviledge_member_price;comment:是否有会员价格特权" json:"priviledge_member_price"` // 是否有会员价格特权
	PriviledgeBirthday    int32   `gorm:"column:priviledge_birthday;comment:是否有生日特权" json:"priviledge_birthday"`           // 是否有生日特权
	Note                  string  `gorm:"column:note" json:"note"`
}

// TableName UmsMemberLevel's table name
func (*UmsMemberLevel) TableName() string {
	return TableNameUmsMemberLevel
}
