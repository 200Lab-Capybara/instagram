package builder

import (
	"github.com/gin-gonic/gin"
	"instagram/app/infras/services/profile/repository/mysqlprofile"
	httpprofile "instagram/app/infras/services/profile/transport/http"
	"instagram/app/internals/services/profile/usecaseprofile"
	"instagram/common"
)

func BuildProfileService(con common.SQLDatabase, v1 *gin.RouterGroup) {
	profileStorage := mysqlprofile.NewMySQLStorage(con)
	setupUseCase := usecaseprofile.NewSetupUseCase(profileStorage)
	profileHandler := httpprofile.NewProfileHandler(setupUseCase)
	profileHandler.SetupProfileV1Router(v1)
}
