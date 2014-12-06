package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Run struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func FindRun(s *mgo.Session, id bson.ObjectId) *Run {
	var run Run
	s.DB("").C("runs").FindId(id).One(&run)
	return &run
}

func FindAllRuns(s *mgo.Session) *[]Run {
	var results []Run

	err := s.DB("").C("runs").Find(nil).All(&results)
	if err != nil {
		fmt.Printf("Unable to retrieve all Run documents")
		return nil
	}
	return &results
}
