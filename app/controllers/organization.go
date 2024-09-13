package controllers

import (
	"strconv"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
)

type OrganizationApi struct {
}

func (p *OrganizationApi) List(c *gin.Context) {

	res, err := services.NewOrganizationService().List()
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *OrganizationApi) Create(c *gin.Context) {
	var req request.CreateOrganization
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewOrganizationService().Create(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (p *OrganizationApi) Delete(c *gin.Context) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		response.ValidateFail(c, "")
	}
	if id, err := strconv.ParseInt(idstr, 10, 64); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		if err := services.NewOrganizationService().Delete(id); err != nil {
			response.BusinessFail(c, err.Error())
		} else {
			response.Success[any](c, nil)
		}
	}
}

func (p *OrganizationApi) Update(c *gin.Context) {
	var req request.UpdateOrganization
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(req, err))
	}

	if res, err := services.NewOrganizationService().Update(&req); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
