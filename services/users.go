package services

import (
	"fmt"

	"github.com/memclutter/gontacts/components"
	"github.com/memclutter/gontacts/models"
	"github.com/memclutter/gontacts/utils"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	collection *mgo.Collection
}

func NewUsers() *Users {
	return &Users{
		collection: components.MongoDB.C("users"),
	}
}

func (s *Users) Registration(model *models.User) error {
	model.Id = bson.NewObjectId()

	// Hashing password
	if ph, err := utils.HashPassword(model.PasswordHash); err != nil {
		return err
	} else {
		model.PasswordHash = ph
	}

	// Generate confirmation token
	model.ConfirmationToken = utils.GenerateRandomToken(16)

	if err := s.collection.Insert(model); err != nil {
		return err
	}

	// Send email confirmation
	message := gomail.NewMessage()
	message.SetHeader("From", "noreply@gontacts.example.com")
	message.SetHeader("To", model.Email)
	message.SetHeader("Subject", "User confirmation")

	// Html template
	html := fmt.Sprintf("<h1>User confirmation</h1><p><a href=\"http://localhost:8000/confirmation?token=%v\">Confirmation link</a></p>", model.ConfirmationToken)
	message.SetBody("text/html", html)

	// Text template
	text := fmt.Sprintf("User confirmation\n\nConfirmation link: http://localhost:8000/confirmation?token=%v", model.ConfirmationToken)
	message.AddAlternative("text/plain", text)

	// Send email via mailer daemon
	components.MailerCh <- message

	// return success
	return nil
}

func (s *Users) Confirmation(token string) error {
	model := new(models.User)

	if err := s.collection.Find(bson.M{"confirmation_token": token}).One(&model); err != nil {
		return err
	}

	updateQuery := bson.M{
		"$set": bson.M{
			"confirmation_token": "",
			"is_confirmed":       true,
		},
	}

	return s.collection.UpdateId(model.Id, updateQuery)
}

func (s *Users) Login(model *models.Login) (string, error) {
	var userModel models.User

	findQuery := bson.M{
		"email":        model.Email,
		"is_confirmed": true,
	}

	if err := s.collection.Find(findQuery).One(&userModel); err != nil {
		return "", err
	} else if !utils.CheckPasswordHash(model.Password, userModel.PasswordHash) {
		return "", errors.New("Password mismatch")
	} else {
		return utils.CreateJwtToken(userModel.Id, userModel.Email), nil
	}
}
