package tests

import "github.com/revel/revel"

type RunTest struct {
	revel.TestSuite
}

func (t *RunTest) Before() {
	println("Set up run test")
}

func (t *RunTest) testBasicURL() {
	t.Get("/runs")
	t.AssertOk()
	// TODO : extract json
}
