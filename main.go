package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	pkg "github.com/MuhammadImran696/RSS-Reader-package"
	"github.com/gorilla/mux"
)

type rssFeeds struct {
	Items []pkg.RssItem `json:"items"`
}

type Urls struct {
	Array []string `json:"urls"`
}

func unmarsh(content []pkg.RssItem) rssFeeds {
	var feeds rssFeeds

	out, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	var items pkg.RssItem

	json.Unmarshal([]byte(out), &feeds.Items)

	for i := range content {
		items.Title = feeds.Items[i].Title
		items.Source = feeds.Items[i].Source
		items.SourceURL = feeds.Items[i].SourceURL
		items.Link = feeds.Items[i].Link
		items.PublishDate = feeds.Items[i].PublishDate
		items.Description = feeds.Items[i].Description

		feeds.Items = append(feeds.Items, items)
	}
	return feeds
}
func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	urlBody := string(reqBody)
	if urlBody == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	var urls Urls
	json.Unmarshal([]byte(reqBody), &urls)
	if len(urls.Array) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	response := pkg.Parse(urls.Array)
	feeds := unmarsh(response)
	json.NewEncoder(w).Encode(feeds)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/getdata", GetData).Methods("POST")
	log.Fatal(http.ListenAndServe(":9000", router))

}
