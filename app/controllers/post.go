package controllers

import (
	"strconv"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
)

type PostApi struct {
}

func (p *PostApi) PageList(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.BusinessFail(c, err.Error())
	}

	res, err := services.NewPostService().PageList(page.PageIndex, page.PageSize)
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *PostApi) List(c *gin.Context) {
	res, err := services.NewPostService().List()
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *PostApi) Create(c *gin.Context) {
	var req request.CreatePost
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewPostService().Create(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *PostApi) Delete(c *gin.Context) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		response.ValidateFail(c, "")
	}
	if id, err := strconv.ParseInt(idstr, 10, 64); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		if err := services.NewPostService().Delete(id); err != nil {
			response.BusinessFail(c, err.Error())
		} else {
			response.Success[any](c, nil)
		}
	}
}

func (p *PostApi) Update(c *gin.Context) {
	var req request.UpdatePost
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewPostService().Update(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
