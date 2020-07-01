package routers

import (
	"discord_club/controllers"
	. "discord_club/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})

	dashboardNS := beego.NewNamespace("/dashboard",
		// Checks if user is logged in
		beego.NSBefore(func(ctx *context.Context) {
			token, ok := ctx.Input.Session("token").(TokenData)
			if !ok {
				// User is not logged in
				ctx.Redirect(302, "/oauth/login")
				return
			}

			// Expose access token to the frontend
			ctx.Output.Cookie("token", token.AccessToken)
		}),
		beego.NSRouter("/", &controllers.DashboardController{}),
	)
	beego.AddNamespace(dashboardNS)
}