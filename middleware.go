package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
	"regexp"
	"strings"
)

// RegisterMiddleware contains the middleware that runs for all routes
func RegisterMiddleware(router *chi.Mux) {
	corsMiddleware := cors.New(cors.Options{
		MaxAge:          300,
		AllowOriginFunc: handleCheckOrigin,
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:  []string{"Accept", "Content-Type", "C-CSRF-Token"},
		ExposedHeaders:  []string{"Link"},
	})
	router.Use(corsMiddleware.Handler)
}

func handleCheckOrigin(r *http.Request, origin string) bool {
	allowedOrigins := EnvDefault("ALLOWED_ORIGINS", "").(string)
	if allowedOrigins == "" {
		Error.Println("Can't check cross origin because the ALLOWED_ORIGINS env var is not defined")
		return false
	}
	allowedOrigins = strings.Replace(allowedOrigins, " ", "", -1)
	isAllowed := false
	for _, o := range strings.Split(allowedOrigins, ",") {
		originExp, _ := regexp.Compile(fmt.Sprintf("^(http[s]?://)?(%s)$", o))
		if originExp.MatchString(origin) {
			isAllowed = true
		}
	}
	return isAllowed
}
