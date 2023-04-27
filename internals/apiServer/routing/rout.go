package routing

import (
	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	g.GET("signup", SignUp)
	g.GET("signin", SignIn)
	signInGroup := g.Group("signin")
	signInGroup.POST("search", authMiddleware(false), Search)
	signInGroup.POST("newdir", authMiddleware(false), AddDirPath)
	signInGroup.POST("change-engine", authMiddleware(true), ChangeEngine)
}
