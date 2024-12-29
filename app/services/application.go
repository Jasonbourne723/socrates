package services

import (
	"math"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
	"github.com/Jasonbourne723/socrates/utils"
)

type ApplicationService struct {
}

type IApplicationService interface {
	PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Application], err error)
	List() (list []*response.Application, err error)
	Create(req *request.CreateApplication) (res *response.Application, err error)
	Update(req *request.UpdateApplication) (res *response.Application, err error)
	Delete(id int64) (err error)
}

func NewApplicationService() IApplicationService {
	return &ApplicationService{}
}

func (p *ApplicationService) PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Application], err error) {

	var applications []models.Application
	if err = global.App.DB.Offset((int(pageIndex) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&applications).Error; err != nil {
		return
	}

	rows := []response.Application{}
	for _, item := range applications {
		rows = append(rows, *MapToApplicationResponse(&item))
	}

	var count int64
	err = global.App.DB.Find(&models.Role{}).Count(&count).Error

	pages = response.Page[response.Application]{
		Rows:       rows,
		PageIndex:  int32(pageIndex),
		PageSize:   int32(pageSize),
		TotalCount: count,
		TotalPage:  int64(math.Ceil(float64(count) / float64(pageSize))),
	}

	return
}

func (p *ApplicationService) List() (roots []*response.Application, err error) {
	var applications []models.Application
	if err = global.App.DB.Find(&applications).Error; err != nil {
		return
	}
	roots = []*response.Application{}
	for i := range applications {
		entity := MapToApplicationResponse(&applications[i])
		roots = append(roots, entity)
	}
	return
}

func (p *ApplicationService) Create(req *request.CreateApplication) (res *response.Application, err error) {
	exists := models.Application{}
	result := global.App.DB.Where("name = ?", req.Name).First(&exists)
	if result.RowsAffected != 0 {
		return nil, global.Errors.NameDuplicateError
	}
	appKey, _ := utils.GenerateRandomString(8)
	appSecret, _ := utils.GenerateRandomString(24)
	entity := models.Application{Name: req.Name, Description: req.Description, CallbackUrl: req.CallbackUrl, AppKey: appKey, AppSecret: appSecret}
	err = global.App.DB.Create(&entity).Error
	res = MapToApplicationResponse(&entity)
	return
}

func (p *ApplicationService) Update(req *request.UpdateApplication) (res *response.Application, err error) {
	var exists models.Application
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
	res = MapToApplicationResponse(&exists)
	return
}

func (p *ApplicationService) Delete(id int64) (err error) {
	err = global.App.DB.Delete(&models.Application{}, id).Error
	return err
}

func MapToApplicationResponse(m *models.Application) *response.Application {
	return &response.Application{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
		AppKey:      m.AppKey,
		AppSecret:   m.AppSecret,
		CallbackUrl: m.CallbackUrl,
	}
}
