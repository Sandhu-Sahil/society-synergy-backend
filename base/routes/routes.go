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
			userSimpleGroup.POST("/login", rs.Controller.Login)
			userSimpleGroup.POST("/register", rs.Controller.Register)
		}
	}

	jwtGroup := rg.Group("v2") //with jwt headers
	{
		jwtGroup.Use(middlewares.JwtAuthMiddleware())

		userGroup := jwtGroup.Group("/user")
		{
			userGroup.GET("/:id", rs.Controller.GetUser)
		}

		// teamGroup := jwtGroup.Group("/team")
		// {
		// 	teamGroup.GET("/leaderboard", rs.Controller.TeamLeaderboard)
		// }
	}
}
