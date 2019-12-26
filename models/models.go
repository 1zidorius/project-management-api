package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	Id        string     `json:"id,omitempty"`
	Subject   string     `json:"subject,omitempty"`
	Status    string     `json:"status,omitempty"`
	Priority  int        `json:"priority,omitempty"`
	CreatedOn time.Time  `json:"createdOn,omitempty" bson:"createdOn"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty" bson:"updatedOn"`
	AddedBy   string     `json:"addedBy,omitempty"`
	Assignee  string     `json:"assignee"`
}

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Username string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty"`
	Email    string             `json:"email,omitempty"`
	Name     string             `json:"name,omitempty"`
	Surname  string             `json:"surname,omitempty"`
}
