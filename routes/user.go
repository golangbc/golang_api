package routes

import (
	"database/sql"

	"github.com/KatsutoshiOtogawa/golang_api/handler/auth"
	"github.com/KatsutoshiOtogawa/golang_api/handler/users"
	"github.com/gin-gonic/gin"
)

// 一般ユーザー権限のみ使える処理をここにまとめます。
func User(r *gin.Engine, db *sql.DB) {

	// ユーザー権限でログインできた人のみ
	user := r.Group("user", auth.CheckLogin(auth.UserLoginFlg, db))

	// user.Use(middleware.LoginCheckMiddleware())
	// {
	// 	user.GET("/users/hello", func(c *gin.Context) { users.Hello(c, db) })
	// }
	user.GET("/hello", users.Hello(db))

}
