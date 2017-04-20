package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"

	"encoding/json"
	"strconv"
)

type DataStr struct {
	Title string `json:"title"`
	Lines string `json:"lines"`
}

func CountLines(folder string) string {
	f, err := os.Open(folder)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	ScanF := bufio.NewScanner(f)

	c := 0
	for ScanF.Scan() {
		c += 1
	}
	return strconv.Itoa(c)
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	Data := DataStr{}

	folder := os.Args[1]

	BookName := folder + mux.Vars(r)["book"]

	Data.Title = mux.Vars(r)["book"]
	Data.Lines = CountLines(BookName)

	jsonStr, err := json.Marshal(Data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonStr))

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

func main() {
	//	folder := os.Args[1]
	//	folder := "./httplinecount/books/alice.txt"

	m := mux.NewRouter()
	m.HandleFunc("/books/{book}", BookHandler)
	http.ListenAndServe(":8080", m)
}
