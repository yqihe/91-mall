package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	redisClient "github.com/yqihe/91-mall/gomall/app/user_admin/biz/dal/redis"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/model"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/mymodel"
	"github.com/yqihe/91-mall/gomall/app/user_admin/pack"
	"github.com/yqihe/91-mall/gomall/pkg/constants"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"github.com/yqihe/91-mall/gomall/pkg/utils"
	"github.com/yqihe/91-mall/gomall/pkg/utils/auth"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type LoginService struct {
	ctx context.Context
}

// NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user_admin.LoginReq) (resp *user_admin.LoginResp, err error) {
	// Finish your business logic.
	resp = new(user_admin.LoginResp)

	// 并行加载用户信息和资源列表
	//var wg sync.WaitGroup
	var userDetails *mymodel.AdminUserDetails
	//var userErr, resourceErr error
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	userDetails, userErr = s.loadUserByUsername(req.Username)
	//}()

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	// 获取资源列表的相关逻辑可以在这里并行化
	//	// resourceList, resourceErr = s.loadResourceListForUser(req.Username)
	//}()
	//
	//wg.Wait()

	// 加载用户详情
	userDetails, err = s.loadUserByUsername(req.Username)
	if err != nil {
		klog.Errorf("[user_admin-service] Login loadUserByUsername failed, err: %s", err.Error())
		return nil, err
	}
	// 密码需要客户端加密后传递
	if !utils.Matches(req.Password, userDetails.UmsAdmin.Password) {
		return nil, errno.PasswordNotCorrectError
	}
	// 检查用户是否被禁用
	if !userDetails.IsEnabled() {
		return nil, errno.UserNameIsNotEnabledError
	}
	// 生成 JWT Token
	token, err := auth.CreateToken(userDetails.UmsAdmin.Username)
	if err != nil {
		klog.Errorf("[user_admin-service] Login CreateToken failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	// 插入登录日志
	if err = s.insertLoginLog(userDetails.UmsAdmin.Username); err != nil {
		klog.Errorf("[user_admin-service] Login InsertLoginLog failed, err: %s", err.Error())
		return nil, err
	}
	resp.Resp = pack.BuildBaseResp(nil)
	resp.Data = &user_admin.TokenMap{
		Token:     token,
		TokenHead: constants.TOKENHEAD,
	}
	return resp, nil
}

func (s *LoginService) loadUserByUsername(username string) (*mymodel.AdminUserDetails, error) {
	admin, err := getAdminByUsername(s.ctx, username)
	if err != nil {
		klog.Errorf("[user_admin-service] Login getAdminByUsername failed, err: %s", err.Error())
		return nil, err
	}
	resourceList, err := s.getResourceList(admin.ID)
	if err != nil {
		klog.Errorf("[user_admin-service] Login getResourceList failed, err: %s", err.Error())
		return nil, err
	}
	return &mymodel.AdminUserDetails{
		UmsAdmin:     admin,
		ResourceList: resourceList,
	}, nil
}

func (s *LoginService) getResourceList(adminId int64) ([]*model.UmsResource, error) {
	// 先从缓存中获取数据
	key := constants.REDIS_DATABASE_UMS + ":" + constants.REDIS_KEY_RESOURCE_LIST + strconv.FormatInt(adminId, 10)
	resourceBytes, err := redisClient.GetDataFromCache(s.ctx, key)
	if err == nil && len(resourceBytes) > 0 {
		var resources []*model.UmsResource
		if err = json.Unmarshal(resourceBytes, &resources); err != nil {
			klog.Errorf("[user_admin-service] Login Redis Get resource list failed, err: %s", err.Error())
			return nil, errno.ServiceInternalError
		}
		return resources, nil
	}
	// 缓存中没有从数据库中获取
	var resources []*model.UmsResource
	resources, err = mymodel.GetResourceListByAdminId(s.ctx, adminId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		klog.Errorf("[user_admin-service] Login GetResourceListByAdminId failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	// 将数据库中的数据存入缓存中
	resourceBytes, _ = json.Marshal(resources)
	if err = redisClient.SetCache(s.ctx, key, resourceBytes); err != nil {
		klog.Errorf("[user_admin-service] Login Redis Set resource list failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	return resources, nil
}

func (s *LoginService) insertLoginLog(username string) error {
	admin, err := getAdminByUsername(s.ctx, username)
	if err != nil {
		klog.Errorf("[user_admin-service] Login getAdminByUsername failed, err: %s", err.Error())
		return err
	}
	if admin == nil {
		return errno.UserNameOrPasswordError
	}
	loginLog := &model.UmsAdminLoginLog{
		AdminID:    admin.ID,
		CreateTime: time.Now(),
		// todo: rpc  没实现 IP 获取
		// IP: s.getClientIP(),
	}
	if err = mymodel.InsertLoginLog(s.ctx, loginLog); err != nil {
		klog.Errorf("[user_admin-service] Login InsertLoginLog failed, err: %s", err.Error())
		return errno.ServiceInternalError
	}
	return nil
}

//// 获取客户端 IP 地址
//func (s *LoginService) getClientIP() string {
//	// 获取 HTTP 请求中的 IP 地址
//	if req, ok := s.ctx.Value("httpRequest").(*http.Request); ok {
//		ip := req.RemoteAddr
//		// 判断是否有代理，如果有代理则返回真实 IP
//		forwarded := req.Header.Get("X-Forwarded-For")
//		if len(forwarded) > 0 {
//			ip = forwarded
//		}
//		return ip
//	}
//	return ""
//}
