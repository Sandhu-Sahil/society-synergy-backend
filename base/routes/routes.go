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
			userSimpleGroup.POST("/register", rs.UserController.Register)
			userSimpleGroup.POST("/otpsendemail", rs.UserController.EmailSendOtp)
			userSimpleGroup.POST("/otpverifyemail", rs.UserController.VerifyOtpEmail)
		}
		testGroup := simpleGroup.Group("/test")
		{
			testGroup.GET("/ping", rs.UserController.Ping)
		}
		departmentGroup := simpleGroup.Group("/department")
		{
			departmentGroup.GET("/:id", rs.UserController.DepartmentLeaderboard)
		}
		eventGroup := simpleGroup.Group("/event")
		{
			eventGroup.GET("/:id", rs.UserController.EventLeaderboard)
		}
		homeGroup := simpleGroup.Group("/home")
		{
			homeGroup.GET("/leaderboard", rs.UserController.HomeLeaderboard)
		}
	}

	jwtGroup := rg.Group("v2") //with jwt headers
	{
		jwtGroup.Use(middlewares.JwtAuthMiddleware())

		userGroup := jwtGroup.Group("/user")
		{
			userGroup.GET("/:id", rs.UserController.GetUser)
			userGroup.POST("/changePassword", rs.UserController.ChangePassword)
			userGroup.POST("/updateUser", rs.UserController.UpdateUser)
		}
		emailGroup := jwtGroup.Group("/email")
		{
			emailGroup.POST("/otpsend", rs.UserController.SendOtp)
			emailGroup.POST("/otpverify", rs.UserController.VerifyOtp)
		}
		departmentGroup := jwtGroup.Group("/department")
		{
			// departmentGroup.GET("/:id", rs.UserController.DepartmentLeaderboard)
			departmentGroup.POST("/create", rs.UserController.CreateDepartment)
		}
		memberGroup := jwtGroup.Group("/member")
		{
			memberGroup.POST("/create", rs.UserController.CreateMember)
		}
		eventGroup := jwtGroup.Group("/event")
		{
			eventGroup.POST("/create", rs.UserController.CreateEvent)
			emailGroup.POST("/addrsvp", rs.UserController.AddRsvp)
		}
		sandhuGroup := jwtGroup.Group("/sandhu")
		{
			sandhuGroup.POST("/createAdmin", rs.UserController.SandhuCreateAdmin)
		}

		// teamGroup := jwtGroup.Group("/team")
		// {
		// 	teamGroup.GET("/leaderboard", rs.Controller.TeamLeaderboard)
		// }
	}
}
