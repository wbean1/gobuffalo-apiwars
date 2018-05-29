package actions

import (
        "fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/rs/cors"
	"github.com/wbean1/apiwars/models"

	"github.com/markbates/goth/gothic"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env: ENV,
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_apiwars_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		app.GET("/", HomeHandler)

		app.Use(SetCurrentUser)
		app.Use(Authorize)
		app.Middleware.Skip(Authorize, HomeHandler)

		auth := app.Group("/auth")
		bah := buffalo.WrapHandlerFunc(gothic.BeginAuthHandler)
		auth.GET("/{provider}", bah)
		auth.GET("/{provider}/callback", AuthCallback)
		auth.DELETE("", AuthDestroy)
		auth.Middleware.Skip(Authorize, bah, AuthCallback)

                app.GET("/protected", Protect)

	}

	return app
}

func Protect (c buffalo.Context) error {
        m := fmt.Sprintf("super secret info here, userid: %s", c.Value("current_user"))
        return c.Render(200, r.JSON(map[string]string{"message": m}))

}
