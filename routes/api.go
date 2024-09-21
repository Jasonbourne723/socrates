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
		roleRouter.GET("/role", roleApi.List)
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
		organizationRouter.GET("/organization/all", organizationApi.All)
		organizationRouter.PUT("/organization", organizationApi.Update)
	}

	applicationRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName)).Use(middleware.Cors())
	{
		applicationApi := &controllers.ApplicationApi{}
		applicationRouter.POST("/application", applicationApi.Create)
		applicationRouter.DELETE("/application/:id", applicationApi.Delete)
		applicationRouter.GET("/application", applicationApi.List)
		applicationRouter.PUT("/application", applicationApi.Update)
		applicationRouter.GET("/application/pagelist", applicationApi.PageList)
	}

	postRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName)).Use(middleware.Cors())
	{
		postApi := &controllers.PostApi{}
		postRouter.POST("/post", postApi.Create)
		postRouter.DELETE("/post/:id", postApi.Delete)
		postRouter.GET("/post", postApi.List)
		postRouter.PUT("/post", postApi.Update)
		postRouter.GET("/post/pagelist", postApi.PageList)
	}

	userRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName)).Use(middleware.Cors())
	{
		userApi := &controllers.UserApi{}
		userRouter.POST("/user", userApi.Create)
		userRouter.DELETE("/user/:id", userApi.Delete)
		userRouter.GET("/user", userApi.List)
		userRouter.PUT("/user", userApi.Update)
		userRouter.GET("/user/pagelist", userApi.PageList)
	}

	resourceRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName)).Use(middleware.Cors())
	{
		resourceApi := &controllers.ResourceApi{}
		resourceRouter.POST("/resource", resourceApi.Create)
		resourceRouter.DELETE("/resource/:id", resourceApi.Delete)
		resourceRouter.GET("/resource", resourceApi.List)
		resourceRouter.PUT("/resource", resourceApi.Update)
		resourceRouter.GET("/resource/pagelist", resourceApi.PageList)
		resourceRouter.GET("/resource/:id", resourceApi.GetOne)
	}
}
