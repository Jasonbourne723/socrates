package services

import (
	"math"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
	"github.com/Jasonbourne723/socrates/utils"
)

type SsoApplicationService struct {
}

type ISsoApplicationService interface {
	PageList(pageIndex int32, pageSize int32) (pages response.Page[response.SsoApplication], err error)
	List() (list []*response.SsoApplication, err error)
	Create(req *request.CreateSsoApplication) (res *response.SsoApplication, err error)
	Update(req *request.UpdateSsoApplication) (res *response.SsoApplication, err error)
	GetByName(name string) (response *response.SsoApplication, err error)
	Delete(id int64) (err error)
}

func NewSsoApplicationService() ISsoApplicationService {
	return &SsoApplicationService{}
}

func (p *SsoApplicationService) PageList(pageIndex int32, pageSize int32) (pages response.Page[response.SsoApplication], err error) {

	var SsoApplications []models.SsoApplication
	if err = global.App.DB.Offset((int(pageIndex) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&SsoApplications).Error; err != nil {
		return
	}

	rows := []response.SsoApplication{}
	for _, item := range SsoApplications {
		rows = append(rows, *MapToSsoApplicationResponse(&item))
	}

	var count int64
	err = global.App.DB.Find(&models.Role{}).Count(&count).Error

	pages = response.Page[response.SsoApplication]{
		Rows:       rows,
		PageIndex:  int32(pageIndex),
		PageSize:   int32(pageSize),
		TotalCount: count,
		TotalPage:  int64(math.Ceil(float64(count) / float64(pageSize))),
	}

	return
}

func (p *SsoApplicationService) List() (roots []*response.SsoApplication, err error) {
	var SsoApplications []models.SsoApplication
	if err = global.App.DB.Find(&SsoApplications).Error; err != nil {
		return
	}
	roots = []*response.SsoApplication{}
	for i := range SsoApplications {
		entity := MapToSsoApplicationResponse(&SsoApplications[i])
		roots = append(roots, entity)
	}
	return
}

func (p *SsoApplicationService) GetByName(name string) (response *response.SsoApplication, err error) {
	var SsoApplications models.SsoApplication
	if err = global.App.DB.Where("name = ?", name).First(&SsoApplications).Error; err != nil {
		return
	}
	response = MapToSsoApplicationResponse(&SsoApplications)
	return
}

func (p *SsoApplicationService) Create(req *request.CreateSsoApplication) (res *response.SsoApplication, err error) {
	exists := models.SsoApplication{}
	result := global.App.DB.Where("name = ?", req.Name).First(&exists)
	if result.RowsAffected != 0 {
		return nil, global.Errors.NameDuplicateError
	}
	appKey, _ := utils.GenerateRandomString(8)
	appSecret, _ := utils.GenerateRandomString(24)
	entity := models.SsoApplication{Name: req.Name, Description: req.Description, CallbackUrl: req.CallbackUrl, AppKey: appKey, AppSecret: appSecret}
	err = global.App.DB.Create(&entity).Error
	res = MapToSsoApplicationResponse(&entity)
	return
}

func (p *SsoApplicationService) Update(req *request.UpdateSsoApplication) (res *response.SsoApplication, err error) {
	var exists models.SsoApplication
	result := global.App.DB.First(&exists, req.Id)
	if result.Error != nil {
		return
	}
	if result.RowsAffected == 0 {
		err = global.Errors.RecordNotFoundError
		return
	}

	exists.Name = req.Name
	exists.Description = req.Description

	err = global.App.DB.Save(&exists).Error
	if err != nil {
		return
	}
	res = MapToSsoApplicationResponse(&exists)
	return
}

func (p *SsoApplicationService) Delete(id int64) (err error) {
	err = global.App.DB.Delete(&models.SsoApplication{}, id).Error
	return err
}

func MapToSsoApplicationResponse(m *models.SsoApplication) *response.SsoApplication {
	return &response.SsoApplication{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
		AppKey:      m.AppKey,
		AppSecret:   m.AppSecret,
		CallbackUrl: m.CallbackUrl,
	}
}
