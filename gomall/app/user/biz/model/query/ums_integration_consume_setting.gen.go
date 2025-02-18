// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/yqihe/91-mall/gomall/app/user/biz/model/model"
)

func newUmsIntegrationConsumeSetting(db *gorm.DB, opts ...gen.DOOption) umsIntegrationConsumeSetting {
	_umsIntegrationConsumeSetting := umsIntegrationConsumeSetting{}

	_umsIntegrationConsumeSetting.umsIntegrationConsumeSettingDo.UseDB(db, opts...)
	_umsIntegrationConsumeSetting.umsIntegrationConsumeSettingDo.UseModel(&model.UmsIntegrationConsumeSetting{})

	tableName := _umsIntegrationConsumeSetting.umsIntegrationConsumeSettingDo.TableName()
	_umsIntegrationConsumeSetting.ALL = field.NewAsterisk(tableName)
	_umsIntegrationConsumeSetting.ID = field.NewInt64(tableName, "id")
	_umsIntegrationConsumeSetting.DeductionPerAmount = field.NewInt32(tableName, "deduction_per_amount")
	_umsIntegrationConsumeSetting.MaxPercentPerOrder = field.NewInt32(tableName, "max_percent_per_order")
	_umsIntegrationConsumeSetting.UseUnit = field.NewInt32(tableName, "use_unit")
	_umsIntegrationConsumeSetting.CouponStatus = field.NewInt32(tableName, "coupon_status")

	_umsIntegrationConsumeSetting.fillFieldMap()

	return _umsIntegrationConsumeSetting
}

// umsIntegrationConsumeSetting 积分消费设置
type umsIntegrationConsumeSetting struct {
	umsIntegrationConsumeSettingDo umsIntegrationConsumeSettingDo

	ALL                field.Asterisk
	ID                 field.Int64
	DeductionPerAmount field.Int32 // 每一元需要抵扣的积分数量
	MaxPercentPerOrder field.Int32 // 每笔订单最高抵用百分比
	UseUnit            field.Int32 // 每次使用积分最小单位100
	CouponStatus       field.Int32 // 是否可以和优惠券同用；0->不可以；1->可以

	fieldMap map[string]field.Expr
}

func (u umsIntegrationConsumeSetting) Table(newTableName string) *umsIntegrationConsumeSetting {
	u.umsIntegrationConsumeSettingDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsIntegrationConsumeSetting) As(alias string) *umsIntegrationConsumeSetting {
	u.umsIntegrationConsumeSettingDo.DO = *(u.umsIntegrationConsumeSettingDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsIntegrationConsumeSetting) updateTableName(table string) *umsIntegrationConsumeSetting {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.DeductionPerAmount = field.NewInt32(table, "deduction_per_amount")
	u.MaxPercentPerOrder = field.NewInt32(table, "max_percent_per_order")
	u.UseUnit = field.NewInt32(table, "use_unit")
	u.CouponStatus = field.NewInt32(table, "coupon_status")

	u.fillFieldMap()

	return u
}

func (u *umsIntegrationConsumeSetting) WithContext(ctx context.Context) IUmsIntegrationConsumeSettingDo {
	return u.umsIntegrationConsumeSettingDo.WithContext(ctx)
}

func (u umsIntegrationConsumeSetting) TableName() string {
	return u.umsIntegrationConsumeSettingDo.TableName()
}

func (u umsIntegrationConsumeSetting) Alias() string { return u.umsIntegrationConsumeSettingDo.Alias() }

func (u umsIntegrationConsumeSetting) Columns(cols ...field.Expr) gen.Columns {
	return u.umsIntegrationConsumeSettingDo.Columns(cols...)
}

func (u *umsIntegrationConsumeSetting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsIntegrationConsumeSetting) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 5)
	u.fieldMap["id"] = u.ID
	u.fieldMap["deduction_per_amount"] = u.DeductionPerAmount
	u.fieldMap["max_percent_per_order"] = u.MaxPercentPerOrder
	u.fieldMap["use_unit"] = u.UseUnit
	u.fieldMap["coupon_status"] = u.CouponStatus
}

func (u umsIntegrationConsumeSetting) clone(db *gorm.DB) umsIntegrationConsumeSetting {
	u.umsIntegrationConsumeSettingDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsIntegrationConsumeSetting) replaceDB(db *gorm.DB) umsIntegrationConsumeSetting {
	u.umsIntegrationConsumeSettingDo.ReplaceDB(db)
	return u
}

type umsIntegrationConsumeSettingDo struct{ gen.DO }

type IUmsIntegrationConsumeSettingDo interface {
	gen.SubQuery
	Debug() IUmsIntegrationConsumeSettingDo
	WithContext(ctx context.Context) IUmsIntegrationConsumeSettingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsIntegrationConsumeSettingDo
	WriteDB() IUmsIntegrationConsumeSettingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsIntegrationConsumeSettingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsIntegrationConsumeSettingDo
	Not(conds ...gen.Condition) IUmsIntegrationConsumeSettingDo
	Or(conds ...gen.Condition) IUmsIntegrationConsumeSettingDo
	Select(conds ...field.Expr) IUmsIntegrationConsumeSettingDo
	Where(conds ...gen.Condition) IUmsIntegrationConsumeSettingDo
	Order(conds ...field.Expr) IUmsIntegrationConsumeSettingDo
	Distinct(cols ...field.Expr) IUmsIntegrationConsumeSettingDo
	Omit(cols ...field.Expr) IUmsIntegrationConsumeSettingDo
	Join(table schema.Tabler, on ...field.Expr) IUmsIntegrationConsumeSettingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationConsumeSettingDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationConsumeSettingDo
	Group(cols ...field.Expr) IUmsIntegrationConsumeSettingDo
	Having(conds ...gen.Condition) IUmsIntegrationConsumeSettingDo
	Limit(limit int) IUmsIntegrationConsumeSettingDo
	Offset(offset int) IUmsIntegrationConsumeSettingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsIntegrationConsumeSettingDo
	Unscoped() IUmsIntegrationConsumeSettingDo
	Create(values ...*model.UmsIntegrationConsumeSetting) error
	CreateInBatches(values []*model.UmsIntegrationConsumeSetting, batchSize int) error
	Save(values ...*model.UmsIntegrationConsumeSetting) error
	First() (*model.UmsIntegrationConsumeSetting, error)
	Take() (*model.UmsIntegrationConsumeSetting, error)
	Last() (*model.UmsIntegrationConsumeSetting, error)
	Find() ([]*model.UmsIntegrationConsumeSetting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsIntegrationConsumeSetting, err error)
	FindInBatches(result *[]*model.UmsIntegrationConsumeSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsIntegrationConsumeSetting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsIntegrationConsumeSettingDo
	Assign(attrs ...field.AssignExpr) IUmsIntegrationConsumeSettingDo
	Joins(fields ...field.RelationField) IUmsIntegrationConsumeSettingDo
	Preload(fields ...field.RelationField) IUmsIntegrationConsumeSettingDo
	FirstOrInit() (*model.UmsIntegrationConsumeSetting, error)
	FirstOrCreate() (*model.UmsIntegrationConsumeSetting, error)
	FindByPage(offset int, limit int) (result []*model.UmsIntegrationConsumeSetting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsIntegrationConsumeSettingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsIntegrationConsumeSettingDo) Debug() IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Debug())
}

func (u umsIntegrationConsumeSettingDo) WithContext(ctx context.Context) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsIntegrationConsumeSettingDo) ReadDB() IUmsIntegrationConsumeSettingDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsIntegrationConsumeSettingDo) WriteDB() IUmsIntegrationConsumeSettingDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsIntegrationConsumeSettingDo) Session(config *gorm.Session) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsIntegrationConsumeSettingDo) Clauses(conds ...clause.Expression) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsIntegrationConsumeSettingDo) Returning(value interface{}, columns ...string) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsIntegrationConsumeSettingDo) Not(conds ...gen.Condition) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsIntegrationConsumeSettingDo) Or(conds ...gen.Condition) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsIntegrationConsumeSettingDo) Select(conds ...field.Expr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsIntegrationConsumeSettingDo) Where(conds ...gen.Condition) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsIntegrationConsumeSettingDo) Order(conds ...field.Expr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsIntegrationConsumeSettingDo) Distinct(cols ...field.Expr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsIntegrationConsumeSettingDo) Omit(cols ...field.Expr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsIntegrationConsumeSettingDo) Join(table schema.Tabler, on ...field.Expr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsIntegrationConsumeSettingDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsIntegrationConsumeSettingDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsIntegrationConsumeSettingDo) Group(cols ...field.Expr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsIntegrationConsumeSettingDo) Having(conds ...gen.Condition) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsIntegrationConsumeSettingDo) Limit(limit int) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsIntegrationConsumeSettingDo) Offset(offset int) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsIntegrationConsumeSettingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsIntegrationConsumeSettingDo) Unscoped() IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsIntegrationConsumeSettingDo) Create(values ...*model.UmsIntegrationConsumeSetting) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsIntegrationConsumeSettingDo) CreateInBatches(values []*model.UmsIntegrationConsumeSetting, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsIntegrationConsumeSettingDo) Save(values ...*model.UmsIntegrationConsumeSetting) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsIntegrationConsumeSettingDo) First() (*model.UmsIntegrationConsumeSetting, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationConsumeSetting), nil
	}
}

func (u umsIntegrationConsumeSettingDo) Take() (*model.UmsIntegrationConsumeSetting, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationConsumeSetting), nil
	}
}

func (u umsIntegrationConsumeSettingDo) Last() (*model.UmsIntegrationConsumeSetting, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationConsumeSetting), nil
	}
}

func (u umsIntegrationConsumeSettingDo) Find() ([]*model.UmsIntegrationConsumeSetting, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsIntegrationConsumeSetting), err
}

func (u umsIntegrationConsumeSettingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsIntegrationConsumeSetting, err error) {
	buf := make([]*model.UmsIntegrationConsumeSetting, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsIntegrationConsumeSettingDo) FindInBatches(result *[]*model.UmsIntegrationConsumeSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsIntegrationConsumeSettingDo) Attrs(attrs ...field.AssignExpr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsIntegrationConsumeSettingDo) Assign(attrs ...field.AssignExpr) IUmsIntegrationConsumeSettingDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsIntegrationConsumeSettingDo) Joins(fields ...field.RelationField) IUmsIntegrationConsumeSettingDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsIntegrationConsumeSettingDo) Preload(fields ...field.RelationField) IUmsIntegrationConsumeSettingDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsIntegrationConsumeSettingDo) FirstOrInit() (*model.UmsIntegrationConsumeSetting, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationConsumeSetting), nil
	}
}

func (u umsIntegrationConsumeSettingDo) FirstOrCreate() (*model.UmsIntegrationConsumeSetting, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationConsumeSetting), nil
	}
}

func (u umsIntegrationConsumeSettingDo) FindByPage(offset int, limit int) (result []*model.UmsIntegrationConsumeSetting, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u umsIntegrationConsumeSettingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsIntegrationConsumeSettingDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsIntegrationConsumeSettingDo) Delete(models ...*model.UmsIntegrationConsumeSetting) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsIntegrationConsumeSettingDo) withDO(do gen.Dao) *umsIntegrationConsumeSettingDo {
	u.DO = *do.(*gen.DO)
	return u
}
