package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/aarol/reload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	user "github.com/zachatrocity/htmx-hyperscript-starter/api/routes"
)

func main() {
	isDevelopment := flag.Bool("dev", true, "Development mode")
	port := flag.String("port", "3000", "Port to serve the app")
	flag.Parse()

	router := echo.New()

	// Add Middlewares Here
	// e.Use(middleware.Logger())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:4000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// server the public folder as static
	router.Static("/", "public/")

	api := router.Group("/api")
	{
		// htmx components
		api.GET("/components/*", func(c echo.Context) error {
			if *isDevelopment {
				fmt.Println("Component Requested: " + c.Request().URL.Path)
			}
			component := strings.ReplaceAll(c.Request().URL.Path, "/api/components/", "")
			// yet the cache for dev
			c.Response().Header().Set("Cache-Control", "no-store")
			return c.File("public/components/" + component + ".html")
		})
		// all other API requests
		api.GET("/users/:id", user.GetUser)
	}

	// hot reload from aarol/reload
	if *isDevelopment {
		// Watch for HTML changes in the public folder to trigger browser reload
		reload := reload.New("public/")

		// reload.OnReload = func() {
		// build templates if that's your thing
		// }
		router.GET("/reload_ws", echo.WrapHandler(reload.Handle(http.DefaultServeMux)))

		fmt.Println("Hot Reload Enabled...")
	}

	fmt.Printf("Listening on port %s\n", *port)
	router.Logger.Fatal(router.Start(":" + *port))
}
