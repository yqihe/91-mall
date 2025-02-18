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

func newUmsMember(db *gorm.DB, opts ...gen.DOOption) umsMember {
	_umsMember := umsMember{}

	_umsMember.umsMemberDo.UseDB(db, opts...)
	_umsMember.umsMemberDo.UseModel(&model.UmsMember{})

	tableName := _umsMember.umsMemberDo.TableName()
	_umsMember.ALL = field.NewAsterisk(tableName)
	_umsMember.ID = field.NewInt64(tableName, "id")
	_umsMember.MemberLevelID = field.NewInt64(tableName, "member_level_id")
	_umsMember.Username = field.NewString(tableName, "username")
	_umsMember.Password = field.NewString(tableName, "password")
	_umsMember.Nickname = field.NewString(tableName, "nickname")
	_umsMember.Phone = field.NewString(tableName, "phone")
	_umsMember.Status = field.NewInt32(tableName, "status")
	_umsMember.CreateTime = field.NewTime(tableName, "create_time")
	_umsMember.Icon = field.NewString(tableName, "icon")
	_umsMember.Gender = field.NewInt32(tableName, "gender")
	_umsMember.Birthday = field.NewTime(tableName, "birthday")
	_umsMember.City = field.NewString(tableName, "city")
	_umsMember.Job = field.NewString(tableName, "job")
	_umsMember.PersonalizedSignature = field.NewString(tableName, "personalized_signature")
	_umsMember.SourceType = field.NewInt32(tableName, "source_type")
	_umsMember.Integration = field.NewInt32(tableName, "integration")
	_umsMember.Growth = field.NewInt32(tableName, "growth")
	_umsMember.LuckeyCount = field.NewInt32(tableName, "luckey_count")
	_umsMember.HistoryIntegration = field.NewInt32(tableName, "history_integration")

	_umsMember.fillFieldMap()

	return _umsMember
}

// umsMember 会员表
type umsMember struct {
	umsMemberDo

	ALL                   field.Asterisk
	ID                    field.Int64
	MemberLevelID         field.Int64
	Username              field.String // 用户名
	Password              field.String // 密码
	Nickname              field.String // 昵称
	Phone                 field.String // 手机号码
	Status                field.Int32  // 帐号启用状态:0->禁用；1->启用
	CreateTime            field.Time   // 注册时间
	Icon                  field.String // 头像
	Gender                field.Int32  // 性别：0->未知；1->男；2->女
	Birthday              field.Time   // 生日
	City                  field.String // 所做城市
	Job                   field.String // 职业
	PersonalizedSignature field.String // 个性签名
	SourceType            field.Int32  // 用户来源
	Integration           field.Int32  // 积分
	Growth                field.Int32  // 成长值
	LuckeyCount           field.Int32  // 剩余抽奖次数
	HistoryIntegration    field.Int32  // 历史积分数量

	fieldMap map[string]field.Expr
}

func (u umsMember) Table(newTableName string) *umsMember {
	u.umsMemberDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsMember) As(alias string) *umsMember {
	u.umsMemberDo.DO = *(u.umsMemberDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsMember) updateTableName(table string) *umsMember {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.MemberLevelID = field.NewInt64(table, "member_level_id")
	u.Username = field.NewString(table, "username")
	u.Password = field.NewString(table, "password")
	u.Nickname = field.NewString(table, "nickname")
	u.Phone = field.NewString(table, "phone")
	u.Status = field.NewInt32(table, "status")
	u.CreateTime = field.NewTime(table, "create_time")
	u.Icon = field.NewString(table, "icon")
	u.Gender = field.NewInt32(table, "gender")
	u.Birthday = field.NewTime(table, "birthday")
	u.City = field.NewString(table, "city")
	u.Job = field.NewString(table, "job")
	u.PersonalizedSignature = field.NewString(table, "personalized_signature")
	u.SourceType = field.NewInt32(table, "source_type")
	u.Integration = field.NewInt32(table, "integration")
	u.Growth = field.NewInt32(table, "growth")
	u.LuckeyCount = field.NewInt32(table, "luckey_count")
	u.HistoryIntegration = field.NewInt32(table, "history_integration")

	u.fillFieldMap()

	return u
}

func (u *umsMember) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsMember) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 19)
	u.fieldMap["id"] = u.ID
	u.fieldMap["member_level_id"] = u.MemberLevelID
	u.fieldMap["username"] = u.Username
	u.fieldMap["password"] = u.Password
	u.fieldMap["nickname"] = u.Nickname
	u.fieldMap["phone"] = u.Phone
	u.fieldMap["status"] = u.Status
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["icon"] = u.Icon
	u.fieldMap["gender"] = u.Gender
	u.fieldMap["birthday"] = u.Birthday
	u.fieldMap["city"] = u.City
	u.fieldMap["job"] = u.Job
	u.fieldMap["personalized_signature"] = u.PersonalizedSignature
	u.fieldMap["source_type"] = u.SourceType
	u.fieldMap["integration"] = u.Integration
	u.fieldMap["growth"] = u.Growth
	u.fieldMap["luckey_count"] = u.LuckeyCount
	u.fieldMap["history_integration"] = u.HistoryIntegration
}

func (u umsMember) clone(db *gorm.DB) umsMember {
	u.umsMemberDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsMember) replaceDB(db *gorm.DB) umsMember {
	u.umsMemberDo.ReplaceDB(db)
	return u
}

type umsMemberDo struct{ gen.DO }

type IUmsMemberDo interface {
	gen.SubQuery
	Debug() IUmsMemberDo
	WithContext(ctx context.Context) IUmsMemberDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsMemberDo
	WriteDB() IUmsMemberDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsMemberDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsMemberDo
	Not(conds ...gen.Condition) IUmsMemberDo
	Or(conds ...gen.Condition) IUmsMemberDo
	Select(conds ...field.Expr) IUmsMemberDo
	Where(conds ...gen.Condition) IUmsMemberDo
	Order(conds ...field.Expr) IUmsMemberDo
	Distinct(cols ...field.Expr) IUmsMemberDo
	Omit(cols ...field.Expr) IUmsMemberDo
	Join(table schema.Tabler, on ...field.Expr) IUmsMemberDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberDo
	Group(cols ...field.Expr) IUmsMemberDo
	Having(conds ...gen.Condition) IUmsMemberDo
	Limit(limit int) IUmsMemberDo
	Offset(offset int) IUmsMemberDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberDo
	Unscoped() IUmsMemberDo
	Create(values ...*model.UmsMember) error
	CreateInBatches(values []*model.UmsMember, batchSize int) error
	Save(values ...*model.UmsMember) error
	First() (*model.UmsMember, error)
	Take() (*model.UmsMember, error)
	Last() (*model.UmsMember, error)
	Find() ([]*model.UmsMember, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMember, err error)
	FindInBatches(result *[]*model.UmsMember, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsMember) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsMemberDo
	Assign(attrs ...field.AssignExpr) IUmsMemberDo
	Joins(fields ...field.RelationField) IUmsMemberDo
	Preload(fields ...field.RelationField) IUmsMemberDo
	FirstOrInit() (*model.UmsMember, error)
	FirstOrCreate() (*model.UmsMember, error)
	FindByPage(offset int, limit int) (result []*model.UmsMember, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsMemberDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsMemberDo) Debug() IUmsMemberDo {
	return u.withDO(u.DO.Debug())
}

func (u umsMemberDo) WithContext(ctx context.Context) IUmsMemberDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsMemberDo) ReadDB() IUmsMemberDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsMemberDo) WriteDB() IUmsMemberDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsMemberDo) Session(config *gorm.Session) IUmsMemberDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsMemberDo) Clauses(conds ...clause.Expression) IUmsMemberDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsMemberDo) Returning(value interface{}, columns ...string) IUmsMemberDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsMemberDo) Not(conds ...gen.Condition) IUmsMemberDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsMemberDo) Or(conds ...gen.Condition) IUmsMemberDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsMemberDo) Select(conds ...field.Expr) IUmsMemberDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsMemberDo) Where(conds ...gen.Condition) IUmsMemberDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsMemberDo) Order(conds ...field.Expr) IUmsMemberDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsMemberDo) Distinct(cols ...field.Expr) IUmsMemberDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsMemberDo) Omit(cols ...field.Expr) IUmsMemberDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsMemberDo) Join(table schema.Tabler, on ...field.Expr) IUmsMemberDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsMemberDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsMemberDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsMemberDo) Group(cols ...field.Expr) IUmsMemberDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsMemberDo) Having(conds ...gen.Condition) IUmsMemberDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsMemberDo) Limit(limit int) IUmsMemberDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsMemberDo) Offset(offset int) IUmsMemberDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsMemberDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsMemberDo) Unscoped() IUmsMemberDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsMemberDo) Create(values ...*model.UmsMember) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsMemberDo) CreateInBatches(values []*model.UmsMember, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsMemberDo) Save(values ...*model.UmsMember) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsMemberDo) First() (*model.UmsMember, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMember), nil
	}
}

func (u umsMemberDo) Take() (*model.UmsMember, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMember), nil
	}
}

func (u umsMemberDo) Last() (*model.UmsMember, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMember), nil
	}
}

func (u umsMemberDo) Find() ([]*model.UmsMember, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsMember), err
}

func (u umsMemberDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMember, err error) {
	buf := make([]*model.UmsMember, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsMemberDo) FindInBatches(result *[]*model.UmsMember, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsMemberDo) Attrs(attrs ...field.AssignExpr) IUmsMemberDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsMemberDo) Assign(attrs ...field.AssignExpr) IUmsMemberDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsMemberDo) Joins(fields ...field.RelationField) IUmsMemberDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsMemberDo) Preload(fields ...field.RelationField) IUmsMemberDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsMemberDo) FirstOrInit() (*model.UmsMember, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMember), nil
	}
}

func (u umsMemberDo) FirstOrCreate() (*model.UmsMember, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMember), nil
	}
}

func (u umsMemberDo) FindByPage(offset int, limit int) (result []*model.UmsMember, count int64, err error) {
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

func (u umsMemberDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsMemberDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsMemberDo) Delete(models ...*model.UmsMember) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsMemberDo) withDO(do gen.Dao) *umsMemberDo {
	u.DO = *do.(*gen.DO)
	return u
}
