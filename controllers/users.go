package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/memclutter/gontacts/models"
	"github.com/memclutter/gontacts/services"
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	service *services.Users
}

func NewUsers() *Users {
	return &Users{
		service: services.NewUsers(),
	}
}

func (c *Users) Info(ctx *gin.Context) {
	if claims, ok := ctx.Get("claims"); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	} else if userId := claims.(jwt.MapClaims)["aud"]; !bson.IsObjectIdHex(userId.(string)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	} else if user, err := c.service.Get(bson.ObjectIdHex(userId.(string))); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *Users) Registration(ctx *gin.Context) {
	var model models.User

	if err := ctx.ShouldBind(&model); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if err := c.service.Registration(&model); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, model)
	}
}

func (c *Users) Login(ctx *gin.Context) {
	var model models.Login

	if err := ctx.ShouldBind(&model); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if token, err := c.service.Login(&model); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"token": token,
		})
	}
}

func (c *Users) Confirmation(ctx *gin.Context) {
	if token, ok := ctx.GetQuery("token"); !ok {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	} else if err := c.service.Confirmation(token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "User confirmation success",
		})
	}
}
