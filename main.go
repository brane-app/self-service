package main

import (
	"github.com/gastrodon/groudon"
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monkelib/middleware"

	"net/http"
	"os"
)

func main() {
	monkebase.Connect(os.Getenv("MONKEBASE_CONNECTION"))
	groudon.RegisterMiddleware(middleware.MustAuth)

	groudon.RegisterHandler("GET", "^/$", getSelf)
	http.Handle("/", http.HandlerFunc(groudon.Route))
	http.ListenAndServe(":8000", nil)
}
