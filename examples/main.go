package main 

import (
	"log"
	"net/http"

	"github.com/nguyendangminh/gotpl"
)

var tplHelper *gotpl.Helper

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["user"] = "minhnd"
	tplHelper.Render(w, "index", data)
}

func main() {
	var err error
	tplHelper, err = gotpl.New("templates")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", IndexHandler)
	log.Println("Server is running at http://localhost:1203")
	log.Println(http.ListenAndServe(":1203", nil))
}