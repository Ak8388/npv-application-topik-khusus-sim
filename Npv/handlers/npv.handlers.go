package handlers

import (
	"net/http"
	"npv-application/npv"

	"github.com/gin-gonic/gin"
)

func NewHandlersNpv(npvUs npv.UsecaseNpv, r *gin.RouterGroup) {
	eng := npvHandlers{
		usecaseNpv: npvUs,
	}

	v2 := r.Group("npv")
	v2.POST("", eng.PostNpv)
	v2.GET("", eng.GetNpv)
}

type npvHandlers struct {
	usecaseNpv npv.UsecaseNpv
}

func (hand npvHandlers) PostNpv(ctx *gin.Context) {
	err := hand.usecaseNpv.PostNpv(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (hand npvHandlers) GetNpv(ctx *gin.Context) {
	res, err := hand.usecaseNpv.GetNpv(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Result": res})
}
