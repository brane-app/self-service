package main

import (
	"github.com/brane-app/tools-library/middleware"
	"github.com/gastrodon/groudon/v2"

	"os"
)

var (
	prefix = os.Getenv("PATH_PREFIX")

	routeRoot = "^" + prefix + "/$"
)

func register_handlers() {
	groudon.AddMiddleware("GET", routeRoot, middleware.MustAuth)

	groudon.AddHandler("GET", routeRoot, getSelf)
}
