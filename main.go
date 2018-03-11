package main

import (
	"flag"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/memclutter/gontacts/components"
	"github.com/memclutter/gontacts/controllers"
	"github.com/memclutter/gontacts/middlewares"
	"github.com/memclutter/gontacts/utils"
	"github.com/pkg/profile"
)

var (
	addr       string
	mongoUrl   string
	mailerHost string
	mailerPort string
	mailerUser string
	mailerPass string
)

func init() {
	utils.FlagStringVarEnv(&addr, "addr", ":8000", "Set address", "ADDR")
	utils.FlagStringVarEnv(&mongoUrl, "mongoUrl", "mongodb://localhost:27017/gontacts", "MongoDB connection url", "MONGO_URL")
	utils.FlagStringVarEnv(&mailerHost, "mailerHost", "smtp.mailtrap.io", "Mailer SMTP host", "MAILER_HOST")
	utils.FlagStringVarEnv(&mailerPort, "mailerPort", "25", "Mailer SMTP port", "MAILER_PORT")
	utils.FlagStringVarEnv(&mailerUser, "mailerUser", "a8856fa3e2904b", "Mailer SMTP username", "MAILER_USER")
	utils.FlagStringVarEnv(&mailerPass, "mailerPass", "e2f8970c426f95", "Mailer SMTP user password", "MAILER_PASS")

	flag.Parse()
}

func main() {
	if gin.IsDebugging() {
		defer profile.Start().Stop()
	}

	components.MongoInit(mongoUrl)
	defer components.MongoClose()

	mailerPortInt, _ := strconv.Atoi(mailerPort)
	components.MailerInit(mailerHost, mailerPortInt, mailerUser, mailerPass)
	defer close(components.MailerCh)

	router := gin.Default()
	authorization := middlewares.Authorization()

	app := controllers.NewApp()

	router.GET("/", app.Status)

	users := controllers.NewUsers()

	router.GET("/users/info", authorization, users.Info)
	router.POST("/users/registration", users.Registration)
	router.POST("/users/login", users.Login)
	router.POST("/users/confirmation", users.Confirmation)

	contacts := controllers.NewContacts()

	router.GET("/contacts", authorization, contacts.Index)
	router.POST("/contacts", authorization, contacts.Create)
	router.GET("/contacts/:id", authorization, contacts.Show)
	//router.PATCH("/contacts/:id", contacts.PartialUpdate)
	router.PUT("/contacts/:id", authorization, contacts.Update)
	router.DELETE("/contacts/:id", authorization, contacts.Destroy)

	router.Run(addr)
}
