package services

import (
	"fmt"
	"github.com/velomatrix/surge/api/app/lib"
	. "github.com/velomatrix/surge/api/app/models"
	"gopkg.in/mgo.v2/bson"
)

func FindRun(id string) (*Run, error) {
	var run Run

	objectId := bson.ObjectIdHex(id)
	err := mongo.Config.Session.DB("").C("runs").FindId(objectId).One(&run)
	if err != nil {
		fmt.Printf("Unable to retrieve run")
		return nil, err
	}
	return &run, nil
}

func FindAllRuns() (*[]Run, error) {
	var results []Run

	err := mongo.Config.Session.DB("").C("runs").Find(nil).All(&results)
	if err != nil {
		fmt.Printf("Unable to retrieve all Run documents")
		return nil, err
	}
	return &results, nil
}
