package controllers

import (
	"strconv"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
)

type ApplicationApi struct {
}

func (p *ApplicationApi) PageList(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.BusinessFail(c, err.Error())
	}

	res, err := services.NewApplicationService().PageList(page.PageIndex, page.PageSize)
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *ApplicationApi) List(c *gin.Context) {
	res, err := services.NewApplicationService().List()
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *ApplicationApi) Create(c *gin.Context) {
	var req request.CreateApplication
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewApplicationService().Create(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *ApplicationApi) Delete(c *gin.Context) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		response.ValidateFail(c, "")
	}
	if id, err := strconv.ParseInt(idstr, 10, 64); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		if err := services.NewApplicationService().Delete(id); err != nil {
			response.BusinessFail(c, err.Error())
		} else {
			response.Success[any](c, nil)
		}
	}
}

func (p *ApplicationApi) Update(c *gin.Context) {
	var req request.UpdateApplication
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewApplicationService().Update(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
