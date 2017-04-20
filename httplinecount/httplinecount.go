package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"

	"encoding/json"
	"strconv"
"errors"
)

type DataStr struct {
	Title string `json:"title"`
	Lines string `json:"lines"`
}

func CountLines(folder string) (string, error) {
	f, err := os.Open(folder)
	if err != nil {
		return "", errors.New("Can't open the file")
	}

	defer f.Close()

	ScanF := bufio.NewScanner(f)

	c := 0
	for ScanF.Scan() {
		c += 1
	}
	return strconv.Itoa(c), nil
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	Data := DataStr{}

	folder := os.Args[1]

	BookName := folder + mux.Vars(r)["book"]

	Data.Title = mux.Vars(r)["book"]

	if _, err := CountLines(BookName); err != nil {
		fmt.Fprintf(w, fmt.Sprint(err))
		return
		}

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
