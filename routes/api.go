package routes

import (
	"github.com/Jasonbourne723/socrates/app/controllers"
	"github.com/Jasonbourne723/socrates/app/controllers/app"
	"github.com/Jasonbourne723/socrates/app/controllers/common"
	"github.com/Jasonbourne723/socrates/app/middleware"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {

	router.Use(middleware.Cors()).OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)

	})

	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}

	roleRouter := router.Group("").Use(middleware.Cors())
	{
		roleApi := &controllers.RoleApi{}
		roleRouter.POST("/role", roleApi.Create)
		roleRouter.DELETE("/role/:id", roleApi.Delete)
		roleRouter.GET("/role/pagelist", roleApi.PageList)
		roleRouter.PUT("/role", roleApi.Update)
	}

	permissionSpaceRouter := router.Group("").Use(middleware.Cors())
	{
		permissionSpaceApi := &controllers.PermissionSpaceApi{}
		permissionSpaceRouter.POST("/permission_space", permissionSpaceApi.Create)
		permissionSpaceRouter.DELETE("/permission_space/:id", permissionSpaceApi.Delete)
		permissionSpaceRouter.PUT("/permission_space", permissionSpaceApi.Update)
		permissionSpaceRouter.GET("/permission_space/pagelist", permissionSpaceApi.PageList)
		permissionSpaceRouter.GET("/permission_space", permissionSpaceApi.List)
	}

	organizationRouter := router.Group("").Use(middleware.Cors())
	{
		organizationApi := &controllers.OrganizationApi{}

		organizationRouter.POST("/organization", organizationApi.Create)
		organizationRouter.DELETE("/organization/:id", organizationApi.Delete)
		organizationRouter.GET("/organization", organizationApi.List)
		organizationRouter.PUT("/organization", organizationApi.Update)
	}
}
