package routes

import (
	"database/sql"

	"github.com/KatsutoshiOtogawa/golang_api/handler/users"
	"github.com/gin-gonic/gin"
)

// 管理者権限、バッチのみ使える処理をここにまとめます。
func Admin(r *gin.Engine, db *sql.DB) {
	// 管理者権限のみ使える処理。
	admin := r.Group("admin")
	admin.GET("/hello", func(c *gin.Context) { users.Hello(c, db) })
}
