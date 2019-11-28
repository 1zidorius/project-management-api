package models

import (
	"time"
)

type Task struct {
	Id        string    `json:"id,omitempty"`
	Subject   string    `json:"subject,omitempty"`
	Status    string    `json:"status,omitempty"`
	Priority  string    `json:"priority,omitempty"`
	CreatedOn time.Time `json:"createdOn,omitempty" bson:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn,omitempty" bson:"updatedOn"`
	AddedBy   User      `json:"addedBy,omitempty"`
	Assignee   User      `json:"assignee"`
}

type User struct {
	Id          string `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
}
