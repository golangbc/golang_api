package auth

import (
	"database/sql"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 管理者権限、バッチのみ使える処理をここにまとめます。
func Login(c *gin.Context, db *sql.DB) {

	session := sessions.Default(c)

	if IsLogin() {

	} else if LoginAuth() {
		session.Set("hello", "world")
		session.Save()
	}
	// if session.Get("hello") != "world" {
	// 	session.Set("hello", "world")
	// 	session.Save()
	// }

	c.JSON(200, gin.H{"hello": session.Get("hello")})
}

func IsLogin() bool {
	// loginしてたら、loginCheckMiddleWareを呼ぶ。
	return true
}

func LoginAuth() bool {
	// loginしてたら、
	return true
}
