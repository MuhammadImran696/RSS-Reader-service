package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	pkg "github.com/MuhammadImran696/RSS-Reader-package"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/getdata", GetData).Methods("POST")

	return r
}
func TestGetDataEndpoint(t *testing.T) {

	f := []string{"http://feeds.nytimes.com/nyt/rss/HomePage", "https://www.latimes.com/world/rss2.0.xml"}
	data, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	request, _ := http.NewRequest("POST", "/getdata", reader)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
func TestEmptyEndpoint(t *testing.T) {

	f := []string{}
	data, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	request, _ := http.NewRequest("POST", "/getdata", reader)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code, "OK response is expected")
}
func TestUnmarshMethod(t *testing.T) {
	urls := []string{"http://feeds.nytimes.com/nyt/rss/HomePage", "https://www.latimes.com/world/rss2.0.xml"}
	// var data []RssItem
	data := pkg.Parse(urls)
	if len(data) == 0 {
		t.Fatalf(`No data Found`)
	}
}
