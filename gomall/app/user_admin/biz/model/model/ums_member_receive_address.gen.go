// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUmsMemberReceiveAddress = "ums_member_receive_address"

// UmsMemberReceiveAddress 会员收货地址表
type UmsMemberReceiveAddress struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MemberID      int64  `gorm:"column:member_id" json:"member_id"`
	Name          string `gorm:"column:name;comment:收货人名称" json:"name"` // 收货人名称
	PhoneNumber   string `gorm:"column:phone_number" json:"phone_number"`
	DefaultStatus int32  `gorm:"column:default_status;comment:是否为默认" json:"default_status"`    // 是否为默认
	PostCode      string `gorm:"column:post_code;comment:邮政编码" json:"post_code"`               // 邮政编码
	Province      string `gorm:"column:province;comment:省份/直辖市" json:"province"`               // 省份/直辖市
	City          string `gorm:"column:city;comment:城市" json:"city"`                           // 城市
	Region        string `gorm:"column:region;comment:区" json:"region"`                        // 区
	DetailAddress string `gorm:"column:detail_address;comment:详细地址(街道)" json:"detail_address"` // 详细地址(街道)
}

// TableName UmsMemberReceiveAddress's table name
func (*UmsMemberReceiveAddress) TableName() string {
	return TableNameUmsMemberReceiveAddress
}
