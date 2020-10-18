package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1", v1)
	http.HandleFunc("/api/v2", v2)
	// http.HandleFunc("/api/v3", foo)

	http.ListenAndServe(":3000", nil)
}

func v1(w http.ResponseWriter, r *http.Request) {
	// content, err := ioutil.ReadFile("data/resp_data/top.geojson")
	var p = getAllPoint()

	w.Header().Set("Content-Type", "application/json")
	w.Write(convert_geojson(p))

}
func v2(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("data/resp_data/top.geojson")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(content)

}

// func v3(w http.ResponseWriter, r *http.Request) {
// 	content, err := ioutil.ReadFile("testdata/hello")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)

// }
