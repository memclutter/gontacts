package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/memclutter/gontacts/components"
	"github.com/memclutter/gontacts/controllers"
	"github.com/memclutter/gontacts/utils"
)

var (
	addr     string
	mongoUrl string
)

func init() {
	utils.FlagStringVarEnv(&addr, "addr", ":8000", "Set address", "ADDR")
	utils.FlagStringVarEnv(&mongoUrl, "mongoUrl", "mongodb://localhost:27017/gontacts", "MongoDB connection url", "MONGO_URL")
	flag.Parse()
}

func main() {
	components.MongoInit(mongoUrl)
	defer components.MongoClose()

	router := gin.Default()

	app := controllers.NewApp()

	router.GET("/", app.Status)

	contacts := controllers.NewContacts()

	router.GET("/contacts", contacts.Index)
	router.POST("/contacts", contacts.Create)
	router.GET("/contacts/:id", contacts.Show)
	//router.PATCH("/contacts/:id", contacts.PartialUpdate)
	router.PUT("/contacts/:id", contacts.Update)
	router.DELETE("/contacts/:id", contacts.Destroy)

	router.Run(addr)
}
