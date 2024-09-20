package controllers

import (
	"strconv"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (p *UserApi) PageList(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.BusinessFail(c, err.Error())
	}

	res, err := services.NewUserService().PageList(page.PageIndex, page.PageSize)
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *UserApi) List(c *gin.Context) {
	res, err := services.NewUserService().List()
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *UserApi) Create(c *gin.Context) {
	var req request.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewUserService().Create(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *UserApi) Delete(c *gin.Context) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		response.ValidateFail(c, "")
	}
	if id, err := strconv.ParseInt(idstr, 10, 64); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		if err := services.NewUserService().Delete(id); err != nil {
			response.BusinessFail(c, err.Error())
		} else {
			response.Success[any](c, nil)
		}
	}
}

func (p *UserApi) Update(c *gin.Context) {
	var req request.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewUserService().Update(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
