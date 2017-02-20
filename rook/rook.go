package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"net/url"
	"fmt"
	"html"
)

type Rook struct{
	bishops *Bishops
}

func main() {
	rooter := mux.NewRouter().StrictSlash(true);
	rook := Rook{
		CreateBishops(),
	}
	rook.bishops.AddServerFromUrl(url.URL{
		Scheme: "http",

		Host:   "www.google.com",

		Path:   "/yolo",
	})
	fmt.Println("")
	fmt.Printf("Current servers : %s\n", rook.bishops.String())
	rooter.HandleFunc("/", rook.Index)
	rooter.HandleFunc("/bishops", rook.BishopIndex)
	rooter.HandleFunc("/bishops/{url}", rook.AddServerFromExternal)
	log.Fatal(http.ListenAndServe(":8080", rooter))
}

func (r *Rook) Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
}

func (r *Rook) BishopIndex(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer, "servers in bishops :")
	fmt.Fprintf(writer, r.bishops.String())
}

func (r *Rook) AddServerFromExternal(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer, "add a server in bishop :")
	vars := mux.Vars(request)
	urls := vars["url"]
	fmt.Fprintln(writer, "url to add :", urls)

	r.bishops.AddServerFromUrl(url.URL{
		Scheme: "http",

		Host: urls,
	})

	fmt.Fprintf(writer, r.bishops.String())
}