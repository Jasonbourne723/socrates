package controllers

import (
	"strconv"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
)

type RoleApi struct {
}

// Role 创建角色
// @Summary 创建角色
// @Description 创建角色
// @Tags 角色
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param default body request.CreaeteRole true "参数"
// @Success 200 {object} response.Response[dto.ProjectDto]
// @Router /role [post]
func (r *RoleApi) Create(c *gin.Context) {
	var role request.CreaeteRole
	if err := c.ShouldBindJSON(&role); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(role, err))
	}

	if res, err := services.NewRoleService().Create(role); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (r *RoleApi) Update(c *gin.Context) {
	var role request.UpdateRole
	if err := c.ShouldBindJSON(&role); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(role, err))
	}

	if res, err := services.NewRoleService().Update(role); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func (r *RoleApi) Delete(c *gin.Context) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		response.ValidateFail(c, "")
	}
	if id, err := strconv.ParseInt(idstr, 10, 64); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		if err := services.NewRoleService().Delete(id); err != nil {
			response.BusinessFail(c, err.Error())
		} else {
			response.Success[any](c, nil)
		}
	}
}

func (r *RoleApi) PageList(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.BusinessFail(c, err.Error())
	}

	res, err := services.NewRoleService().PageList(page.PageIndex, page.PageSize)
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
