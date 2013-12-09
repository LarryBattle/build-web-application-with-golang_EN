//@TODO test.
//@todo Add front end gui ?

package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "Get User %s", uid)
}

func modifyUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "Modify User %s", uid)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "Delete User %s", uid)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprint(w, "Add User %s", uid)
}

var PORT = "8088"

func main() {
	mux := routes.New()
	mux.Get("/User/:uid", getUser)
	mux.Post("/User/:uid", modifyUser)
	mux.Del("/User/:uid", deleteUser)
	mux.Put("/User/:uid", addUser)
	http.Handle("/", mux)
	fmt.Println("Listing to localhost:" + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
