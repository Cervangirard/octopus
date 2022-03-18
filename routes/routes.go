package routes

import (
	"fmt"
	"net/http"
	"octopus/login"
	"octopus/shinyprocess"
	"time"
)

// type DataHtml map[string]interface{}

// Get ShinyProcess struct
var p shinyprocess.ShinyProcess

// Home route
func Home(w http.ResponseWriter, r *http.Request) {

	// Middleware this ??
	login.IsLogged(w, r)

	fmt.Printf("home route:")
	fmt.Printf(r.URL.Path)
	http.ServeFile(w, r, "views/index.html")

}

// Login route
func Login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/login.html")
}

// AppShiny route. Will launch also a ShinyProcess
func AppShiny(w http.ResponseWriter, r *http.Request) {

	login.IsLogged(w, r)

	p.LaunchApp(r)

	// Sleep to be sure that process is starting.
	// Should be something cleaner
	time.Sleep(3 * time.Second)

	http.ServeFile(w, r, "views/shiny.html")

}

// BackHome route
func BackHome(w http.ResponseWriter, r *http.Request) {

	p.KillSession()

	http.Redirect(w, r, "/", http.StatusFound)

}
