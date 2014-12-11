package models

import (
	"gopkg.in/mgo.v2/bson"
)

type RunCollection struct {
	Data []Run `json:"data"`
}

type Run struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
