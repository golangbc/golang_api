package routes

import (
	"database/sql"

	"github.com/KatsutoshiOtogawa/golang_api/handler/users"
	"github.com/gin-gonic/gin"
)

// ログインしなくても入れるルートです。
func NoLogin(r *gin.Engine, db *sql.DB) {
	// 管理者権限のみ使える処理。
	nologin := r.Group("nologin")
	nologin.GET("/hello", users.Hello(db))
}
