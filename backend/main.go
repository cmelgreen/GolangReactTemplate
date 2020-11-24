package main

import (
	"net/http"
	"log"

	"github.com/julienschmidt/httprouter"
)

const frontend = "../frontend/build/"

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.ServeFiles("/static/*filepath", http.Dir(frontend + "static"))

	addPathsToRedirect(router, parseRoutes("../frontend/src/routes.json"))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, frontend + "index.html")
}

func addPathsToRedirect(router *httprouter.Router, paths []string) {
	for _, path := range paths {
		router.GET(path, redirectWithCookie)
	}
}

func redirectWithCookie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.SetCookie(w, &http.Cookie{Name: "redirect", Value: r.URL.Path, Path:"/" })
	http.Redirect(w, r, "/", http.StatusFound)
}