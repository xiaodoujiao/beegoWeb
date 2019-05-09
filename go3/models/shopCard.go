package models

import "github.com/astaxie/beego/orm"
import _ "github.com/go-sql-driver/mysql"

type ShopCard struct {
	Id int
	Name string
	Price int
}

func init1(){
	//orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/yzdb")
	orm.RegisterDataBase("test","mysql","root:123456@tcp(127.0.0.1:3306)/test")
	orm.RegisterModel(new(ShopCard))
	//orm.RunSyncdb("default",false,true)
	orm.RunSyncdb("test",false,true)

}
