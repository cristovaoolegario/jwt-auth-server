package routes

import (
	"encoding/json"
	"github.com/cristovaoolegario/free-auth-server/dto"
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
	defer r.Body.Close()
	var user dto.InsertUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := user.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := api.UserService.GetUserByEmail(user.Email); err == nil {
		respondWithError(w, http.StatusBadRequest, "This email is already being used by an account.")
		return
	}

	createdUser, err := api.UserService.CreateNewUser(user)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, createdUser)
}
