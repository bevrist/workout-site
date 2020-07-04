package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func showFrontend(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	// log.Println(r.URL.String()) //TODO: change form target to match the URL string

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		weight, _ := strconv.Atoi(r.FormValue("weight"))
		height, _ := strconv.Atoi(r.FormValue("height"))
		waist, _ := strconv.Atoi(r.FormValue("waist"))
		values := map[string]int{"weight": weight, "height": height, "waist": waist}
		jsonValue, _ := json.Marshal(values)
		_, _ = http.Post("http://example.com", "application/json", bytes.NewBuffer(jsonValue))
		// log.Println(string(jsonValue))

		http.ServeFile(w, r, "index.html")
	} else if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func main() {
	http.HandleFunc("/", showFrontend)

	fmt.Printf("listening at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
