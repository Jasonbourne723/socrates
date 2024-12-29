package controllers

import (
	"strconv"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
)

type ResourceApi struct {
}

func (r *ResourceApi) Create(c *gin.Context) {
	var Resource request.CreateResource
	if err := c.ShouldBindJSON(&Resource); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(Resource, err))
		return
	}

	if err := services.NewResourceService().Create(Resource); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success[any](c, nil)
	}
}

func (r *ResourceApi) Update(c *gin.Context) {
	var Resource request.UpdateResource
	if err := c.ShouldBindJSON(&Resource); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(Resource, err))
		return
	}

	if err := services.NewResourceService().Update(Resource); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success[any](c, nil)
	}
}

func (r *ResourceApi) Delete(c *gin.Context) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		response.ValidateFail(c, "")
		return
	}
	if id, err := strconv.ParseInt(idstr, 10, 64); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		if err := services.NewResourceService().Delete(id); err != nil {
			response.BusinessFail(c, err.Error())
		} else {
			response.Success[any](c, nil)
		}
	}
}

func (r *ResourceApi) PageList(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	res, err := services.NewResourceService().PageList(page.PageIndex, page.PageSize)
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (r *ResourceApi) List(c *gin.Context) {
	res, err := services.NewResourceService().List()
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (r *ResourceApi) GetOne(c *gin.Context) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		response.ValidateFail(c, "")
		return
	}
	if id, err := strconv.ParseInt(idstr, 10, 64); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		if res, err := services.NewResourceService().GetOne(id); err != nil {
			response.BusinessFail(c, err.Error())
		} else {
			response.Success(c, res)
		}
	}
}
