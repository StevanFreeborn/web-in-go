package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	cookieName         = "cookie-name"
	authenticationName = "authenticated"
	key                = []byte("super-secret-key")
	store              = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	if auth, ok := session.Values[authenticationName].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	session.Values[authenticationName] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	session.Values[authenticationName] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
