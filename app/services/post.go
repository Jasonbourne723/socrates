package services

import (
	"math"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
)

type PostService struct {
}

type IPostService interface {
	PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Post], err error)
	List() (list []*response.Post, err error)
	Create(req *request.CreatePost) (res *response.Post, err error)
	Update(req *request.UpdatePost) (res *response.Post, err error)
	Delete(id int64) (err error)
}

func NewPostService() IPostService {
	return &PostService{}
}

func (p *PostService) PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Post], err error) {

	var Posts []models.Post
	if err = global.App.DB.Offset((int(pageIndex) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&Posts).Error; err != nil {
		return
	}

	rows := []response.Post{}
	for _, item := range Posts {
		postOrganizations := []models.PostOrganization{}
		err = global.App.DB.Where("post_id = ?", item.Id).Find(&postOrganizations).Error
		if err != nil {
			return
		}
		organizationIds := []int64{}
		for _, item := range postOrganizations {
			organizationIds = append(organizationIds, item.OrganizationId)
		}
		rows = append(rows, *MapToPostResponse(&item, organizationIds))
	}

	var count int64
	err = global.App.DB.Find(&models.Role{}).Count(&count).Error

	pages = response.Page[response.Post]{
		Rows:       rows,
		PageIndex:  int32(pageIndex),
		PageSize:   int32(pageSize),
		TotalCount: count,
		TotalPage:  int64(math.Ceil(float64(count) / float64(pageSize))),
	}

	return
}

func (p *PostService) List() (list []*response.Post, err error) {
	var Posts []models.Post
	if err = global.App.DB.Find(&Posts).Error; err != nil {
		return
	}
	list = []*response.Post{}
	for i := range Posts {
		postOrganizations := []models.PostOrganization{}
		global.App.DB.Select("id").Where("post_id = ?", Posts[i].Id).Find(&postOrganizations)
		organizationIds := []int64{}
		for _, item := range postOrganizations {
			organizationIds = append(organizationIds, item.OrganizationId)
		}
		entity := MapToPostResponse(&Posts[i], organizationIds)
		list = append(list, entity)
	}
	//todo: organization_post
	return
}

func (p *PostService) Create(req *request.CreatePost) (res *response.Post, err error) {
	exists := models.Post{}
	result := global.App.DB.Where("name = ?", req.Name).Or("code = ?", req.Code).First(&exists)
	if result.RowsAffected != 0 {
		if exists.Code == req.Code {
			return nil, global.Errors.CodeDuplicateError
		} else {
			return nil, global.Errors.NameDuplicateError
		}
	}
	entity := models.Post{Name: req.Name, Code: req.Code}
	//todo: organization_post
	err = global.App.DB.Create(&entity).Error
	if req.OrganizationIds != nil && len(req.OrganizationIds) > 0 {
		organizationPosts := []models.PostOrganization{}
		for _, item := range req.OrganizationIds {
			organizationPosts = append(organizationPosts, models.PostOrganization{OrganizationId: item, PostId: entity.Id})
		}
		global.App.DB.Create(organizationPosts)
	}
	res = MapToPostResponse(&entity, req.OrganizationIds)
	return
}

func (p *PostService) Update(req *request.UpdatePost) (res *response.Post, err error) {
	var exists models.Post
	result := global.App.DB.First(&exists, req.Id)
	if result.Error != nil {
		return
	}
	if result.RowsAffected == 0 {
		err = global.Errors.RecordNotFoundError
		return
	}

	exists.Name = req.Name
	exists.Code = req.Code
	err = global.App.DB.Save(&exists).Error
	if err != nil {
		return
	}
	var postOrganizations = []models.PostOrganization{}

	err = global.App.DB.Where("post_id = ?", exists.Id).Find(&postOrganizations).Error
	if err != nil {
		return
	}
	if len(postOrganizations) > 0 {
		err = global.App.DB.Delete(&postOrganizations).Error
		if err != nil {
			return
		}
	}
	if req.OrganizationIds != nil && len(req.OrganizationIds) > 0 {
		organizationPosts := []models.PostOrganization{}
		for _, item := range req.OrganizationIds {
			organizationPosts = append(organizationPosts, models.PostOrganization{OrganizationId: item, PostId: exists.Id})
		}
		global.App.DB.Create(organizationPosts)
	}
	res = MapToPostResponse(&exists, req.OrganizationIds)
	return
}

func (p *PostService) Delete(id int64) (err error) {
	err = global.App.DB.Delete(&models.Post{}, id).Error
	return err
}

func MapToPostResponse(m *models.Post, organizationIds []int64) *response.Post {
	return &response.Post{
		Id:              m.Id,
		Name:            m.Name,
		Code:            m.Code,
		OrganizationIds: organizationIds,
	}
}
