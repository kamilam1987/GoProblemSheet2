//Author: Kmila Michel g00340498
//Problem set: Web applications in Go

package main

import (
	"net/http"
)

func main() {
	//code adapted from:https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
	// FileServe serves index.html from src folder
	//http.Handle function tells the http package to handle all requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/index.html")
	})

	http.ListenAndServe(":8080", nil)

}
