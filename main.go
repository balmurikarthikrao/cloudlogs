package main

import (
	"cloud/controllers"
	_ "cloud/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:admin@123@tcp(127.0.0.1:3306)/CloudWatchLog")
	name := "default"
	force := true
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

}
func main() {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	maincontroller := new(controllers.MainController)
	router.GET("/logstream", maincontroller.LogStreams)
	router.GET("/logevent", maincontroller.LogEvents)
	router.Run(":8081")

}
