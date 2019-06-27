package router

import (
	"golang-cell5/src/pkg/types/routes"
	"net/http"

	"github.com/go-xorm/xorm"

	AuthHandler "golang-cell5/src/controllers/auth"
	HomeHandler "golang-cell5/src/controllers/home"
	"log"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside main middleware.")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	AuthHandler.Init(db)
	HomeHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
		routes.Route{"AuthStore", "POST", "/auth/login", AuthHandler.Login},
		routes.Route{"AuthCheck", "GET", "/auth/check", AuthHandler.Check},
	}
}
