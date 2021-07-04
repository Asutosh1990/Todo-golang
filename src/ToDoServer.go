package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"weekend_project/Todo-golang/src/services"
)

var td = services.NewToDoService()

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "Welcome to HomePage!<br/><br/>")

	items := td.ListToDoitem()

	if len(items) > 0 {
		fmt.Fprintf(w, "Your Task for today are:<br/><ul>")
		for _, v := range items {
			fmt.Fprintf(w, "<li>%s</li>", v.Description)
		}
	} else {
		fmt.Fprintf(w, "No new task for today")
	}
	fmt.Fprintf(w, "</ul>")
	fmt.Println("End Point Hit :Home Page")
}

//Mapping all http request here
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)                                                     //homePage mapping
	myRouter.HandleFunc("/todo", td.HandleListToDoItems).Methods(http.MethodGet)           //Get http request mapping
	myRouter.HandleFunc("/todo", td.HandlePostToDoItems).Methods(http.MethodPost)          //Post http request mapping
	myRouter.HandleFunc("/todo/{id}", td.HandleDeleteToDoItems).Methods(http.MethodDelete) //Delete http request mapping
	myRouter.HandleFunc("/todo/{id}", td.HandleGetToDoItem).Methods(http.MethodGet)        //Get mapping
	myRouter.HandleFunc("/todo/{id}", td.HandlePatchToDoItem).Methods(http.MethodPatch)    //Patch  mapping

	log.Fatal(http.ListenAndServe("localhost:8080", myRouter))
}

func main() {
	handleRequest()
}
