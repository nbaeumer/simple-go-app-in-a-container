package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/addition", additionHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func addition(num1 string, num2 string) int {
	i, err := strconv.Atoi(num1)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	j, err := strconv.Atoi(num2)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	return (i + j)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	type HealthResp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	resp := HealthResp{"200", "Healthy"}

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func additionHandler(w http.ResponseWriter, r *http.Request) {

	num1, ok := r.URL.Query()["num1"]
	if !ok || len(num1) < 1 {
		fmt.Fprintf(w, "Error: Url Param 'num1' is missing")
		return
	}

	num2, ok := r.URL.Query()["num2"]
	if !ok || len(num2) < 1 {
		fmt.Fprintf(w, "Error: Url Param 'num2' is missing")
		return
	}

	val1 := num1[0]
	val2 := num2[0]

	fmt.Fprintf(w, "Input: %s, %s\n", string(val1), val2)
	w.(http.Flusher).Flush()

	sum := addition(val1, val2)
	v := strconv.Itoa(sum)

	delay, ok := r.URL.Query()["delay"]
	if ok && len(delay) > 0 {
		val3 := delay[0]

		faktor, err := strconv.Atoi(val3)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		time.Sleep(time.Duration(faktor) * time.Millisecond)

	}

	fmt.Fprintf(w, "Addition: %s + %s = %s\n", val1, val2, v)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Incorrect Url!", 405)
}
