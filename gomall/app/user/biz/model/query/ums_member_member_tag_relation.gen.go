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

func newUmsMemberMemberTagRelation(db *gorm.DB, opts ...gen.DOOption) umsMemberMemberTagRelation {
	_umsMemberMemberTagRelation := umsMemberMemberTagRelation{}

	_umsMemberMemberTagRelation.umsMemberMemberTagRelationDo.UseDB(db, opts...)
	_umsMemberMemberTagRelation.umsMemberMemberTagRelationDo.UseModel(&model.UmsMemberMemberTagRelation{})

	tableName := _umsMemberMemberTagRelation.umsMemberMemberTagRelationDo.TableName()
	_umsMemberMemberTagRelation.ALL = field.NewAsterisk(tableName)
	_umsMemberMemberTagRelation.ID = field.NewInt64(tableName, "id")
	_umsMemberMemberTagRelation.MemberID = field.NewInt64(tableName, "member_id")
	_umsMemberMemberTagRelation.TagID = field.NewInt64(tableName, "tag_id")

	_umsMemberMemberTagRelation.fillFieldMap()

	return _umsMemberMemberTagRelation
}

// umsMemberMemberTagRelation 用户和标签关系表
type umsMemberMemberTagRelation struct {
	umsMemberMemberTagRelationDo

	ALL      field.Asterisk
	ID       field.Int64
	MemberID field.Int64
	TagID    field.Int64

	fieldMap map[string]field.Expr
}

func (u umsMemberMemberTagRelation) Table(newTableName string) *umsMemberMemberTagRelation {
	u.umsMemberMemberTagRelationDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsMemberMemberTagRelation) As(alias string) *umsMemberMemberTagRelation {
	u.umsMemberMemberTagRelationDo.DO = *(u.umsMemberMemberTagRelationDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsMemberMemberTagRelation) updateTableName(table string) *umsMemberMemberTagRelation {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.MemberID = field.NewInt64(table, "member_id")
	u.TagID = field.NewInt64(table, "tag_id")

	u.fillFieldMap()

	return u
}

func (u *umsMemberMemberTagRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsMemberMemberTagRelation) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 3)
	u.fieldMap["id"] = u.ID
	u.fieldMap["member_id"] = u.MemberID
	u.fieldMap["tag_id"] = u.TagID
}

func (u umsMemberMemberTagRelation) clone(db *gorm.DB) umsMemberMemberTagRelation {
	u.umsMemberMemberTagRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsMemberMemberTagRelation) replaceDB(db *gorm.DB) umsMemberMemberTagRelation {
	u.umsMemberMemberTagRelationDo.ReplaceDB(db)
	return u
}

type umsMemberMemberTagRelationDo struct{ gen.DO }

type IUmsMemberMemberTagRelationDo interface {
	gen.SubQuery
	Debug() IUmsMemberMemberTagRelationDo
	WithContext(ctx context.Context) IUmsMemberMemberTagRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsMemberMemberTagRelationDo
	WriteDB() IUmsMemberMemberTagRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsMemberMemberTagRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsMemberMemberTagRelationDo
	Not(conds ...gen.Condition) IUmsMemberMemberTagRelationDo
	Or(conds ...gen.Condition) IUmsMemberMemberTagRelationDo
	Select(conds ...field.Expr) IUmsMemberMemberTagRelationDo
	Where(conds ...gen.Condition) IUmsMemberMemberTagRelationDo
	Order(conds ...field.Expr) IUmsMemberMemberTagRelationDo
	Distinct(cols ...field.Expr) IUmsMemberMemberTagRelationDo
	Omit(cols ...field.Expr) IUmsMemberMemberTagRelationDo
	Join(table schema.Tabler, on ...field.Expr) IUmsMemberMemberTagRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberMemberTagRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberMemberTagRelationDo
	Group(cols ...field.Expr) IUmsMemberMemberTagRelationDo
	Having(conds ...gen.Condition) IUmsMemberMemberTagRelationDo
	Limit(limit int) IUmsMemberMemberTagRelationDo
	Offset(offset int) IUmsMemberMemberTagRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberMemberTagRelationDo
	Unscoped() IUmsMemberMemberTagRelationDo
	Create(values ...*model.UmsMemberMemberTagRelation) error
	CreateInBatches(values []*model.UmsMemberMemberTagRelation, batchSize int) error
	Save(values ...*model.UmsMemberMemberTagRelation) error
	First() (*model.UmsMemberMemberTagRelation, error)
	Take() (*model.UmsMemberMemberTagRelation, error)
	Last() (*model.UmsMemberMemberTagRelation, error)
	Find() ([]*model.UmsMemberMemberTagRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberMemberTagRelation, err error)
	FindInBatches(result *[]*model.UmsMemberMemberTagRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsMemberMemberTagRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsMemberMemberTagRelationDo
	Assign(attrs ...field.AssignExpr) IUmsMemberMemberTagRelationDo
	Joins(fields ...field.RelationField) IUmsMemberMemberTagRelationDo
	Preload(fields ...field.RelationField) IUmsMemberMemberTagRelationDo
	FirstOrInit() (*model.UmsMemberMemberTagRelation, error)
	FirstOrCreate() (*model.UmsMemberMemberTagRelation, error)
	FindByPage(offset int, limit int) (result []*model.UmsMemberMemberTagRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsMemberMemberTagRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsMemberMemberTagRelationDo) Debug() IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Debug())
}

func (u umsMemberMemberTagRelationDo) WithContext(ctx context.Context) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsMemberMemberTagRelationDo) ReadDB() IUmsMemberMemberTagRelationDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsMemberMemberTagRelationDo) WriteDB() IUmsMemberMemberTagRelationDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsMemberMemberTagRelationDo) Session(config *gorm.Session) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsMemberMemberTagRelationDo) Clauses(conds ...clause.Expression) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsMemberMemberTagRelationDo) Returning(value interface{}, columns ...string) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsMemberMemberTagRelationDo) Not(conds ...gen.Condition) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsMemberMemberTagRelationDo) Or(conds ...gen.Condition) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsMemberMemberTagRelationDo) Select(conds ...field.Expr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsMemberMemberTagRelationDo) Where(conds ...gen.Condition) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsMemberMemberTagRelationDo) Order(conds ...field.Expr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsMemberMemberTagRelationDo) Distinct(cols ...field.Expr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsMemberMemberTagRelationDo) Omit(cols ...field.Expr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsMemberMemberTagRelationDo) Join(table schema.Tabler, on ...field.Expr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsMemberMemberTagRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsMemberMemberTagRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsMemberMemberTagRelationDo) Group(cols ...field.Expr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsMemberMemberTagRelationDo) Having(conds ...gen.Condition) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsMemberMemberTagRelationDo) Limit(limit int) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsMemberMemberTagRelationDo) Offset(offset int) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsMemberMemberTagRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsMemberMemberTagRelationDo) Unscoped() IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsMemberMemberTagRelationDo) Create(values ...*model.UmsMemberMemberTagRelation) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsMemberMemberTagRelationDo) CreateInBatches(values []*model.UmsMemberMemberTagRelation, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsMemberMemberTagRelationDo) Save(values ...*model.UmsMemberMemberTagRelation) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsMemberMemberTagRelationDo) First() (*model.UmsMemberMemberTagRelation, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberMemberTagRelation), nil
	}
}

func (u umsMemberMemberTagRelationDo) Take() (*model.UmsMemberMemberTagRelation, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberMemberTagRelation), nil
	}
}

func (u umsMemberMemberTagRelationDo) Last() (*model.UmsMemberMemberTagRelation, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberMemberTagRelation), nil
	}
}

func (u umsMemberMemberTagRelationDo) Find() ([]*model.UmsMemberMemberTagRelation, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsMemberMemberTagRelation), err
}

func (u umsMemberMemberTagRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberMemberTagRelation, err error) {
	buf := make([]*model.UmsMemberMemberTagRelation, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsMemberMemberTagRelationDo) FindInBatches(result *[]*model.UmsMemberMemberTagRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsMemberMemberTagRelationDo) Attrs(attrs ...field.AssignExpr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsMemberMemberTagRelationDo) Assign(attrs ...field.AssignExpr) IUmsMemberMemberTagRelationDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsMemberMemberTagRelationDo) Joins(fields ...field.RelationField) IUmsMemberMemberTagRelationDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsMemberMemberTagRelationDo) Preload(fields ...field.RelationField) IUmsMemberMemberTagRelationDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsMemberMemberTagRelationDo) FirstOrInit() (*model.UmsMemberMemberTagRelation, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberMemberTagRelation), nil
	}
}

func (u umsMemberMemberTagRelationDo) FirstOrCreate() (*model.UmsMemberMemberTagRelation, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberMemberTagRelation), nil
	}
}

func (u umsMemberMemberTagRelationDo) FindByPage(offset int, limit int) (result []*model.UmsMemberMemberTagRelation, count int64, err error) {
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

func (u umsMemberMemberTagRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsMemberMemberTagRelationDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsMemberMemberTagRelationDo) Delete(models ...*model.UmsMemberMemberTagRelation) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsMemberMemberTagRelationDo) withDO(do gen.Dao) *umsMemberMemberTagRelationDo {
	u.DO = *do.(*gen.DO)
	return u
}
