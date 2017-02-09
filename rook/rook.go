package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	rooter := mux.NewRouter().StrictSlash(true);

	rooter.HandleFunc("/", Index)
	rooter.HandleFunc("/bishops", BishopIndex)
	log.Fatal(http.ListenAndServe(":8080", rooter))
}

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
}

func BishopIndex(writer http.ResponseWriter, request *http.Request){
	tab := make([]string, 0)
	tab = append(tab,"url1","url2")

	fmt.Fprintln(writer, "Bishop index : ")
	for _, url := range tab{
		fmt.Fprintf(writer, "- %s \n", url)
	}
}
