package routes

import (
	"database/sql"

	"github.com/KatsutoshiOtogawa/golang_api/handler/users"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 一般ユーザー権限のみ使える処理をここにまとめます。
func User(r *gin.Engine, db *sql.DB) {
	// 管理者権限のみ使える処理。
	user := r.Group("user")

	store := cookie.NewStore([]byte("secret"))
	user.Use(sessions.Sessions("mysession", store))

	user.GET("/users/hello", func(c *gin.Context) {users.Hello(c,db) })
}