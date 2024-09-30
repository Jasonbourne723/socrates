//go:build wireinject

package inject

import (
	"github.com/Jasonbourne723/socrates/app/controllers"
	"github.com/Jasonbourne723/socrates/app/services"
	"github.com/google/wire"
)

func InitializeAuthApi() controllers.AuthApi {
	wire.Build(controllers.NewAuthApi, services.NewUserService)
	return controllers.AuthApi{}
}
