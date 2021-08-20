package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserNo    int64              `form:"userno" json:"userno"`
	FirstName string             `form:"firstName" json:"firstName"`
	LastName  string             `form:"lastName"  json:"lastName"`
	Username  string             `form:"username" json:"username" validate:"required"`
	Password  string             `form:"password" json:"password" validate:"required"`
}
