package main

import (
	_ "GoGoCar/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.DirectoryIndex = true
	beego.SetStaticPath("app", "static/dist")
	beego.SetStaticPath("bower_components", "static/bower_components")
	beego.SetStaticPath("custom_components", "static/custom_components")
	beego.Run()
}
