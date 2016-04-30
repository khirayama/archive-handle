package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joeshaw/envdecode"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

var (
	authClient *oauth.Client
	creds      *oauth.Credentials
)

func setUpAppAuth() {
	var ts struct {
		SecurityKey string `env:"SECURITY_KEY",required`
	}
}

func setUpGoogleAuth() {
	var ts struct {
		ConsumerKey    string `env:"GOOGLE_KEY",required`
		ConsumerSecret string `env:"GOOGLE_SECRET",required`
	}
	if err := envdecode.Decode(&ts); err != nil {
		log.Fatalln(err)
	}
	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token: ts.ConsumerKey,
			Secret: ts.ConsumerSecret
		}
	}
}

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
