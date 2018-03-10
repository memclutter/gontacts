package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id                bson.ObjectId `bson:"_id" json:"id"`
	Email             string        `bson:"email" json:"email" form:"email" binding:"required"`
	PasswordHash      string        `bson:"password_hash" json:"-" form:"password" binding:"required"`
	IsConfirmed       bool          `bson:"is_confirmed" json:"-"`
	ConfirmationToken string        `bson:"confirmation_token" json:"-"`
}
