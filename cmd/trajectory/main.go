package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"soundspace/internal/api"
	"soundspace/internal/store"
)

func main() {
	path := "soundspace.db"
	if v := os.Getenv("SOUNDSPACE_DB"); v != "" {
		path = v
	}
	s, err := store.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	server := api.New(s)
	addr := ":8080"
	if v := os.Getenv("SOUNDSPACE_ADDR"); v != "" {
		addr = v
	}
	fmt.Println("soundspace listening on", addr)
	log.Fatal(http.ListenAndServe(addr, server.Handler()))
}
