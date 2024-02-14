package server

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	router *echo.Echo
}

func Start(port string) error {
	server := &Server{
		router: echo.New(),
	}

	server.router.Use(middleware.Logger())
	server.router.Use(middleware.Recover())

	// Setup HTTPS redirection
	server.router.Pre(middleware.HTTPSNonWWWRedirect())
	server.router.Pre(middleware.HTTPSRedirect())

	// Add Secure middleware
	server.router.Use(middleware.Secure())

	// Setup proxy
	url1, err := url.Parse("http://localhost:12000")
	if err != nil {
		server.router.Logger.Fatal(err)
	}
	url2, err := url.Parse("http://localhost:12000")
	if err != nil {
		server.router.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
		{
			URL: url2,
		},
	}
	server.router.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	server.router.AutoTLSManager.HostPolicy = autocert.HostWhitelist("ranaria.store")
	server.router.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	return server.router.StartAutoTLS(port)
}
