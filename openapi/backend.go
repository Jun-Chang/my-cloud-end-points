package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request Method: %s\n", r.Method)
	fmt.Fprintf(w, "Request Path: %s\n", r.URL.Path)
	r.ParseForm()
	for key, values := range r.Form {
		fmt.Fprintf(w, "Request Parameter: %s\n", key)
		for _, value := range values {
			fmt.Fprintf(w, "Value: %s\n", url.QueryEscape(value))
		}
	}
	userInfo := r.Header.Get("X-Endpoint-API-UserInfo")
	fmt.Fprintf(w, "Raw X-Endpoint-API-UserInfo: %s\n", userInfo)
	decoded, err := base64.StdEncoding.DecodeString(userInfo)
	if err != nil {
		fmt.Fprintf(w, "Error decoding X-Endpoint-API-UserInfo: %v\n", err)
	} else {
		fmt.Fprintf(w, "Decoded X-Endpoint-API-UserInfo: %s\n", string(decoded))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
