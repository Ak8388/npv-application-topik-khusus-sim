package middleware

import (
	"github.com/gin-gonic/gin"
)

func Add(r *gin.Engine, h ...gin.HandlerFunc) {
	for _, v := range h {
		r.Use(v)
	}
}
