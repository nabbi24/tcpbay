package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Hello, world.")
	} else {
		s := strings.Split(r.URL.Path, "/")
		if len(s) == 2 || s[2] == "" {
			fmt.Fprintf(w, fmt.Sprintf("Hello, %s.", s[1]))
		} else {
			fmt.Fprintf(w, fmt.Sprintf("Hello, %s.\nYou speak a lot!", s[1]))
		}
	}
}

func main() {
	port := 80
	if 2 <= len(os.Args) {
		if v, err := strconv.Atoi(os.Args[1]); err == nil {
			port = v
		}
	}

	http.HandleFunc("/", HelloWorld)

	fmt.Printf("Server starts to listen with port: %v... [Enter Ctrl+C when exit]\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
