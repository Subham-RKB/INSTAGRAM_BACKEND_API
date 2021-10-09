package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id   bson.ObjectId `json:"id" bosn:"_id"`
	Name string        `json:"name" bson:"name"`
}
