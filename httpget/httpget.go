package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {

	URL := os.Args[1]
//	URL := "http://httpbin.org/ip"

	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}