package routes

import (
	"github.com/Jasonbourne723/socrates/app/controllers"
	"github.com/Jasonbourne723/socrates/app/controllers/common"
	"github.com/Jasonbourne723/socrates/app/middleware"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {

	router.Use(middleware.Cors()).OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)

	})

	authApi := &controllers.AuthApi{}

	unauthRouter := router.Group("").Use(middleware.Cors())
	{
		unauthRouter.POST("/auth/register", authApi.Register)
		unauthRouter.POST("/auth/login", authApi.Login)
	}

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName)).Use(middleware.Cors())
	{

		authRouter.POST("/auth/info", authApi.Info)
		authRouter.POST("/auth/logout", authApi.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}

	roleRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName)).Use(middleware.Cors())
	{
		roleApi := &controllers.RoleApi{}
		roleRouter.POST("/role", roleApi.Create)
		roleRouter.DELETE("/role/:id", roleApi.Delete)
		roleRouter.GET("/role/pagelist", roleApi.PageList)
		roleRouter.PUT("/role", roleApi.Update)
	}

	permissionSpaceRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName)).Use(middleware.Cors())
	{
		permissionSpaceApi := &controllers.PermissionSpaceApi{}
		permissionSpaceRouter.POST("/permission_space", permissionSpaceApi.Create)
		permissionSpaceRouter.DELETE("/permission_space/:id", permissionSpaceApi.Delete)
		permissionSpaceRouter.PUT("/permission_space", permissionSpaceApi.Update)
		permissionSpaceRouter.GET("/permission_space/pagelist", permissionSpaceApi.PageList)
		permissionSpaceRouter.GET("/permission_space", permissionSpaceApi.List)
	}

	organizationRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName)).Use(middleware.Cors())
	{
		organizationApi := &controllers.OrganizationApi{}

		organizationRouter.POST("/organization", organizationApi.Create)
		organizationRouter.DELETE("/organization/:id", organizationApi.Delete)
		organizationRouter.GET("/organization", organizationApi.List)
		organizationRouter.PUT("/organization", organizationApi.Update)
	}
}
