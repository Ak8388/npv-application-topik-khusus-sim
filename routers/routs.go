package routers

import (
	"npv-application/Npv/handlers"
	"npv-application/Npv/repository"
	"npv-application/Npv/usecase"
	"npv-application/entity"
	"npv-application/middleware"

	"github.com/gin-gonic/gin"
)

type Npv struct {
	R         *gin.Engine
	NpvEntity entity.NPV
}

func (npv Npv) Routers() {
	v1 := npv.R.Group("api")

	middleware.Add(npv.R, middleware.CORSMiddleware())
	repoNpv := repository.NewRepoNpv(npv.NpvEntity)
	usecaseNpv := usecase.NewUsecaseNpv(&repoNpv)
	handlers.NewHandlersNpv(usecaseNpv, v1)
}
