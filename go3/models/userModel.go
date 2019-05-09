package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int
	Name string
	Pwd string
}
type Child struct {
	Id int
	Num int
}

type Father struct {
	Id int
	Name string
}
type Article struct {
	Id int `orm:"pk;auto"`
	Title string `orm:"unique;size(40)"`
	Content string `orm:"size(500)"`
	Img string	`orm:"null"`
	Time time.Time `orm:"type(datetime);auto_now_add"`
	ReadCount int `orm:"default(0)"`
}
func init(){
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test430")
	orm.RegisterModel(new(User),new(Article))
	orm.RunSyncdb("default",false,true)
}
