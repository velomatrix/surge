package tests

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/velomatrix/surge/api/app/models"
	"gopkg.in/mgo.v2"
)

type RunTest struct {
	revel.TestSuite
	*mgo.Session
}

// Simple type of array of run for the API -- will need to convert once/if we support pagination
type RunsResult []models.Run

func (t *RunTest) Before() {

	// next step should be to use the configuration information
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	t.Session = session
	t.Session.SetMode(mgo.Monotonic, true)
	c := t.Session.DB("test").C("runs")
	err = c.Insert(&models.Run{"000000000001", "Run A"})
	if err != nil {
		// lets be forgiving because right now it will fail if it's already there
		// should be more robust later
	}
}

func (t *RunTest) TestGetRuns() {

	t.Get("/runs")
	t.AssertContentType("application/json; charset=utf-8")
	t.AssertOk()

	r := RunsResult{}
	err := json.Unmarshal(t.ResponseBody, &r)

	t.Assert(err == nil)
	t.Assert(r[0].Name == "Run A")
	t.Assert(len(r) == 1)

}

func (t *RunTest) TestBasicURL() {
	t.Get("/runs")
	t.AssertOk()
}
