package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func printRemoteAddr() {
	conn, err := net.Dial("tcp", "<somehost>.local:8080")
	fmt.Println(conn, err)
	if err != nil {
		panic(err)
	}
	fmt.Println(conn.RemoteAddr())
}
