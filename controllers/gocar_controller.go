package controllers

import (
	"GoGoCar/models"
	"fmt"
	"github.com/astaxie/beego"
)

type DBController struct {
	beego.Controller
}

var b models.DB

func init() {
	dbUser := beego.AppConfig.String("mysqluser")
	dbPasswd := beego.AppConfig.String("mysqlpass")
	dbIP := beego.AppConfig.String("mysqlurls")
	dbName := beego.AppConfig.String("mysqldb")

	b = models.DB{dbUser, dbPasswd, dbIP, dbName}
	b.Init()
}

func (this *DBController) Prepare() {
	//fmt.Println("Validation")

}

func (this *DBController) GetClients() {
	queryString := "SELECT * FROM client"
	num, result := b.Query(queryString)
	fmt.Println(num)
	this.Data["json"] = result
	this.ServeJson()
}

func (this *DBController) GetDetails() {
	queryString := "SELECT * FROM maintain_detail"
	_, result := b.Query(queryString)

	this.Data["json"] = result
	this.ServeJson()
}
