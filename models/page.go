package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Id int
	Website string
	Email string
}

func init(){
	orm.RegisterDataBase("default","mysql","root:199700@tcp(127.0.0.1:3306)/test?charset=utf8&local")
	orm.RegisterModel(new(Page))
}

/*
func GetPage()Page{
	o:=orm.NewOrm()
	p:=Page{Id:1}
	err:=o.Read(&p)
	fmt.Println(err)
	return p
}
 */
