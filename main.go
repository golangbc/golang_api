package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/KatsutoshiOtogawa/golang_api/routes"
)

func main() {
	r := gin.Default()

	db, _ := sql.Open("sqlite3", "sqlite3.db")

	// 管理者権限のルート
	routes.Admin(r, db)

	r.Run() // listen and serve on 0.0.0.0:8080
}
