package controllers

import (
	"github.com/revel/revel"
	"github.com/velomatrix/revmgo"
	"github.com/velomatrix/surge/api/app/models"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	*revel.Controller
	revmgo.MongoController
}

func (c User) Index() revel.Result {
	users, _ := models.FindAllUsers(c.MongoSession)
	return c.RenderJson(users)
}

func (c User) Show(id string) revel.Result {
	user := models.FindUser(c.MongoSession, bson.ObjectIdHex(id))
	return c.RenderJson(user)
}

func (c User) Create() revel.Result {
	return nil
}
