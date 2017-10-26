//Author: Kmila Michel g00340498
//Problem sheet 2 : Web applications in Go

package main

import (
	"net/http"
	"text/template"
)

//Code adapted from :https://golang.org/pkg/text/template/
type Game struct {
	Message string
}

//Function tempHandler
func tempHandler(w http.ResponseWriter, r *http.Request) {
	//Parse guess template file
	t, _ := template.ParseFiles("template/guess.html")
	//Sends message from guess template
	//code adapted from:https://www.youtube.com/watch?v=GTSq1VPPFco&feature=youtu.be
	t.Execute(w, Game{Message: "Guess a number between 1 and 20"})

} //End of tempHandler

//Function main
func main() {
	//code adapted from:https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
	// FileServe serves index.html from src folder
	//http.Handle function tells the http package to handle all requests from src folder (index.html)
	http.Handle("/", http.FileServer(http.Dir("./src")))

	//Handles requests from guess template
	http.HandleFunc("/guess", tempHandler)

	//http.Handle function tells the http package to handle all requests
	/*http.HandleFunc("/guess/", func(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./src/guess.html")*/

	http.ListenAndServe(":8080", nil)

} //End of function main
