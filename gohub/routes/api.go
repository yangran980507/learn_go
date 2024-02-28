// Package routes 注册路由
package routes

import (
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")

	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")

		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			// 注册用户
			authGroup.POST("/signup/phone/exist",
				middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist",
				middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(),
				suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(),
				suc.SignupUsingEmail)

			//发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/phone",
				middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email",
				middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)
			// 图片验证码
			authGroup.POST("/verify-codes/captcha",
				middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)

			// 登陆用户
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(),
				lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(),
				lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(),
				lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(),
				pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(),
				pwc.ResetByEmail)

			uc := new(controllers.UsersController)
			// 获取当前用户
			v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
			usersGroup := v1.Group("/users")
			{
				usersGroup.GET("", uc.Index)
			}

			cgc := new(controllers.CategoriesController)
			cgcGroup := v1.Group("/categories")
			{
				cgcGroup.GET("", cgc.Index)
				cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
				cgcGroup.PUT("/:id", middlewares.GuestJWT(), cgc.Update)
				cgcGroup.DELETE("/:id", middlewares.GuestJWT(), cgc.Delete)
			}

			tpc := new(controllers.TopicsController)
			tpcGroup := v1.Group("/topics")
			{
				tpcGroup.GET("", tpc.Index)
				tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
				tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
				tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
				tpcGroup.GET("/:id", tpc.Show)
			}

			lsc := new(controllers.LinksController)
			linksGroup := v1.Group("/links")
			{
				linksGroup.GET("", lsc.Index)
			}
		}
	}
}
