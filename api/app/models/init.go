package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/revel/revel"
)

var (
	mgoSession *mgo.Session
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		connectString, _ := revel.Config.String("mongo.uri")
		if connectString == "" {
			connectString = "mongodb://localhost/surge_development"
		}
		fmt.Printf("connection string: %s\n", connectString)
		mgoSession, err = mgo.Dial(connectString)
		if err != nil {
			fmt.Printf("Can't connect to mongo, go error %v\n", err)
		}
	}
	return mgoSession.Clone()
}
