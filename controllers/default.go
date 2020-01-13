package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"imooc/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*
	m:=models.GetPage()
	c.Data["Id"]=m.Id
	c.Data["Website"] = m.Website
	c.Data["Email"] = m.Email
	*/
	c.TplName = "login.tpl"
}



func (c *MainController)RegPost(){
	Wb,_:=c.GetInt("Id")
	o:=orm.NewOrm()
	p:=models.Page{Id:Wb}
	err:=o.Read(&p)
	c.Data["Id"]=p.Id
	c.Data["Website"]=p.Website
	c.Data["Email"]=p.Email
	c.TplName="log.html"
	if p.Website== ""{
		c.TplName="login.tpl"
	}
	fmt.Println(err)
	return
}
