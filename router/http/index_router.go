package http

import (
	"github.com/ChaukeBasani/gofarm/controller/context/pages"
	"github.com/labstack/echo/v4"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}