package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	values := url.Values{}
	values.Add("city", "400040")
	resp, err := http.Get("http://weather.livedoor.com/forecast/webservice/json/v1" + "?" + values.Encode())

	if err != nil {
		fmt.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func main() {
	http.HandleFunc("/weather", weatherHandler)
	http.Handle("/", &templateHandler{filename: "api.html"})
	http.ListenAndServe(":8000", nil)
}
