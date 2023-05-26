package main

import (
	"npv-application/entity"
	"npv-application/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var npv entity.NPV

	routs := routers.Npv{
		R:         r,
		NpvEntity: npv,
	}
	routs.Routers()

	r.Run()
}
