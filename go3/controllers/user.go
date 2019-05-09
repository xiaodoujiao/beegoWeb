package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go3/models"
)

type UserController struct {
	beego.Controller
}
func(c *UserController) Register(){
	c.TplName = "register.html"

}
func (c *UserController)HandleReg(){
	name := c.GetString("userName")
	pwd := c.GetString("password")
	if name=="" || pwd == "" {
		beego.Error("data-not-full")
		c.TplName = "register.html"
		return
	}
	var user models.User
	user.Name = name
	user.Pwd = pwd

	o := orm.NewOrm()
	_,err:=o.Insert(&user)
	if err!=nil{
		beego.Error("insert err")
		c.TplName = "register.html"
		return
	}
	c.TplName = "login.html"
	c.Redirect("/yz/login.html",302)
}
//LoginShow
func (c *UserController)LoginShow(){
	c.TplName = "login.html"
}

//LoginHandle
func(c *UserController)LoginHandle(){
	name := c.GetString("userName")
	pwd := c.GetString("password")
	var user models.User
	user.Pwd = pwd
	user.Name = name

	o := orm.NewOrm()
	err := o.Read(&user,"Name")
	if err!=nil{
		beego.Error("not exists username")
		c.TplName = "login.html"
		return
	}
	if user.Pwd != pwd {
		beego.Error("pwd error")
		c.TplName = "login.html"
		return
	}

	//c.Ctx.WriteString("login success")
	c.Redirect("/index",302)
}
func (c *UserController)ShowIndex(){
	c.TplName = "index.html"
}
