package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {
	var addr = flag.String("addr", ":8080", "app address")
	flag.Parse() // read flag

	// setup Gomniauth
	SECURITY_KEY := os.Getenv("SECURITY_KEY")
	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_SECRET_KEY := os.Getenv("GOOGLE_SECRET_KEY")

	CALLBACK_PATH := "http://localhost:8080/auth/callback"

	gomniauth.SetSecurityKey(SECURITY_KEY)
	gomniauth.WithProviders(google.New(GOOGLE_CLIENT_ID, GOOGLE_SECRET_KEY, CALLBACK_PATH+"/google"))

	// if publish static files
	// http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("/assets/"))))

	// routing
	http.Handle("/", MustAuth(&templateHandler{filename: "main.html"}, &templateHandler{filename: "home.html"}))
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	// start web server
	log.Println("start web server. port:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
