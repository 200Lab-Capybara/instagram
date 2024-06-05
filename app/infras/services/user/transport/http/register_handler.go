package userhttp

import (
	usermodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/model"
	"github.com/nghiatrann0502/instagram-clone/common"
	"net/http"
)

func (u *userHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	dataCreation := usermodel.UserCreation{}

	err := common.ReadJSON(r, &dataCreation)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := dataCreation.Validate(); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := u.userUseCase.Register(r.Context(), &dataCreation)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	common.SimpleSuccess(w, http.StatusOK, id)
}
