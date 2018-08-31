package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ViewModel

// Resp Response VM
type Resp struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

// Handlers

func handleHello(w http.ResponseWriter, r *http.Request) {
	// Case 1: w.Write byte
	w.Write([]byte("Hello World"))

	// Case 2: fmt.Fprintf
	// fmt.Fprintf(w, "Hello World")

	// Case 3: io.Write
	// io.WriteString(w, "Hello World")
}

// http://localhost:8888/query?q=products&page=1
func handleQuery(w http.ResponseWriter, r *http.Request) {
	url := r.URL         // net/url.URL
	query := url.Query() // Values (map[string][]string)
	fmt.Fprintf(w, query.Encode())
	// q := query["q"]           // []string{“products”}
	// page := query.Get("page") // "1" 如果 没有传值 是 ""
}

func handleGetJSON(w http.ResponseWriter, r *http.Request) {
	status := "ok"
	msg := "result"
	result := Resp{Status: status, Result: msg}

	// Case 1: Set header and w.Write
	// js, err := json.Marshal(result)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)

	// Case 2: Encoder ResonseWriter
	enc := json.NewEncoder(w)
	err := enc.Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlePostForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", handleHello)
	http.HandleFunc("/query", handleQuery)
	http.HandleFunc("/getjson", handleGetJSON)
	http.ListenAndServe(":8888", nil)
}
