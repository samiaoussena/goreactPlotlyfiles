package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type item struct {
	Name           string
	Path           string
	Size           int64
	Downloaded_url string
}

func main() {
	/*
		directory := flag.String("d", ".", "the directory of static file to host")
		flag.Parse()
	*/
	r := mux.NewRouter()
	r.HandleFunc("/list", withCORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		files, err := ioutil.ReadDir("./JsonDir")
		if err != nil {
			log.Fatal(err)
		}
		var items []item
		for _, f := range files {
			res2A := item{
				Name:           f.Name(),
				Path:           "JsonDir/" + f.Name(),
				Size:           f.Size(),
				Downloaded_url: "http://localhost:8080/fdetails/" + f.Name(),
			}
			items = append(items, res2A)

		}
		b, _ := json.Marshal(items)
		w.Write(b)
	}))

	r.HandleFunc("/fdetails/{id}", withCORS(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		fName := vars["id"]
		fpath := "./JsonDir/" + fName
		fmt.Println(fpath)
		dat, _ := ioutil.ReadFile(fpath)
		w.Write(dat)

	}))
	http.ListenAndServe(":8080", r)
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		//w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		fn(w, r)
	}
}
