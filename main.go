package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", authMetadata)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func authMetadata(writer http.ResponseWriter, req *http.Request) {
	authMetadataResponse := struct {
		Issuer                string   `json:"issuer"`
		AuthorizationEndpoint string   `json:"authorization_endpoint"`
		TokenEndpoint         string   `json:"token_endpoint"`
		CodeChallengeMethods  []string `json:"code_challenge_methods_supported"`
	}{
		Issuer:                "https://indieauth.example.com/",
		AuthorizationEndpoint: "https://indieauth.example.com/auth",
		TokenEndpoint:         "https://indieauth.example.com/token",
		CodeChallengeMethods:  []string{"S256"},
	}
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(authMetadataResponse); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
