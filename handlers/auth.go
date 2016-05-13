package handlers

import (
	"encoding/json"
	//"fmt"
	//"io"
	//"io/ioutil"
	"net/http"
	"github.com/gokrokvertskhov/gauth"
	"github.com/gokrokvertskhov/gauth/provider/github"
	"config"
	//"strconv"

	//"github.com/gorilla/mux"
)
var prov gauth.Provider
func GetProvider() gauth.Provider {
	if prov == nil {
		prov = github.New(config.Conf.Auth.Client_id, config.Conf.Auth.Client_key, "http://localhost:8080/auth/callback")
	}
  return prov
}
func AuthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Get access token from verification code.
	provider := GetProvider()
	token, err := provider.Authorize(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the authenticating user
	user, err := provider.User(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal user info and write to the response.
	output, _ := json.Marshal(user)
	w.Write([]byte(output))
}

func Login(w http.ResponseWriter, r *http.Request) {
	provider := GetProvider()
	authURL, _, err := provider.AuthURL("state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}