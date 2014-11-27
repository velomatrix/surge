package tests

import "github.com/revel/revel"

type CheckTest struct {
	revel.TestSuite
}

func (t *CheckTest) TestHealth() {
	t.Get("/healthcheck")
	t.AssertOk()
}
