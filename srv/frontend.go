package srv

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

func StartWebserver() {
	n := negroni.New()

	// static folders for assets
	n.Use(negroni.NewStatic(http.Dir("/www/assets/")))

	// recovery
	n.Use(negroni.NewRecovery())
}
