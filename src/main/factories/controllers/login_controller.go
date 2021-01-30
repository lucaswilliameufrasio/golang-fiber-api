package fctrls

import (
	factories "lucaswilliameufrasio/golang-fiber-api/src/main/factories/usecases"
	"lucaswilliameufrasio/golang-fiber-api/src/presentation/controllers"
	presprotcls "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

func MakeLoginController() presprotcls.Controller {
	return controllers.NewLoginController(factories.MakeDbAuthentication())
}
