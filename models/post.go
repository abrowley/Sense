package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type(
	Post struct {
		Id	bson.ObjectId `json:"id" bson:"_id"`
		Sender 	string `json:"sender" bson:"sender"`
		Message string `json:"message" bson:"message"`
		TimeReceived time.Time `json:"time_recv" bson:"time_recv"`
	}
)