package shopify

import (
	"encoding/json"
	"fmt"
)

type FulfillmentRoot struct {
	Fulfillment Fulfillment `json:"fulfillment"`
}
type Properties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type AmountSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type DiscountAllocations struct {
	Amount                   string    `json:"amount"`
	DiscountApplicationIndex int       `json:"discount_application_index"`
	AmountSet                AmountSet `json:"amount_set"`
}
type Receipt struct {}
type Fulfillment struct {
	ID                int           `json:"id"`
	OrderID           int           `json:"order_id"`
	NotifyCustomer 	  bool 			`json:"notify_customer"`
	Status            string        `json:"status"`
	CreatedAt         string        `json:"created_at"`
	Service           string        `json:"service"`
	UpdatedAt         string        `json:"updated_at"`
	TrackingCompany   string        `json:"tracking_company"`
	ShipmentStatus    interface{}   `json:"shipment_status"`
	LocationID        int           `json:"location_id"`
	LineItems         []LineItems   `json:"line_items"`
	TrackingNumber    interface{}   `json:"tracking_number"`
	TrackingNumbers   []interface{} `json:"tracking_numbers"`
	TrackingURL       string        `json:"tracking_url"`
	TrackingUrls      []string      `json:"tracking_urls"`
	Receipt           Receipt       `json:"receipt"`
	Name              string        `json:"name"`
	AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
}

func (store *Store) CreateFulfillment(orderId int, fulfilment Fulfillment) Fulfillment {
	fulfilmentParent := FulfillmentRoot{fulfilment}
	response := doRequest(store, "post", fmt.Sprintf("/orders/%v/fulfillments.json", orderId), nil, fulfilmentParent)
	data := FulfillmentRoot{}
	err := json.Unmarshal(response, &data)
	handleError(err)
	return data.Fulfillment
}
