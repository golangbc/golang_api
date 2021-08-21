package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// 各ルートの割当をここに書きます。
func Invoke(r *gin.Engine, db *sql.DB) {
	// 管理者権限のみ入るルート
	Admin(r,db)
	// ユーザーがログインしたら入れるルート
	User(r,db)
	// ログインしなくても入れるルート
	NoLogin(r,db)
}
