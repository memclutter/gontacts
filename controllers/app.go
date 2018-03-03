package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time": time.Now().UTC(),
	})
}
