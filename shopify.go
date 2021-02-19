package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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

type ShopMoney struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}
type PresentmentMoney struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}
type PriceSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type TotalDiscountSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type TaxLines struct {
	Title    string   `json:"title"`
	Price    string   `json:"price"`
	Rate     float64  `json:"rate"`
	PriceSet PriceSet `json:"price_set"`
}
type LineItems struct {
	ID                         int                   `json:"id"`
	VariantID                  int                   `json:"variant_id"`
	Title                      string                `json:"title"`
	Quantity                   int                   `json:"quantity"`
	Sku                        string                `json:"sku"`
	VariantTitle               string                `json:"variant_title"`
	Vendor                     interface{}           `json:"vendor"`
	FulfillmentService         string                `json:"fulfillment_service"`
	ProductID                  int                   `json:"product_id"`
	RequiresShipping           bool                  `json:"requires_shipping"`
	Taxable                    bool                  `json:"taxable"`
	GiftCard                   bool                  `json:"gift_card"`
	Name                       string                `json:"name"`
	VariantInventoryManagement string                `json:"variant_inventory_management"`
	Properties                 []Properties          `json:"properties"`
	ProductExists              bool                  `json:"product_exists"`
	FulfillableQuantity        int                   `json:"fulfillable_quantity"`
	Grams                      int                   `json:"grams"`
	Price                      string                `json:"price"`
	TotalDiscount              string                `json:"total_discount"`
	FulfillmentStatus          interface{}           `json:"fulfillment_status"`
	PriceSet                   PriceSet              `json:"price_set"`
	TotalDiscountSet           TotalDiscountSet      `json:"total_discount_set"`
	DiscountAllocations        []DiscountAllocations `json:"discount_allocations"`
	Duties                     []interface{}         `json:"duties"`
	AdminGraphqlAPIID          string                `json:"admin_graphql_api_id"`
	TaxLines                   []TaxLines            `json:"tax_lines"`
	OriginLocation             OriginLocation   	 `json:"origin_location"`
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
	if body == nil {
		bodyJson = []byte{}
	}
	handleError(err)
	queryString := url.Values{}
	if len(params) > 0 {
		for _, value := range params {
			queryString.Add(value.Key, fmt.Sprintf("%v", value.Value))
		}
	}
	endpoint, client := fmt.Sprintf("https://%s.myshopify.com/admin/api/%s%s?%s", store.Name, store.Version, endpoint, queryString.Encode()), &http.Client{}
	req, err := http.NewRequest(strings.ToUpper(method), endpoint, bytes.NewBuffer(bodyJson))
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