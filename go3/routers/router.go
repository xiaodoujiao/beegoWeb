package routers

import (
	"go3/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
   /* beego.Router("/yz/login",&controllers.YzController{},"post:Login;get:ShowLogin")
    beego.Router("/yz/register0",&controllers.YzController{},"get:Reg;post:HandleReg")
    beego.Router("/yz/query",&controllers.YzController{},"*:Query")
	beego.Router("/yz/update",&controllers.YzController{},"*:Update")
    beego.Router("/yz/hello",&controllers.YzController{},"get:Hello")
*/
	beego.Router("/yz/login",&controllers.UserController{},"get:LoginShow;post:LoginHandle")

	beego.Router("/yz/register",&controllers.UserController{},"get:Register;post:HandleReg")
	beego.Router("/index",&controllers.ArticleController{},"*:ShowIndex")
	//	addArticle
	beego.Router("/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:AddArticle")

}
