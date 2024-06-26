package httpprofile

import "instagram/app/internals/services/profile/usecaseprofile"

type profileHandler struct {
	setupUseCase usecaseprofile.SetupUseCase
}

func NewProfileHandler(setupUseCase usecaseprofile.SetupUseCase) *profileHandler {
	return &profileHandler{setupUseCase: setupUseCase}
}
