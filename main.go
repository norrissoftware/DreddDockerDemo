package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	CurrentItems ItemsWrapper
	validToken   string
	validUser    = "Carlos"
)

func init() {
	CurrentItems = ItemsWrapper{
		Items: []Item{
			Item{"dog", "brown"},
			Item{"cat", "white"},
		},
	}
	salt := os.Getenv("DEMOSHASALT")
	if salt == "" {
		log.Fatal("illegal salt")
	}
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, validUser)
	validToken = fmt.Sprintf("%x", h.Sum(nil))
}

type ItemsWrapper struct {
	Items []Item `json:"items"`
}

type Item struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func main() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/items").HandlerFunc(GetAllHandler)
	r.Methods("GET").Path("/items/{name}").HandlerFunc(GetHandler)
	r.Methods("POST").Path("/items").HandlerFunc(authMiddleware(PostHandler))
	r.Methods("GET").Path("/nope").HandlerFunc(NopeHandler)
	log.Fatal(http.ListenAndServe(":8080", jsonMiddleware(r)))
}

func jsonMiddleware(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})

}
func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CurrentItems)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		http.NotFound(w, r)
	}
	for _, v := range CurrentItems.Items {
		if v.Name == name {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.NotFound(w, r)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	i := Item{}
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()))
		return
	}
	CurrentItems.Items = append(CurrentItems.Items, i)
	io.WriteString(w, `{"success": true}`)
}

func authMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Auth")
		if auth != validToken {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error": "You aren't Carlos"}`)
			return
		}
		h(w, r)
	})

}

func NopeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, `{"error": "this request will never work"}`)
}
