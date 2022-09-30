package main

import(
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/books", func (w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World")
	})

	log.Println("API is running on 4040")
	http.ListenandServe(":4040", router)
}