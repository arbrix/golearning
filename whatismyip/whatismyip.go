package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func main()  {

//	URL := os.Args[1]
	URL := "http://httpbin.org/ip"

	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	result := make(map[string]string)
	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(&result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("My IP address is: %v", result["origin"])
}