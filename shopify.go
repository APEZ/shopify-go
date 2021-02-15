package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Param struct {
	Key string
	Value interface{}
}

type Store struct {
	Name string
	Password string
	Version string
}

// create new shopify instance
func New(name, password string) (store Store) {
	store.Name = name
	store.Password = password
	store.Version = "2021-04"
	return
}

// handle errors
func handleError(err error) {
	if err != nil {
		log.Fatalf("[shopify-go] Error: %v", err)
	}
}

// handle api request
func doRequest(store *Store, method, endpoint string, params []Param, body interface{}) (respBody []byte) {
	bodyJson, err := json.Marshal(body)
	handleError(err)
	queryString := url.Values{}
	if len(params) > 0 {
		for _, value := range params {
			queryString.Add(value.Key, fmt.Sprintf("%v", value.Value))
		}
	}
	endpoint, client := fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/%s?%s", store.Name, store.Version, endpoint, queryString.Encode()), &http.Client{}
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(bodyJson))
	handleError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Shopify-Access-Token", store.Password)
	response, err := client.Do(req)
	handleError(err)
	defer response.Body.Close()
	respBody, err = ioutil.ReadAll(response.Body)
	handleError(err)
	return
}
