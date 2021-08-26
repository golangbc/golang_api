package auth

import (
	"database/sql"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// middleware
func CheckLogin(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		session := sessions.Default(c)

		if IsLogin() {

		}
		// if session.Get("hello") != "world" {
		// 	session.Set("hello", "world")
		// 	session.Save()
		// }

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	}
}

func IsLogin() bool {
	// loginしてたら、loginCheckMiddleWareを呼ぶ。
	return true
}

func LoginAuth(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		session := sessions.Default(c)

		session.Set("hello", "world")
		session.Save()

	}
}
