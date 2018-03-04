package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/memclutter/gontacts/components"
	"github.com/memclutter/gontacts/models"
	"github.com/memclutter/gontacts/services"
	"gopkg.in/mgo.v2/bson"
)

type Contacts struct {
	service *services.Contacts
}

func NewContacts() *Contacts {
	return &Contacts{
		service: services.NewContacts(),
	}
}

func (c *Contacts) Index(ctx *gin.Context) {
	var pagination components.Pagination

	ctx.ShouldBindQuery(&pagination)

	if model, totalCount, err := c.service.All(pagination.Offset, pagination.Limit, pagination.Order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"items":       model,
			"total_count": totalCount,
		})
	}
}

func (c *Contacts) Create(ctx *gin.Context) {
	var model models.Contact

	if err := ctx.ShouldBind(&model); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if err := c.service.Create(&model); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, model)
	}
}

func (c *Contacts) Show(ctx *gin.Context) {
	if !bson.IsObjectIdHex(ctx.Param("id")) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	} else {
		id := bson.ObjectIdHex(ctx.Param("id"))

		if model, err := c.service.Get(id); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, model)
		}
	}
}

//func (c *Contacts) PartialUpdate(ctx *gin.Context) {
//	var model models.Contact
//
//	if !bson.IsObjectIdHex(ctx.Param("id")) {
//		ctx.JSON(http.StatusNotFound, gin.H{
//			"message": "Page not found",
//		})
//	} else {
//		id := bson.ObjectIdHex(ctx.Param("id"))
//
//		ctx.ShouldBind(&model)
//
//		if err := c.service.PartialUpdate(id, &model); err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"message": err.Error(),
//			})
//		} else {
//			ctx.JSON(http.StatusOK, model)
//		}
//	}
//}

func (c *Contacts) Update(ctx *gin.Context) {
	var model models.Contact

	if !bson.IsObjectIdHex(ctx.Param("id")) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	} else {
		id := bson.ObjectIdHex(ctx.Param("id"))

		if err := ctx.ShouldBind(&model); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else if err := c.service.Update(id, &model); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, model)
		}
	}
}

func (c *Contacts) Destroy(ctx *gin.Context) {
	if !bson.IsObjectIdHex(ctx.Param("id")) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	} else {
		id := bson.ObjectIdHex(ctx.Param("id"))

		if err := c.service.Destroy(id); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}
