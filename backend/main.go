package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const frontend = "../frontend/build/"

const api = "/api"

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.ServeFiles("/static/*filepath", http.Dir(frontend+"static"))

	router.POST(api+"/post", createPost)
	router.GET(api+"/post", readPost)
	router.PUT(api+"/post", updatePost)
	router.DELETE(api+"/post", deletePost)

	addPathsToRedirect(router, parseRoutes("../frontend/src/routes.json"))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, frontend+"index.html")
}

func addPathsToRedirect(router *httprouter.Router, paths []string) {
	for _, path := range paths {
		router.GET(path, redirectWithCookie)
	}
}

func redirectWithCookie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.SetCookie(w, &http.Cookie{Name: "redirect", Value: r.URL.Path, Path: "/"})
	http.Redirect(w, r, "/", http.StatusFound)
}

func createPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newPost postStruct

	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		log.Println(err)
	}

	postByID[newPost.ID] = newPost.Content
}

type postStruct struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type postList struct {
	Post postStruct `json:"post"`
}

func readPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	postID := r.FormValue("id")

	posts := getPostsByID([]string{postID})

	enc := json.NewEncoder(w)
	for _, post := range posts {
		enc.Encode(postStruct{Content: post})
	}
}

func updatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return
}

func deletePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return
}

var postByID = map[string]string{
	"1": "Hello this is post 1",
	"2": "Post 2 is good for you",
	"3": "But Post 3 is good for me",
}

func getPostsByID(postIDs []string) []string {
	var posts []string

	for _, id := range postIDs {
		posts = append(posts, postByID[id])
	}

	return posts

}
