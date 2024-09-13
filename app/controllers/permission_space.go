package controllers

import (
	"strconv"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
)

type PermissionSpaceApi struct {
}

func (p *PermissionSpaceApi) PageList(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.BusinessFail(c, err.Error())
	}

	res, err := services.NewPermissionSpaceService().PageList(page.PageIndex, page.PageSize)
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *PermissionSpaceApi) List(c *gin.Context) {
	res, err := services.NewPermissionSpaceService().List()
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *PermissionSpaceApi) Create(c *gin.Context) {
	var req request.CreatePermissionSpace
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewPermissionSpaceService().Create(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *PermissionSpaceApi) Delete(c *gin.Context) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		response.ValidateFail(c, "")
	}
	if id, err := strconv.ParseInt(idstr, 10, 64); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		if err := services.NewPermissionSpaceService().Delete(id); err != nil {
			response.BusinessFail(c, err.Error())
		} else {
			response.Success[any](c, nil)
		}
	}
}

func (p *PermissionSpaceApi) Update(c *gin.Context) {
	var req request.UpdatePermissionSpace
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewPermissionSpaceService().Update(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
