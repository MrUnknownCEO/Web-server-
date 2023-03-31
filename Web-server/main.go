// filename: main.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//create handler functions
//home endpoint
func home(w http.ResponseWriter, r *http.Request) {
	//use the serverFile function to serve the html file to the client 
	http.ServeFile(w,r,"index.html")
}

//second handler function 
//greetings endpoint 
func greetings(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()//time function for local time 
	date := time.Now()//use time function again to get date of week
	formatTime := date.Format("Monday")//fromat the time to this particular way 
	formattedtime := currentTime.Format("3:04 PM")//format the time to this way as well
      fmt.Fprintf(w, "Right now is %s\n", formattedtime)
	fmt.Fprintf(w,"Enjoy the rest of your %s\n",formatTime)
}

//third handler function
func random(w http.ResponseWriter, r *http.Request) {
	//map that contains the quotes to randomized and sent as a response when a request is received
	quotes := map[int]string {
	      1: "Many Of Life's Failures Are People Who Did Not Realize How Close They Were To Success When They Gave Up. - (Thomas A. Edison)",
		2: "Don't Be The Person Who Says Yes With Their Mouth But No With Their Actions. - (Ryan Holiday)",
		3: "Life Is A Long Lesson In Humility. - (James M. Barrie)",
		4: "In Three Words I Can Sum Up Everything I've Learned About Life: It Goes On. - (Robert Frost)",
		5: "We Have Two Lives, And The Second Begins When We Realize We Only Have One. - (Confucius)",

	}

	//seeds the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	//chose a random response 
	response := quotes[rand.Intn(len(quotes))+1]

	//write the response to the client
	fmt.Fprintf(w,"%s", response)

}
func main() {
	//multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/greetings", greetings)
	mux.HandleFunc("/random", random)

	//server for the client
	log.Print("starting server on: 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
