package main

import (
	"fmt"
	"log"
	"net/http"
	"octopus/login"
	"octopus/routes"

	"github.com/gorilla/mux"
	"github.com/kataras/rewrite"
)

func main() {

	// NOT USED FOR NOW
	// database.CreateUsers()

	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// HOME
	r.HandleFunc("/", routes.Home)

	// Signin
	r.HandleFunc("/signin", login.Signin)

	// Login route
	r.HandleFunc("/login", routes.Login)

	// AppShiny route, launch a shiny process
	r.HandleFunc("/launchapp", routes.AppShiny)

	// back home route
	r.HandleFunc("/backhome", routes.BackHome)

	// Rewrtie options for shiny
	opts := rewrite.Options{
		RedirectMatch: []string{
			fmt.Sprintf("301 /shiny_router/(.*) %s/", "http://localhost:8000"),
		},
	}
	// Initialize the Rewrite Engine.
	rw, err := rewrite.New(opts)

	if err != nil {
		log.Fatal(err)
	}

	// Rewrite route to be sure to get shinyapp
	r.HandleFunc("/shiny_router", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/shiny_router/", http.StatusFound)
	})

	// On the default page we will simply serve our static index page.
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Our application will run on port 8080. Here we declare the port and pass in our router.
	// TODO add port flag in commad line to launch with port or env var.. dont know yet
	fmt.Printf("Go visit http://localhost:8080")
	http.ListenAndServe(":8080", rw.Handler(r))

}
