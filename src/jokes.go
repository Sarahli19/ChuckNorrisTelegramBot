package main

import "net/http"
import "log"
import (
	"encoding/json"
	"fmt"
)


func getJoke() {

	resp, err := http.Get("http://tambal.azurewebsites.net/joke/random")
	if err != nil {
		log.Panic(err)
	}

	//json.Marshal(resp)
	fmt.Printf(json.Marshal(resp))
}
