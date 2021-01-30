package fctrls

import (
	"lucaswilliameufrasio/golang-fiber-api/src/presentation/controllers"
	presprotcls "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

func MakeProtectedControler() presprotcls.Controller {
	return controllers.NewProtectedController()
}
