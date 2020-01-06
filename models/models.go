package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Subject   string             `json:"subject,omitempty"`
	Status    string             `json:"status,omitempty"`
	Priority  int                `json:"priority,omitempty"`
	CreatedOn time.Time          `json:"createdOn,omitempty" bson:"createdOn"`
	UpdatedOn time.Time         `json:"updatedOn,omitempty" bson:"updatedOn,omitempty"`
	AddedBy   primitive.ObjectID `json:"addedBy,omitempty"`
	Assignee  primitive.ObjectID `json:"assignee,omitempty" bson:"assignee,omitempty"`
}

type ResultUser struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name    string             `json:"name,omitempty"`
	Surname string             `json:"surname,omitempty"`
}

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Username string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty"`
	Email    string             `json:"email,omitempty"`
	Name     string             `json:"name,omitempty"`
	Surname  string             `json:"surname,omitempty"`
}
