package main

import (
	"crypto/sha1"
	"fmt"
	// "html"
	"net/http"
	// "regexp"
	// "net/url"
	b64 "encoding/base64"
	"github.com/gorilla/mux"
)

func DecodeUrl(encoded_str string) (string, error) {
	decoded_str, err := b64.StdEncoding.DecodeString(encoded_str)
	return string(decoded_str), err
}

func EncodeUrl(string_url string) (string, error) {
	return b64.StdEncoding.EncodeToString([]byte(string_url)), nil
}

func GetSHA1HashedURL(long_url string) []byte {
	hash := sha1.New()
	hash.Write([]byte(long_url))
	return hash.Sum(nil)
}

func GetShortUrl(w http.ResponseWriter, r *http.Request) {
	data := GetSHA1HashedURL(long_url)
	short_url, err := EncodeUrl(fmt.Sprintf("%x", data))
	if err != nil {
		fmt.Printf("Could not convert url; error: %s", err)
		return fmt.Sprintf("%s", err)
	}
	w.Write([]byte(short_url))
}

func GetLongUrl(w http.ResponseWriter, r *http.Request) {
	long_url := "dummy"
	w.Write([]byte(long_url))
}

func router(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	route_type := params["route_type"]

	switch route_type {
		case "encode":
			GetShortUrl(w, r)
		case "decode":
			GetLongUrl(w, r)
		default:
			w.Write([]byte("Invalid choice."))
	}
}

func main() {
	CURRENT_HOST := "http://sh.rt"
	r := mux.NewRouter()
	http.HandleFunc("/", router)
	// http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// 	fmt.Fprintf(w, "Hello, long_url: google.com, short_url: %q", CURRENT_HOST + "/" + html.EscapeString(GetShortUrl("google.com"))[:6])
	// }))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("Error")
	}
}

