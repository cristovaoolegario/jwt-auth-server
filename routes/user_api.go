package routes

import (
	"github.com/cristovaoolegario/free-auth-server/service"
	"net/http"
)

type UserAPI struct {
	UserService service.UserService
}

func ProvideUserAPI(s service.UserService) UserAPI {
	return UserAPI{UserService: s}
}

func (api *UserAPI) Register(w http.ResponseWriter, r *http.Request) {

}
