package npv

import (
	"npv-application/entity"

	"github.com/gin-gonic/gin"
)

type UsecaseNpv interface {
	PostNpv(*gin.Context) error
	GetNpv(*gin.Context) (entity.NPV, error)
}

type RepoNpv interface {
	PostNpv(entity.Input) error
	GetNpv() (entity.NPV, error)
}
