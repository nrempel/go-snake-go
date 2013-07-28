package controllers

import "github.com/robfig/revel"
import "encoding/json"

type App struct {
	*revel.Controller
}

func (c App) Register() revel.Result {
	b, err := ioutil.ReadAll(c.Request.Body)
	revel.INFO.Println(string(b), err)
}
