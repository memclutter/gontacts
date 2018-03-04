package services

import (
	"strings"

	"reflect"

	"github.com/memclutter/gontacts/models"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

type Contacts struct {
	storage []models.Contact
}

func NewContacts() *Contacts {
	return &Contacts{
		storage: []models.Contact{
			{
				Id:       bson.NewObjectId(),
				FullName: "John Doe",
				Email:    "john.doe@example.com",
				Phone:    "+112323232",
				Address:  "23, Perfect street, NY",
				Skype:    "john.doe",
				Telegram: "@john.doe",
			},
		},
	}
}

func (s *Contacts) All() ([]models.Contact, int, error) {
	return s.storage, len(s.storage), nil
}

func (s *Contacts) Create(model *models.Contact) error {
	model.Id = bson.NewObjectId()
	s.storage = append(s.storage, *model)
	return nil
}

func (s *Contacts) Get(id bson.ObjectId) (*models.Contact, error) {
	if index, model := s.find(id); index >= 0 {
		return model, nil
	} else {
		return nil, errors.New("Model not found: ")
	}
}

func (s *Contacts) PartialUpdate(id bson.ObjectId, newModel *models.Contact) error {
	if index, oldModel := s.find(id); index < 0 {
		return errors.New("Model not found: ")
	} else {
		oldRef := reflect.ValueOf(oldModel).Elem()
		newRef := reflect.ValueOf(newModel).Elem()

		for i := 0; i < oldRef.NumField(); i++ {
			oldField := oldRef.Field(i)
			newField := newRef.Field(i)

			if oldField.Type().String() == "string" {
				oldStringValue := oldField.Interface().(string)
				newStringValue := newField.Interface().(string)

				if strings.Compare(oldStringValue, newStringValue) != 0 {
					if oldField.CanSet() {
						oldField.SetString(newStringValue)
					}
				}
			}
		}

		newModel = oldModel
		s.storage[index] = *newModel

		return nil
	}
}

func (s *Contacts) Update(id bson.ObjectId, model *models.Contact) error {
	if index, _ := s.find(id); index < 0 {
		return errors.New("Model not found: ")
	} else {
		model.Id = id
		s.storage[index] = *model

		return nil
	}
}

func (s *Contacts) Destroy(id bson.ObjectId) error {
	if index, _ := s.find(id); index < 0 {
		return errors.New("Model not found: ")
	} else {
		s.storage = append(s.storage[:index], s.storage[index+1:]...)

		return nil
	}
}

func (s *Contacts) find(id bson.ObjectId) (int, *models.Contact) {
	for index, model := range s.storage {
		if strings.Compare(model.Id.Hex(), id.Hex()) == 0 {
			return index, &model
		}
	}

	return -1, nil
}
