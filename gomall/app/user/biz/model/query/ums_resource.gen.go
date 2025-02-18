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

func newUmsResource(db *gorm.DB, opts ...gen.DOOption) umsResource {
	_umsResource := umsResource{}

	_umsResource.umsResourceDo.UseDB(db, opts...)
	_umsResource.umsResourceDo.UseModel(&model.UmsResource{})

	tableName := _umsResource.umsResourceDo.TableName()
	_umsResource.ALL = field.NewAsterisk(tableName)
	_umsResource.ID = field.NewInt64(tableName, "id")
	_umsResource.CreateTime = field.NewTime(tableName, "create_time")
	_umsResource.Name = field.NewString(tableName, "name")
	_umsResource.URL = field.NewString(tableName, "url")
	_umsResource.Description = field.NewString(tableName, "description")
	_umsResource.CategoryID = field.NewInt64(tableName, "category_id")

	_umsResource.fillFieldMap()

	return _umsResource
}

// umsResource 后台资源表
type umsResource struct {
	umsResourceDo

	ALL         field.Asterisk
	ID          field.Int64
	CreateTime  field.Time   // 创建时间
	Name        field.String // 资源名称
	URL         field.String // 资源URL
	Description field.String // 描述
	CategoryID  field.Int64  // 资源分类ID

	fieldMap map[string]field.Expr
}

func (u umsResource) Table(newTableName string) *umsResource {
	u.umsResourceDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsResource) As(alias string) *umsResource {
	u.umsResourceDo.DO = *(u.umsResourceDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsResource) updateTableName(table string) *umsResource {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.CreateTime = field.NewTime(table, "create_time")
	u.Name = field.NewString(table, "name")
	u.URL = field.NewString(table, "url")
	u.Description = field.NewString(table, "description")
	u.CategoryID = field.NewInt64(table, "category_id")

	u.fillFieldMap()

	return u
}

func (u *umsResource) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsResource) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["name"] = u.Name
	u.fieldMap["url"] = u.URL
	u.fieldMap["description"] = u.Description
	u.fieldMap["category_id"] = u.CategoryID
}

func (u umsResource) clone(db *gorm.DB) umsResource {
	u.umsResourceDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsResource) replaceDB(db *gorm.DB) umsResource {
	u.umsResourceDo.ReplaceDB(db)
	return u
}

type umsResourceDo struct{ gen.DO }

type IUmsResourceDo interface {
	gen.SubQuery
	Debug() IUmsResourceDo
	WithContext(ctx context.Context) IUmsResourceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsResourceDo
	WriteDB() IUmsResourceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsResourceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsResourceDo
	Not(conds ...gen.Condition) IUmsResourceDo
	Or(conds ...gen.Condition) IUmsResourceDo
	Select(conds ...field.Expr) IUmsResourceDo
	Where(conds ...gen.Condition) IUmsResourceDo
	Order(conds ...field.Expr) IUmsResourceDo
	Distinct(cols ...field.Expr) IUmsResourceDo
	Omit(cols ...field.Expr) IUmsResourceDo
	Join(table schema.Tabler, on ...field.Expr) IUmsResourceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsResourceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsResourceDo
	Group(cols ...field.Expr) IUmsResourceDo
	Having(conds ...gen.Condition) IUmsResourceDo
	Limit(limit int) IUmsResourceDo
	Offset(offset int) IUmsResourceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsResourceDo
	Unscoped() IUmsResourceDo
	Create(values ...*model.UmsResource) error
	CreateInBatches(values []*model.UmsResource, batchSize int) error
	Save(values ...*model.UmsResource) error
	First() (*model.UmsResource, error)
	Take() (*model.UmsResource, error)
	Last() (*model.UmsResource, error)
	Find() ([]*model.UmsResource, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsResource, err error)
	FindInBatches(result *[]*model.UmsResource, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsResource) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsResourceDo
	Assign(attrs ...field.AssignExpr) IUmsResourceDo
	Joins(fields ...field.RelationField) IUmsResourceDo
	Preload(fields ...field.RelationField) IUmsResourceDo
	FirstOrInit() (*model.UmsResource, error)
	FirstOrCreate() (*model.UmsResource, error)
	FindByPage(offset int, limit int) (result []*model.UmsResource, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsResourceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsResourceDo) Debug() IUmsResourceDo {
	return u.withDO(u.DO.Debug())
}

func (u umsResourceDo) WithContext(ctx context.Context) IUmsResourceDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsResourceDo) ReadDB() IUmsResourceDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsResourceDo) WriteDB() IUmsResourceDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsResourceDo) Session(config *gorm.Session) IUmsResourceDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsResourceDo) Clauses(conds ...clause.Expression) IUmsResourceDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsResourceDo) Returning(value interface{}, columns ...string) IUmsResourceDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsResourceDo) Not(conds ...gen.Condition) IUmsResourceDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsResourceDo) Or(conds ...gen.Condition) IUmsResourceDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsResourceDo) Select(conds ...field.Expr) IUmsResourceDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsResourceDo) Where(conds ...gen.Condition) IUmsResourceDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsResourceDo) Order(conds ...field.Expr) IUmsResourceDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsResourceDo) Distinct(cols ...field.Expr) IUmsResourceDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsResourceDo) Omit(cols ...field.Expr) IUmsResourceDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsResourceDo) Join(table schema.Tabler, on ...field.Expr) IUmsResourceDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsResourceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsResourceDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsResourceDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsResourceDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsResourceDo) Group(cols ...field.Expr) IUmsResourceDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsResourceDo) Having(conds ...gen.Condition) IUmsResourceDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsResourceDo) Limit(limit int) IUmsResourceDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsResourceDo) Offset(offset int) IUmsResourceDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsResourceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsResourceDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsResourceDo) Unscoped() IUmsResourceDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsResourceDo) Create(values ...*model.UmsResource) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsResourceDo) CreateInBatches(values []*model.UmsResource, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsResourceDo) Save(values ...*model.UmsResource) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsResourceDo) First() (*model.UmsResource, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResource), nil
	}
}

func (u umsResourceDo) Take() (*model.UmsResource, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResource), nil
	}
}

func (u umsResourceDo) Last() (*model.UmsResource, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResource), nil
	}
}

func (u umsResourceDo) Find() ([]*model.UmsResource, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsResource), err
}

func (u umsResourceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsResource, err error) {
	buf := make([]*model.UmsResource, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsResourceDo) FindInBatches(result *[]*model.UmsResource, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsResourceDo) Attrs(attrs ...field.AssignExpr) IUmsResourceDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsResourceDo) Assign(attrs ...field.AssignExpr) IUmsResourceDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsResourceDo) Joins(fields ...field.RelationField) IUmsResourceDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsResourceDo) Preload(fields ...field.RelationField) IUmsResourceDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsResourceDo) FirstOrInit() (*model.UmsResource, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResource), nil
	}
}

func (u umsResourceDo) FirstOrCreate() (*model.UmsResource, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsResource), nil
	}
}

func (u umsResourceDo) FindByPage(offset int, limit int) (result []*model.UmsResource, count int64, err error) {
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

func (u umsResourceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsResourceDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsResourceDo) Delete(models ...*model.UmsResource) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsResourceDo) withDO(do gen.Dao) *umsResourceDo {
	u.DO = *do.(*gen.DO)
	return u
}
