package routes

import (
	"database/sql"

	"github.com/KatsutoshiOtogawa/golang_api/handler/auth"
	"github.com/KatsutoshiOtogawa/golang_api/handler/users"
	"github.com/gin-gonic/gin"
)

// 管理者権限、バッチのみ使える処理をここにまとめます。
func Admin(r *gin.Engine, db *sql.DB) {

	// 管理者権限でログインできた人のみ
	admin := r.Group("admin", auth.CheckLogin(auth.AdminLoginFlg, db))

	admin.GET("/hello", users.Hello(db))

}
