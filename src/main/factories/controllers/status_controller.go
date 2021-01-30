package fctrls

import (
	"lucaswilliameufrasio/golang-fiber-api/src/presentation/controllers"
	presprotcls "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

func MakeStatusController() presprotcls.Controller {
	return controllers.NewStatusController()
}
