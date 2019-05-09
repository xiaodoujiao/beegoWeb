package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go3/models"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}
//ShowIndex
func (this *ArticleController)ShowIndex(){
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var articles []*models.Article
	//articles = make([]models.Article,0)
	qs.All(&articles)
	this.Data["articles"] = articles
	//total count
	totalCnt,err := qs.Count()
	if err!=nil{
		beego.Info(err)
		return
	}

	//page Num
	pNum := totalCnt/2
	//now page--default:1
	nowPage := this.GetString("nowPage")
	if nowPage=="" {
		nowPage=""
	}
	//this page data
	this.TplName = "index.html"
}

func(c *ArticleController) ShowAddArticle(){
	c.TplName = "add.html"
}
func(c *ArticleController)AddArticle(){

	//articleName
	name := c.GetString("articleName")
	//content
	content := c.GetString("content")
	if name == "" || content == ""{
		beego.Info("input-err-:not empty")
		return
	}
	//uploadname
	file,head,err := c.GetFile("uploadname")
	if err!=nil{
		beego.Info("get-file-err:",err)
		return
	}
	defer file.Close()
	//defer file.Close()
	if head.Size >1000000 {
		beego.Info("filesize-too-big")
		return
	}
	ext := path.Ext(head.Filename)
	if ext!=".jpg" && ext!=".png"{
		beego.Error("获取数据错误")
		return

	}
	filename := time.Now().Format("200601021504052222")

	c.SaveToFile("uploadname","./static/img/"+filename)
	o := orm.NewOrm()

	var article models.Article
	article.Content = content
	article.Title = name
	article.Img = "/static/img/"+filename

	_,err = o.Insert(&article)
	if err!=nil{
		beego.Info("insert-err:",err)
		return
	}



	c.Redirect("/index",302)
}


