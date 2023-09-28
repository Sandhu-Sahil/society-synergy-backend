package routes

import (
	"Society-Synergy/base/middlewares"

	"github.com/gin-gonic/gin"
)

func (rs *RouterService) RegisterRoutes(rg *gin.RouterGroup) {

	simpleGroup := rg.Group("v1") //without jwt
	{
		userSimpleGroup := simpleGroup.Group("/user")
		{
			userSimpleGroup.POST("/login", rs.UserController.Login)
			userSimpleGroup.POST("/register", rs.UserController.Register, rs.LogsController.RegisterLog)
		}
		testGroup := simpleGroup.Group("/test")
		{
			testGroup.GET("/ping", rs.UserController.Ping)
		}
	}

	jwtGroup := rg.Group("v2") //with jwt headers
	{
		jwtGroup.Use(middlewares.JwtAuthMiddleware())

		userGroup := jwtGroup.Group("/user")
		{
			userGroup.GET("/:id", rs.UserController.GetUser)
		}

		// teamGroup := jwtGroup.Group("/team")
		// {
		// 	teamGroup.GET("/leaderboard", rs.Controller.TeamLeaderboard)
		// }
	}
}
