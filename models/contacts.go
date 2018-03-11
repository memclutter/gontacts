package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Contact struct {
	Id       bson.ObjectId `bson:"_id" json:"id"`
	UserId   bson.ObjectId `bson:"user_id" json:"-"`
	FullName string        `bson:"full_name" json:"full_name" form:"full_name" binding:"required"`
	Phone    string        `bson:"phone" json:"phone" form:"phone" binding:"required"`
	Email    string        `bson:"email,omitempty" json:"email,omitempty" form:"email"`
	Address  string        `bson:"address,omitempty" json:"address,omitempty" form:"address"`
	Skype    string        `bson:"skype,omitempty" json:"skype,omitempty" form:"skype"`
	Telegram string        `bson:"telegram,omitempty" json:"telegram,omitempty" form:"telegram"`
}
