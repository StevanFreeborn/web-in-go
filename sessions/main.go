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
	}

	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	session.Values[authenticationName] = true
}
