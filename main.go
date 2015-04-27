package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mickelsonm/storymaker/controllers/stories"
)

var (
	listenAddr = flag.Int("port", 4000, "http listen port")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()

	//r.HandleFunc("/category/{id}", categories.DeleteCategory).Methods("DELETE")
	//r.HandleFunc("/category/{id}", categories.GetCategory).Methods("GET")
	//r.HandleFunc("/category/{id}", categories.UpdateCategory).Methods("PUT")
	//r.HandleFunc("/category", categories.AddCategory).Methods("POST")
	r.HandleFunc("/timelines", stories.GetAllTimelines).Methods("GET")

	r.HandleFunc("/healthstatus", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK: It's all good."))
	})
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Dude, it's an API."))
	})

	n := negroni.New(negroni.NewRecovery())
	n.Run(fmt.Sprintf(":%d", *listenAddr))
}
