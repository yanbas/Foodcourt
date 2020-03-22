package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/foodCourt/app"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	fmt.Println("Running ...")

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	c := app.DBConfig{}
	app := app.App{}
	c.GetConfig()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.Database.Host,
		c.Database.Port,
		c.Database.Username,
		c.Database.DBName,
		c.Database.Password)

	db, err := gorm.Open(c.Database.Driver, stringConnection)
	if err != nil {
		log.Panic("Error Connction Database, ", err.Error())
	}
	defer db.Close()
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(5)
	db.LogMode(true)

	app.DB = *&db
	app.Log = *&f

	app.Running()

}
