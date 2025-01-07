package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Print(nil)
	}
}

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "博客"
	indexData.Desc = "学习案例"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}
