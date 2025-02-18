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

func newUmsResourceCategory(db *gorm.DB, opts ...gen.DOOption) umsResourceCategory {
	_umsResourceCategory := umsResourceCategory{}

	_umsResourceCategory.umsResourceCategoryDo.UseDB(db, opts...)
	_umsResourceCategory.umsResourceCategoryDo.UseModel(&model.UmsResourceCategory{})

	tableName := _umsResourceCategory.umsResourceCategoryDo.TableName()
	_umsResourceCategory.ALL = field.NewAsterisk(tableName)
	_umsResourceCategory.ID = field.NewInt64(tableName, "id")
	_umsResourceCategory.CreateTime = field.NewTime(tableName, "create_time")
	_umsResourceCategory.Name = field.NewString(tableName, "name")
	_umsResourceCategory.Sort = field.NewInt32(tableName, "sort")

	_umsResourceCategory.fillFieldMap()

	return _umsResourceCategory
}

// umsResourceCategory 资源分类表
type umsResourceCategory struct {
	umsResourceCategoryDo umsResourceCategoryDo

	ALL        field.Asterisk
	ID         field.Int64
	CreateTime field.Time   // 创建时间
	Name       field.String // 分类名称
	Sort       field.Int32  // 排序

	fieldMap map[string]field.Expr
}

func (u umsResourceCategory) Table(newTableName string) *umsResourceCategory {
	u.umsResourceCategoryDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsResourceCategory) As(alias string) *umsResourceCategory {
	u.umsResourceCategoryDo.DO = *(u.umsResourceCategoryDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsResourceCategory) updateTableName(table string) *umsResourceCategory {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.CreateTime = field.NewTime(table, "create_time")
	u.Name = field.NewString(table, "name")
	u.Sort = field.NewInt32(table, "sort")

	u.fillFieldMap()

	return u
}

func (u *umsResourceCategory) WithContext(ctx context.Context) IUmsResourceCategoryDo {
	return u.umsResourceCategoryDo.WithContext(ctx)
}

func (u umsResourceCategory) TableName() string { return u.umsResourceCategoryDo.TableName() }

func (u umsResourceCategory) Alias() string { return u.umsResourceCategoryDo.Alias() }

func (u umsResourceCategory) Columns(cols ...field.Expr) gen.Columns {
	return u.umsResourceCategoryDo.Columns(cols...)
}

func (u *umsResourceCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsResourceCategory) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["id"] = u.ID
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["name"] = u.Name
	u.fieldMap["sort"] = u.Sort
}

func (u umsResourceCategory) clone(db *gorm.DB) umsResourceCategory {
	u.umsResourceCategoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsResourceCategory) replaceDB(db *gorm.DB) umsResourceCategory {
	u.umsResourceCategoryDo.ReplaceDB(db)
	return u
}

type umsResourceCategoryDo struct{ gen.DO }

type IUmsResourceCategoryDo interface {
	gen.SubQuery
	Debug() IUmsResourceCategoryDo
	WithContext(ctx context.Context) IUmsResourceCategoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsResourceCategoryDo
	WriteDB() IUmsResourceCategoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsResourceCategoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsResourceCategoryDo
	Not(conds ...gen.Condition) IUmsResourceCategoryDo
	Or(conds ...gen.Condition) IUmsResourceCategoryDo
	Select(conds ...field.Expr) IUmsResourceCategoryDo
	Where(conds ...gen.Condition) IUmsResourceCategoryDo
	Order(conds ...field.Expr) IUmsResourceCategoryDo
	Distinct(cols ...field.Expr) IUmsResourceCategoryDo
	Omit(cols ...field.Expr) IUmsResourceCategoryDo
	Join(table schema.Tabler, on ...field.Expr) IUmsResourceCategoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsResourceCategoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsResourceCategoryDo
	Group(cols ...field.Expr) IUmsResourceCategoryDo
	Having(conds ...gen.Condition) IUmsResourceCategoryDo
	Limit(limit int) IUmsResourceCategoryDo
	Offset(offset int) IUmsResourceCategoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsResourceCategoryDo
	Unscoped() IUmsResourceCategoryDo
	Create(values ...*model.UmsResourceCategory) error
	CreateInBatches(values []*model.UmsResourceCategory, batchSize int) error
	Save(values ...*model.UmsResourceCategory) error
	First() (*model.UmsResourceCategory, error)
	Take() (*model.UmsResourceCategory, error)
	Last() (*model.UmsResourceCategory, error)
	Find() ([]*model.UmsResourceCategory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsResourceCategory, err error)
	FindInBatches(result *[]*model.UmsResourceCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsResourceCategory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsResourceCategoryDo
	Assign(attrs ...field.AssignExpr) IUmsResourceCategoryDo
	Joins(fields ...field.RelationField) IUmsResourceCategoryDo
	Preload(fields ...field.RelationField) IUmsResourceCategoryDo
	FirstOrInit() (*model.UmsResourceCategory, error)
	FirstOrCreate() (*model.UmsResourceCategory, error)
	FindByPage(offset int, limit int) (result []*model.UmsResourceCategory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsResourceCategoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsResourceCategoryDo) Debug() IUmsResourceCategoryDo {
	return u.withDO(u.DO.Debug())
}

func (u umsResourceCategoryDo) WithContext(ctx context.Context) IUmsResourceCategoryDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsResourceCategoryDo) ReadDB() IUmsResourceCategoryDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsResourceCategoryDo) WriteDB() IUmsResourceCategoryDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsResourceCategoryDo) Session(config *gorm.Session) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsResourceCategoryDo) Clauses(conds ...clause.Expression) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsResourceCategoryDo) Returning(value interface{}, columns ...string) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsResourceCategoryDo) Not(conds ...gen.Condition) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsResourceCategoryDo) Or(conds ...gen.Condition) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsResourceCategoryDo) Select(conds ...field.Expr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsResourceCategoryDo) Where(conds ...gen.Condition) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsResourceCategoryDo) Order(conds ...field.Expr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsResourceCategoryDo) Distinct(cols ...field.Expr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsResourceCategoryDo) Omit(cols ...field.Expr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsResourceCategoryDo) Join(table schema.Tabler, on ...field.Expr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsResourceCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsResourceCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsResourceCategoryDo) Group(cols ...field.Expr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsResourceCategoryDo) Having(conds ...gen.Condition) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsResourceCategoryDo) Limit(limit int) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsResourceCategoryDo) Offset(offset int) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsResourceCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsResourceCategoryDo) Unscoped() IUmsResourceCategoryDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsResourceCategoryDo) Create(values ...*model.UmsResourceCategory) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsResourceCategoryDo) CreateInBatches(values []*model.UmsResourceCategory, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsResourceCategoryDo) Save(values ...*model.UmsResourceCategory) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsResourceCategoryDo) First() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) Take() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) Last() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) Find() ([]*model.UmsResourceCategory, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsResourceCategory), err
}

func (u umsResourceCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsResourceCategory, err error) {
	buf := make([]*model.UmsResourceCategory, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsResourceCategoryDo) FindInBatches(result *[]*model.UmsResourceCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsResourceCategoryDo) Attrs(attrs ...field.AssignExpr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsResourceCategoryDo) Assign(attrs ...field.AssignExpr) IUmsResourceCategoryDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsResourceCategoryDo) Joins(fields ...field.RelationField) IUmsResourceCategoryDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsResourceCategoryDo) Preload(fields ...field.RelationField) IUmsResourceCategoryDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsResourceCategoryDo) FirstOrInit() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) FirstOrCreate() (*model.UmsResourceCategory, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResourceCategory), nil
	}
}

func (u umsResourceCategoryDo) FindByPage(offset int, limit int) (result []*model.UmsResourceCategory, count int64, err error) {
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

func (u umsResourceCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsResourceCategoryDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsResourceCategoryDo) Delete(models ...*model.UmsResourceCategory) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsResourceCategoryDo) withDO(do gen.Dao) *umsResourceCategoryDo {
	u.DO = *do.(*gen.DO)
	return u
}
