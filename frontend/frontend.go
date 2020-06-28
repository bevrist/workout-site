package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// log.Println(r.URL.String()) TODO: change form target to match the URL string

	if r.Method == "POST" {
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		weight := r.FormValue("weight")
		height := r.FormValue("height")
		waist := r.FormValue("waist")
		fmt.Fprintf(w, "weight = %s\n", weight)
		fmt.Fprintf(w, "height = %s\n", height)
		fmt.Fprintf(w, "waist = %s\n", waist)
	} else {
		http.ServeFile(w, r, "index.html")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("listening at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
