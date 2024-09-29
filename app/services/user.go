package services

import (
	"errors"
	"math"
	"strconv"

	"github.com/Jasonbourne723/socrates/app/common/mapster"
	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
	"github.com/Jasonbourne723/socrates/utils"
	"gorm.io/gorm"
)

type userService struct {
}

type IUserService interface {
	PageList(pageIndex int32, pageSize int32) (pages response.Page[response.User], err error)
	List() (list []*response.User, err error)
	Create(req *request.CreateUser) (res *response.User, err error)
	Update(req *request.UpdateUser) (res *response.User, err error)
	Delete(id int64) (err error)
}

func NewUserService() IUserService {
	return &userService{}
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password)), Avatar: ""}
	err = global.App.DB.Create(&user).Error
	return
}

func (userService *userService) RegisterByGitHub(params response.GitHubUser) (user models.User, err error) {

	user = models.User{}
	var result = global.App.DB.Where("github_openid = ?", params.ID).Select("id").First(&user)
	if result.RowsAffected != 0 {
		return
	}
	user = models.User{Name: params.Name, Mobile: "", Password: "", GithubOpenid: params.ID, Avatar: params.AvatarURL}
	err = global.App.DB.Create(&user).Error
	return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("mobile = ?", params.Mobile).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

func (p *userService) PageList(pageIndex int32, pageSize int32) (pages response.Page[response.User], err error) {

	var Users []models.User
	if err = global.App.DB.Offset((int(pageIndex) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&Users).Error; err != nil {
		return
	}

	rows := []response.User{}
	for _, user := range Users {
		userRoles := []models.UserRole{}
		if err = global.App.DB.Where("user_id = ?", user.Id).Find(&userRoles).Error; err != nil {
			return
		}
		roleIds := mapster.Map(userRoles, func(t models.UserRole) int64 { return t.RoleId })
		userPost := models.UserPost{}
		if err = global.App.DB.Where("user_id = ?", user.Id).First(&userPost).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		userOrganization := models.UserOrganization{}
		if err = global.App.DB.Where("user_id = ?", user.Id).First(&userOrganization).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		entity := MapToUserResponse(&user, roleIds, userOrganization.OrganizationId, userPost.PostId)
		rows = append(rows, *entity)
	}

	var count int64
	err = global.App.DB.Find(&models.Role{}).Count(&count).Error

	pages = response.Page[response.User]{
		Rows:       rows,
		PageIndex:  int32(pageIndex),
		PageSize:   int32(pageSize),
		TotalCount: count,
		TotalPage:  int64(math.Ceil(float64(count) / float64(pageSize))),
	}

	return
}

func (p *userService) List() (list []*response.User, err error) {
	var Users []models.User
	if err = global.App.DB.Find(&Users).Error; err != nil {
		return
	}
	list = []*response.User{}
	for _, user := range Users {
		userRoles := []models.UserRole{}
		if err = global.App.DB.Where("user_id = ?", user.Id).Find(&userRoles).Error; err != nil {
			return
		}
		roleIds := mapster.Map(userRoles, func(t models.UserRole) int64 { return t.RoleId })
		userPost := models.UserPost{}
		if err = global.App.DB.Where("user_id = ?", user.Id).First(&userPost).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		userOrganization := models.UserOrganization{}
		if err = global.App.DB.Where("user_id = ?", user.Id).First(&userOrganization).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		entity := MapToUserResponse(&user, roleIds, userOrganization.OrganizationId, userPost.PostId)
		list = append(list, entity)
	}
	return
}

func (p *userService) Create(req *request.CreateUser) (res *response.User, err error) {
	exists := models.User{}
	result := global.App.DB.Where("mobile = ?", req.Mobile).First(&exists)
	if result.RowsAffected != 0 {
		return nil, global.Errors.MobileExistedError
	}
	entity := models.User{Name: req.Name, Mobile: req.Mobile, Password: utils.BcryptMake([]byte("123456"))}
	if err = global.App.DB.Create(&entity).Error; err != nil {
		return
	}
	if req.PostId > 0 {
		if err = global.App.DB.Create(&models.UserPost{UserId: entity.Id, PostId: req.PostId}).Error; err != nil {
			return
		}
	}

	if req.OrganizationId > 0 {
		if err = global.App.DB.Create(&models.UserOrganization{UserId: entity.Id, OrganizationId: req.OrganizationId}).Error; err != nil {
			return
		}
	}

	if len(req.RoleIds) > 0 {
		roles := mapster.Map(req.RoleIds, func(roleId int64) models.UserRole {
			return models.UserRole{
				RoleId: roleId,
				UserId: entity.Id,
			}
		})
		if err = global.App.DB.Create(&roles).Error; err != nil {
			return
		}
	}

	res = MapToUserResponse(&entity, req.RoleIds, req.OrganizationId, req.PostId)
	return
}

func (p *userService) Update(req *request.UpdateUser) (res *response.User, err error) {
	var exists models.User
	result := global.App.DB.First(&exists, req.Id)
	if result.Error != nil {
		return
	}
	if result.RowsAffected == 0 {
		err = global.Errors.RecordNotFoundError
		return
	}

	exists.Name = req.Name
	exists.Mobile = req.Mobile
	err = global.App.DB.Save(&exists).Error
	if err != nil {
		return
	}

	global.App.DB.Where("user_id = ?", req.Id).Delete(&models.UserPost{})
	if req.PostId > 0 {
		if err = global.App.DB.Create(&models.UserPost{UserId: req.Id, PostId: req.PostId}).Error; err != nil {
			return
		}
	}
	global.App.DB.Where("user_id = ?", req.Id).Delete(&models.UserOrganization{})
	if req.OrganizationId > 0 {
		if err = global.App.DB.Create(&models.UserOrganization{UserId: req.Id, OrganizationId: req.OrganizationId}).Error; err != nil {
			return
		}
	}
	global.App.DB.Where("user_id = ?", req.Id).Delete(&models.UserRole{})
	if len(req.RoleIds) > 0 {
		roles := mapster.Map(req.RoleIds, func(roleId int64) models.UserRole {
			return models.UserRole{
				RoleId: roleId,
				UserId: req.Id,
			}
		})
		if err = global.App.DB.Create(&roles).Error; err != nil {
			return
		}
	}
	res = MapToUserResponse(&exists, req.RoleIds, req.OrganizationId, req.PostId)
	return
}

func (p *userService) Delete(id int64) (err error) {
	err = global.App.DB.Delete(&models.User{}, id).Error
	return err
}

func MapToUserResponse(m *models.User, roleIds []int64, organizationId int64, postId int64) *response.User {
	return &response.User{
		Id:             m.Id,
		Name:           m.Name,
		Mobile:         m.Mobile,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		RoleIds:        roleIds,
		OrganizationId: organizationId,
		PostId:         postId,
	}
}
