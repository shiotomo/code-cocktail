package main

import (
	"./config"
	"./controllers"
	"./models"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// 初期準備
	flag.Parse()
	e := echo.New()
	var db *gorm.DB

	// 実行時引数がproductionであればproduction環境で実行する
	// それ以外であればdevelopment環境で実行する
	if flag.Arg(0) == "production" {
		db = config.PQPro()
		fmt.Println("=== production ===")
	} else {
		db = config.PQDev()
		fmt.Println("=== development ===")
	}

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Migrateをする
	models.Migrate(db)

	// == ルーティング ==
	// サンプル
	e.GET("/", controllers.Sample())

	// index
	e.GET("/user", controllers.UserIndex())
	e.GET("/question", controllers.QueIndex())
	e.GET("/contest", controllers.ConIndex())
	e.GET("/contest/team", controllers.TeamIndex())

	// show
	e.GET("/user/:id", controllers.UserShow())
	e.GET("/question/:id", controllers.QueShow())
	e.GET("/contest/:id", controllers.ConShow())
	e.GET("/contest/team/:id", controllers.TeamShow())

	e.Logger.Fatal(e.Start(":3000"))
}
