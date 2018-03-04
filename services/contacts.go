package services

import (
	"github.com/memclutter/gontacts/components"
	"github.com/memclutter/gontacts/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Contacts struct {
	collection *mgo.Collection
	storage    []models.Contact
}

func NewContacts() *Contacts {
	return &Contacts{
		collection: components.MongoDB.C("contacts"),
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

func (s *Contacts) All(skip, limit int, order string) ([]models.Contact, int, error) {
	var err error

	all := make([]models.Contact, 0)
	totalCount := 0
	where := bson.M{}
	query := s.collection.Find(where)

	if totalCount, err = query.Count(); err != nil {
		return all, totalCount, err
	}

	err = query.Skip(skip).Limit(limit).All(&all)

	return all, totalCount, err
}

func (s *Contacts) Create(model *models.Contact) error {
	model.Id = bson.NewObjectId()
	return s.collection.Insert(model)
}

func (s *Contacts) Get(id bson.ObjectId) (one *models.Contact, err error) {
	err = s.collection.FindId(id).One(&one)
	return
}

//func (s *Contacts) PartialUpdate(id bson.ObjectId, newModel *models.Contact) error {
//	if index, oldModel := s.find(id); index < 0 {
//		return errors.New("Model not found: ")
//	} else {
//		oldRef := reflect.ValueOf(oldModel).Elem()
//		newRef := reflect.ValueOf(newModel).Elem()
//
//		for i := 0; i < oldRef.NumField(); i++ {
//			oldField := oldRef.Field(i)
//			newField := newRef.Field(i)
//
//			if oldField.Type().String() == "string" {
//				oldStringValue := oldField.Interface().(string)
//				newStringValue := newField.Interface().(string)
//
//				if strings.Compare(oldStringValue, newStringValue) != 0 {
//					if oldField.CanSet() {
//						oldField.SetString(newStringValue)
//					}
//				}
//			}
//		}
//
//		newModel = oldModel
//
//		return s.collection.UpdateId(id, bson.M{
//		})
//	}
//}

func (s *Contacts) Update(id bson.ObjectId, model *models.Contact) error {
	model.Id = id
	return s.collection.UpdateId(id, model)
}

func (s *Contacts) Destroy(id bson.ObjectId) error {
	return s.collection.RemoveId(id)
}
