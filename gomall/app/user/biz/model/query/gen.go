// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                                = new(Query)
	UmsAdmin                         *umsAdmin
	UmsAdminLoginLog                 *umsAdminLoginLog
	UmsAdminPermissionRelation       *umsAdminPermissionRelation
	UmsAdminRoleRelation             *umsAdminRoleRelation
	UmsGrowthChangeHistory           *umsGrowthChangeHistory
	UmsIntegrationChangeHistory      *umsIntegrationChangeHistory
	UmsIntegrationConsumeSetting     *umsIntegrationConsumeSetting
	UmsMember                        *umsMember
	UmsMemberLevel                   *umsMemberLevel
	UmsMemberLoginLog                *umsMemberLoginLog
	UmsMemberMemberTagRelation       *umsMemberMemberTagRelation
	UmsMemberProductCategoryRelation *umsMemberProductCategoryRelation
	UmsMemberReceiveAddress          *umsMemberReceiveAddress
	UmsMemberRuleSetting             *umsMemberRuleSetting
	UmsMemberStatisticsInfo          *umsMemberStatisticsInfo
	UmsMemberTag                     *umsMemberTag
	UmsMemberTask                    *umsMemberTask
	UmsMenu                          *umsMenu
	UmsPermission                    *umsPermission
	UmsResource                      *umsResource
	UmsResourceCategory              *umsResourceCategory
	UmsRole                          *umsRole
	UmsRoleMenuRelation              *umsRoleMenuRelation
	UmsRolePermissionRelation        *umsRolePermissionRelation
	UmsRoleResourceRelation          *umsRoleResourceRelation
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	UmsAdmin = &Q.UmsAdmin
	UmsAdminLoginLog = &Q.UmsAdminLoginLog
	UmsAdminPermissionRelation = &Q.UmsAdminPermissionRelation
	UmsAdminRoleRelation = &Q.UmsAdminRoleRelation
	UmsGrowthChangeHistory = &Q.UmsGrowthChangeHistory
	UmsIntegrationChangeHistory = &Q.UmsIntegrationChangeHistory
	UmsIntegrationConsumeSetting = &Q.UmsIntegrationConsumeSetting
	UmsMember = &Q.UmsMember
	UmsMemberLevel = &Q.UmsMemberLevel
	UmsMemberLoginLog = &Q.UmsMemberLoginLog
	UmsMemberMemberTagRelation = &Q.UmsMemberMemberTagRelation
	UmsMemberProductCategoryRelation = &Q.UmsMemberProductCategoryRelation
	UmsMemberReceiveAddress = &Q.UmsMemberReceiveAddress
	UmsMemberRuleSetting = &Q.UmsMemberRuleSetting
	UmsMemberStatisticsInfo = &Q.UmsMemberStatisticsInfo
	UmsMemberTag = &Q.UmsMemberTag
	UmsMemberTask = &Q.UmsMemberTask
	UmsMenu = &Q.UmsMenu
	UmsPermission = &Q.UmsPermission
	UmsResource = &Q.UmsResource
	UmsResourceCategory = &Q.UmsResourceCategory
	UmsRole = &Q.UmsRole
	UmsRoleMenuRelation = &Q.UmsRoleMenuRelation
	UmsRolePermissionRelation = &Q.UmsRolePermissionRelation
	UmsRoleResourceRelation = &Q.UmsRoleResourceRelation
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                               db,
		UmsAdmin:                         newUmsAdmin(db, opts...),
		UmsAdminLoginLog:                 newUmsAdminLoginLog(db, opts...),
		UmsAdminPermissionRelation:       newUmsAdminPermissionRelation(db, opts...),
		UmsAdminRoleRelation:             newUmsAdminRoleRelation(db, opts...),
		UmsGrowthChangeHistory:           newUmsGrowthChangeHistory(db, opts...),
		UmsIntegrationChangeHistory:      newUmsIntegrationChangeHistory(db, opts...),
		UmsIntegrationConsumeSetting:     newUmsIntegrationConsumeSetting(db, opts...),
		UmsMember:                        newUmsMember(db, opts...),
		UmsMemberLevel:                   newUmsMemberLevel(db, opts...),
		UmsMemberLoginLog:                newUmsMemberLoginLog(db, opts...),
		UmsMemberMemberTagRelation:       newUmsMemberMemberTagRelation(db, opts...),
		UmsMemberProductCategoryRelation: newUmsMemberProductCategoryRelation(db, opts...),
		UmsMemberReceiveAddress:          newUmsMemberReceiveAddress(db, opts...),
		UmsMemberRuleSetting:             newUmsMemberRuleSetting(db, opts...),
		UmsMemberStatisticsInfo:          newUmsMemberStatisticsInfo(db, opts...),
		UmsMemberTag:                     newUmsMemberTag(db, opts...),
		UmsMemberTask:                    newUmsMemberTask(db, opts...),
		UmsMenu:                          newUmsMenu(db, opts...),
		UmsPermission:                    newUmsPermission(db, opts...),
		UmsResource:                      newUmsResource(db, opts...),
		UmsResourceCategory:              newUmsResourceCategory(db, opts...),
		UmsRole:                          newUmsRole(db, opts...),
		UmsRoleMenuRelation:              newUmsRoleMenuRelation(db, opts...),
		UmsRolePermissionRelation:        newUmsRolePermissionRelation(db, opts...),
		UmsRoleResourceRelation:          newUmsRoleResourceRelation(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	UmsAdmin                         umsAdmin
	UmsAdminLoginLog                 umsAdminLoginLog
	UmsAdminPermissionRelation       umsAdminPermissionRelation
	UmsAdminRoleRelation             umsAdminRoleRelation
	UmsGrowthChangeHistory           umsGrowthChangeHistory
	UmsIntegrationChangeHistory      umsIntegrationChangeHistory
	UmsIntegrationConsumeSetting     umsIntegrationConsumeSetting
	UmsMember                        umsMember
	UmsMemberLevel                   umsMemberLevel
	UmsMemberLoginLog                umsMemberLoginLog
	UmsMemberMemberTagRelation       umsMemberMemberTagRelation
	UmsMemberProductCategoryRelation umsMemberProductCategoryRelation
	UmsMemberReceiveAddress          umsMemberReceiveAddress
	UmsMemberRuleSetting             umsMemberRuleSetting
	UmsMemberStatisticsInfo          umsMemberStatisticsInfo
	UmsMemberTag                     umsMemberTag
	UmsMemberTask                    umsMemberTask
	UmsMenu                          umsMenu
	UmsPermission                    umsPermission
	UmsResource                      umsResource
	UmsResourceCategory              umsResourceCategory
	UmsRole                          umsRole
	UmsRoleMenuRelation              umsRoleMenuRelation
	UmsRolePermissionRelation        umsRolePermissionRelation
	UmsRoleResourceRelation          umsRoleResourceRelation
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                               db,
		UmsAdmin:                         q.UmsAdmin.clone(db),
		UmsAdminLoginLog:                 q.UmsAdminLoginLog.clone(db),
		UmsAdminPermissionRelation:       q.UmsAdminPermissionRelation.clone(db),
		UmsAdminRoleRelation:             q.UmsAdminRoleRelation.clone(db),
		UmsGrowthChangeHistory:           q.UmsGrowthChangeHistory.clone(db),
		UmsIntegrationChangeHistory:      q.UmsIntegrationChangeHistory.clone(db),
		UmsIntegrationConsumeSetting:     q.UmsIntegrationConsumeSetting.clone(db),
		UmsMember:                        q.UmsMember.clone(db),
		UmsMemberLevel:                   q.UmsMemberLevel.clone(db),
		UmsMemberLoginLog:                q.UmsMemberLoginLog.clone(db),
		UmsMemberMemberTagRelation:       q.UmsMemberMemberTagRelation.clone(db),
		UmsMemberProductCategoryRelation: q.UmsMemberProductCategoryRelation.clone(db),
		UmsMemberReceiveAddress:          q.UmsMemberReceiveAddress.clone(db),
		UmsMemberRuleSetting:             q.UmsMemberRuleSetting.clone(db),
		UmsMemberStatisticsInfo:          q.UmsMemberStatisticsInfo.clone(db),
		UmsMemberTag:                     q.UmsMemberTag.clone(db),
		UmsMemberTask:                    q.UmsMemberTask.clone(db),
		UmsMenu:                          q.UmsMenu.clone(db),
		UmsPermission:                    q.UmsPermission.clone(db),
		UmsResource:                      q.UmsResource.clone(db),
		UmsResourceCategory:              q.UmsResourceCategory.clone(db),
		UmsRole:                          q.UmsRole.clone(db),
		UmsRoleMenuRelation:              q.UmsRoleMenuRelation.clone(db),
		UmsRolePermissionRelation:        q.UmsRolePermissionRelation.clone(db),
		UmsRoleResourceRelation:          q.UmsRoleResourceRelation.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                               db,
		UmsAdmin:                         q.UmsAdmin.replaceDB(db),
		UmsAdminLoginLog:                 q.UmsAdminLoginLog.replaceDB(db),
		UmsAdminPermissionRelation:       q.UmsAdminPermissionRelation.replaceDB(db),
		UmsAdminRoleRelation:             q.UmsAdminRoleRelation.replaceDB(db),
		UmsGrowthChangeHistory:           q.UmsGrowthChangeHistory.replaceDB(db),
		UmsIntegrationChangeHistory:      q.UmsIntegrationChangeHistory.replaceDB(db),
		UmsIntegrationConsumeSetting:     q.UmsIntegrationConsumeSetting.replaceDB(db),
		UmsMember:                        q.UmsMember.replaceDB(db),
		UmsMemberLevel:                   q.UmsMemberLevel.replaceDB(db),
		UmsMemberLoginLog:                q.UmsMemberLoginLog.replaceDB(db),
		UmsMemberMemberTagRelation:       q.UmsMemberMemberTagRelation.replaceDB(db),
		UmsMemberProductCategoryRelation: q.UmsMemberProductCategoryRelation.replaceDB(db),
		UmsMemberReceiveAddress:          q.UmsMemberReceiveAddress.replaceDB(db),
		UmsMemberRuleSetting:             q.UmsMemberRuleSetting.replaceDB(db),
		UmsMemberStatisticsInfo:          q.UmsMemberStatisticsInfo.replaceDB(db),
		UmsMemberTag:                     q.UmsMemberTag.replaceDB(db),
		UmsMemberTask:                    q.UmsMemberTask.replaceDB(db),
		UmsMenu:                          q.UmsMenu.replaceDB(db),
		UmsPermission:                    q.UmsPermission.replaceDB(db),
		UmsResource:                      q.UmsResource.replaceDB(db),
		UmsResourceCategory:              q.UmsResourceCategory.replaceDB(db),
		UmsRole:                          q.UmsRole.replaceDB(db),
		UmsRoleMenuRelation:              q.UmsRoleMenuRelation.replaceDB(db),
		UmsRolePermissionRelation:        q.UmsRolePermissionRelation.replaceDB(db),
		UmsRoleResourceRelation:          q.UmsRoleResourceRelation.replaceDB(db),
	}
}

type queryCtx struct {
	UmsAdmin                         IUmsAdminDo
	UmsAdminLoginLog                 IUmsAdminLoginLogDo
	UmsAdminPermissionRelation       IUmsAdminPermissionRelationDo
	UmsAdminRoleRelation             IUmsAdminRoleRelationDo
	UmsGrowthChangeHistory           IUmsGrowthChangeHistoryDo
	UmsIntegrationChangeHistory      IUmsIntegrationChangeHistoryDo
	UmsIntegrationConsumeSetting     IUmsIntegrationConsumeSettingDo
	UmsMember                        IUmsMemberDo
	UmsMemberLevel                   IUmsMemberLevelDo
	UmsMemberLoginLog                IUmsMemberLoginLogDo
	UmsMemberMemberTagRelation       IUmsMemberMemberTagRelationDo
	UmsMemberProductCategoryRelation IUmsMemberProductCategoryRelationDo
	UmsMemberReceiveAddress          IUmsMemberReceiveAddressDo
	UmsMemberRuleSetting             IUmsMemberRuleSettingDo
	UmsMemberStatisticsInfo          IUmsMemberStatisticsInfoDo
	UmsMemberTag                     IUmsMemberTagDo
	UmsMemberTask                    IUmsMemberTaskDo
	UmsMenu                          IUmsMenuDo
	UmsPermission                    IUmsPermissionDo
	UmsResource                      IUmsResourceDo
	UmsResourceCategory              IUmsResourceCategoryDo
	UmsRole                          IUmsRoleDo
	UmsRoleMenuRelation              IUmsRoleMenuRelationDo
	UmsRolePermissionRelation        IUmsRolePermissionRelationDo
	UmsRoleResourceRelation          IUmsRoleResourceRelationDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		UmsAdmin:                         q.UmsAdmin.WithContext(ctx),
		UmsAdminLoginLog:                 q.UmsAdminLoginLog.WithContext(ctx),
		UmsAdminPermissionRelation:       q.UmsAdminPermissionRelation.WithContext(ctx),
		UmsAdminRoleRelation:             q.UmsAdminRoleRelation.WithContext(ctx),
		UmsGrowthChangeHistory:           q.UmsGrowthChangeHistory.WithContext(ctx),
		UmsIntegrationChangeHistory:      q.UmsIntegrationChangeHistory.WithContext(ctx),
		UmsIntegrationConsumeSetting:     q.UmsIntegrationConsumeSetting.WithContext(ctx),
		UmsMember:                        q.UmsMember.WithContext(ctx),
		UmsMemberLevel:                   q.UmsMemberLevel.WithContext(ctx),
		UmsMemberLoginLog:                q.UmsMemberLoginLog.WithContext(ctx),
		UmsMemberMemberTagRelation:       q.UmsMemberMemberTagRelation.WithContext(ctx),
		UmsMemberProductCategoryRelation: q.UmsMemberProductCategoryRelation.WithContext(ctx),
		UmsMemberReceiveAddress:          q.UmsMemberReceiveAddress.WithContext(ctx),
		UmsMemberRuleSetting:             q.UmsMemberRuleSetting.WithContext(ctx),
		UmsMemberStatisticsInfo:          q.UmsMemberStatisticsInfo.WithContext(ctx),
		UmsMemberTag:                     q.UmsMemberTag.WithContext(ctx),
		UmsMemberTask:                    q.UmsMemberTask.WithContext(ctx),
		UmsMenu:                          q.UmsMenu.WithContext(ctx),
		UmsPermission:                    q.UmsPermission.WithContext(ctx),
		UmsResource:                      q.UmsResource.WithContext(ctx),
		UmsResourceCategory:              q.UmsResourceCategory.WithContext(ctx),
		UmsRole:                          q.UmsRole.WithContext(ctx),
		UmsRoleMenuRelation:              q.UmsRoleMenuRelation.WithContext(ctx),
		UmsRolePermissionRelation:        q.UmsRolePermissionRelation.WithContext(ctx),
		UmsRoleResourceRelation:          q.UmsRoleResourceRelation.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
