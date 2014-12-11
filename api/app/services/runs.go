package services

import (
	"github.com/velomatrix/surge/api/app/lib"
	. "github.com/velomatrix/surge/api/app/models"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func FindRun(id string) (*Run, error) {
	var run Run

	objectId := bson.ObjectIdHex(id)
	err := mongo.Config.Session.DB("").C("runs").FindId(objectId).One(&run)
	if err != nil {
		log.Printf("Unable to retrieve run")
		return &run, err
	}
	return &run, nil
}

func FindAllRuns() (*[]Run, error) {
	var results RunCollection

	err := mongo.Config.Session.DB("").C("runs").Find(nil).All(&results.Data)
	if err != nil {
		log.Printf("Unable to retrieve all Run documents")
		return &results.Data, err
	}
	return &results.Data, nil
}

func CreateRun(run *Run) error {
	id := bson.NewObjectId()
	_, err := mongo.Config.Session.DB("").C("runs").UpsertId(id, run)
	if err != nil {
		return err
	}

	run.Id = id

	return nil
}
