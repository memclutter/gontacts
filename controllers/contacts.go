package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Contacts struct {
}

func (c *Contacts) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []gin.H{
		{
			"id": 1,
			"full_name": "John Doe",
			"phone": "+728328382",
			"email": "john.doe@gmail.com",
		},
	})
}

func (c *Contacts) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"id": 1,
		"full_name": "John Doe",
		"phone": "+728328382",
		"email": "john.doe@gmail.com",
	})
}

func (c *Contacts) Show(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"id": 1,
		"full_name": "John Doe",
		"phone": "+728328382",
		"email": "john.doe@gmail.com",
	})
}

func (c *Contacts) Patch(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"id": 1,
		"full_name": "John Doe",
		"phone": "+728328382",
		"email": "john.doe@gmail.com",
	})
}

func (c *Contacts) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"id": 1,
		"full_name": "John Doe",
		"phone": "+728328382",
		"email": "john.doe@gmail.com",
	})
}

func (c *Contacts) Destroy(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}
