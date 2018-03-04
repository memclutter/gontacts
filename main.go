package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/memclutter/gontacts/controllers"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8000", "Set address")
	flag.Parse()
}

func main() {
	router := gin.Default()

	app := controllers.NewApp()

	router.GET("/", app.Status)

	contacts := controllers.NewContacts()

	router.GET("/contacts", contacts.Index)
	router.POST("/contacts", contacts.Create)
	router.GET("/contacts/:id", contacts.Show)
	router.PATCH("/contacts/:id", contacts.PartialUpdate)
	router.PUT("/contacts/:id", contacts.Update)
	router.DELETE("/contacts/:id", contacts.Destroy)

	router.Run(addr)
}
