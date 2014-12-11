package mongo

import (
	"gopkg.in/mgo.v2"
)

var Config struct {
	Session *mgo.Session
}

func InitDB() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	Config.Session = session
	Config.Session.SetMode(mgo.Monotonic, true)

}
