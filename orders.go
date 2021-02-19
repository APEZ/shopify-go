package shopify

import (
	"encoding/json"
)

type Orders struct {
	Orders []Order `json:"orders"`
}
type ClientDetails struct {
	AcceptLanguage string      `json:"accept_language"`
	BrowserHeight  int         `json:"browser_height"`
	BrowserIP      string      `json:"browser_ip"`
	BrowserWidth   int         `json:"browser_width"`
	SessionHash    interface{} `json:"session_hash"`
	UserAgent      string      `json:"user_agent"`
}
type TotalLineItemsPriceSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type TotalDiscountsSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type TotalShippingPriceSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type SubtotalPriceSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type TotalPriceSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type TotalTaxSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type OriginLocation struct {
	ID           int64  `json:"id"`
	CountryCode  string `json:"country_code"`
	ProvinceCode string `json:"province_code"`
	Name         string `json:"name"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Zip          string `json:"zip"`
}
type DiscountedPriceSet struct {
	ShopMoney        ShopMoney        `json:"shop_money"`
	PresentmentMoney PresentmentMoney `json:"presentment_money"`
}
type ShippingLines struct {
	ID                            int64              `json:"id"`
	Title                         string             `json:"title"`
	Price                         string             `json:"price"`
	Code                          string             `json:"code"`
	Source                        string             `json:"source"`
	Phone                         interface{}        `json:"phone"`
	RequestedFulfillmentServiceID interface{}        `json:"requested_fulfillment_service_id"`
	DeliveryCategory              interface{}        `json:"delivery_category"`
	CarrierIdentifier             interface{}        `json:"carrier_identifier"`
	DiscountedPrice               string             `json:"discounted_price"`
	PriceSet                      PriceSet           `json:"price_set"`
	DiscountedPriceSet            DiscountedPriceSet `json:"discounted_price_set"`
	DiscountAllocations           []interface{}      `json:"discount_allocations"`
	TaxLines                      []TaxLines         `json:"tax_lines"`
}
type BillingAddress struct {
	FirstName    string      `json:"first_name"`
	Address1     string      `json:"address1"`
	Phone        interface{} `json:"phone"`
	City         string      `json:"city"`
	Zip          string      `json:"zip"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	LastName     string      `json:"last_name"`
	Address2     string      `json:"address2"`
	Company      string      `json:"company"`
	Latitude     float64     `json:"latitude"`
	Longitude    float64     `json:"longitude"`
	Name         string      `json:"name"`
	CountryCode  string      `json:"country_code"`
	ProvinceCode string      `json:"province_code"`
}
type ShippingAddress struct {
	FirstName    string      `json:"first_name"`
	Address1     string      `json:"address1"`
	Phone        interface{} `json:"phone"`
	City         string      `json:"city"`
	Zip          string      `json:"zip"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	LastName     string      `json:"last_name"`
	Address2     string      `json:"address2"`
	Company      string      `json:"company"`
	Latitude     float64     `json:"latitude"`
	Longitude    float64     `json:"longitude"`
	Name         string      `json:"name"`
	CountryCode  string      `json:"country_code"`
	ProvinceCode string      `json:"province_code"`
}
type DefaultAddress struct {
	ID           int64       `json:"id"`
	CustomerID   int64       `json:"customer_id"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Company      string      `json:"company"`
	Address1     string      `json:"address1"`
	Address2     string      `json:"address2"`
	City         string      `json:"city"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	Zip          string      `json:"zip"`
	Phone        interface{} `json:"phone"`
	Name         string      `json:"name"`
	ProvinceCode string      `json:"province_code"`
	CountryCode  string      `json:"country_code"`
	CountryName  string      `json:"country_name"`
	Default      bool        `json:"default"`
}
type Customer struct {
	ID                        int64          `json:"id"`
	Email                     string         `json:"email"`
	AcceptsMarketing          bool           `json:"accepts_marketing"`
	CreatedAt                 string         `json:"created_at"`
	UpdatedAt                 string         `json:"updated_at"`
	FirstName                 string         `json:"first_name"`
	LastName                  string         `json:"last_name"`
	OrdersCount               int            `json:"orders_count"`
	State                     string         `json:"state"`
	TotalSpent                string         `json:"total_spent"`
	LastOrderID               int64          `json:"last_order_id"`
	Note                      interface{}    `json:"note"`
	VerifiedEmail             bool           `json:"verified_email"`
	MultipassIdentifier       interface{}    `json:"multipass_identifier"`
	TaxExempt                 bool           `json:"tax_exempt"`
	Phone                     interface{}    `json:"phone"`
	Tags                      string         `json:"tags"`
	LastOrderName             string         `json:"last_order_name"`
	Currency                  string         `json:"currency"`
	AcceptsMarketingUpdatedAt string         `json:"accepts_marketing_updated_at"`
	MarketingOptInLevel       string         `json:"marketing_opt_in_level"`
	TaxExemptions             []interface{}  `json:"tax_exemptions"`
	AdminGraphqlAPIID         string         `json:"admin_graphql_api_id"`
	DefaultAddress            DefaultAddress `json:"default_address"`
}
type PaymentDetails struct {
	CreditCardBin     string `json:"credit_card_bin"`
	AvsResultCode     string `json:"avs_result_code"`
	CvvResultCode     string `json:"cvv_result_code"`
	CreditCardNumber  string `json:"credit_card_number"`
	CreditCardCompany string `json:"credit_card_company"`
}
type Order struct {
	ID                     int64                  `json:"id"`
	Email                  string                 `json:"email"`
	ClosedAt               interface{}            `json:"closed_at"`
	CreatedAt              string                 `json:"created_at"`
	UpdatedAt              string                 `json:"updated_at"`
	Number                 int                    `json:"number"`
	Note                   interface{}            `json:"note"`
	Token                  string                 `json:"token"`
	Gateway                string                 `json:"gateway"`
	Test                   bool                   `json:"test"`
	TotalPrice             string                 `json:"total_price"`
	SubtotalPrice          string                 `json:"subtotal_price"`
	TotalWeight            int                    `json:"total_weight"`
	TotalTax               string                 `json:"total_tax"`
	TaxesIncluded          bool                   `json:"taxes_included"`
	Currency               string                 `json:"currency"`
	FinancialStatus        string                 `json:"financial_status"`
	Confirmed              bool                   `json:"confirmed"`
	TotalDiscounts         string                 `json:"total_discounts"`
	TotalLineItemsPrice    string                 `json:"total_line_items_price"`
	CartToken              interface{}            `json:"cart_token"`
	BuyerAcceptsMarketing  bool                   `json:"buyer_accepts_marketing"`
	Name                   string                 `json:"name"`
	ReferringSite          string                 `json:"referring_site"`
	LandingSite            string                 `json:"landing_site"`
	CancelledAt            interface{}            `json:"cancelled_at"`
	CancelReason           interface{}            `json:"cancel_reason"`
	TotalPriceUsd          string                 `json:"total_price_usd"`
	CheckoutToken          string                 `json:"checkout_token"`
	Reference              interface{}            `json:"reference"`
	UserID                 interface{}            `json:"user_id"`
	LocationID             interface{}            `json:"location_id"`
	SourceIdentifier       interface{}            `json:"source_identifier"`
	SourceURL              interface{}            `json:"source_url"`
	ProcessedAt            string                 `json:"processed_at"`
	DeviceID               interface{}            `json:"device_id"`
	Phone                  interface{}            `json:"phone"`
	CustomerLocale         string                 `json:"customer_locale"`
	AppID                  int                    `json:"app_id"`
	BrowserIP              string                 `json:"browser_ip"`
	ClientDetails          ClientDetails          `json:"client_details"`
	LandingSiteRef         interface{}            `json:"landing_site_ref"`
	OrderNumber            int                    `json:"order_number"`
	DiscountApplications   []interface{}          `json:"discount_applications"`
	DiscountCodes          []interface{}          `json:"discount_codes"`
	NoteAttributes         []interface{}          `json:"note_attributes"`
	PaymentGatewayNames    []string               `json:"payment_gateway_names"`
	ProcessingMethod       string                 `json:"processing_method"`
	CheckoutID             int64                  `json:"checkout_id"`
	SourceName             string                 `json:"source_name"`
	FulfillmentStatus      interface{}            `json:"fulfillment_status"`
	TaxLines               []TaxLines             `json:"tax_lines"`
	Tags                   string                 `json:"tags"`
	ContactEmail           string                 `json:"contact_email"`
	OrderStatusURL         string                 `json:"order_status_url"`
	PresentmentCurrency    string                 `json:"presentment_currency"`
	TotalLineItemsPriceSet TotalLineItemsPriceSet `json:"total_line_items_price_set"`
	TotalDiscountsSet      TotalDiscountsSet      `json:"total_discounts_set"`
	TotalShippingPriceSet  TotalShippingPriceSet  `json:"total_shipping_price_set"`
	SubtotalPriceSet       SubtotalPriceSet       `json:"subtotal_price_set"`
	TotalPriceSet          TotalPriceSet          `json:"total_price_set"`
	TotalTaxSet            TotalTaxSet            `json:"total_tax_set"`
	LineItems              []LineItems            `json:"line_items"`
	Fulfillments           []interface{}          `json:"fulfillments"`
	Refunds                []interface{}          `json:"refunds"`
	TotalTipReceived       string                 `json:"total_tip_received"`
	OriginalTotalDutiesSet interface{}            `json:"original_total_duties_set"`
	CurrentTotalDutiesSet  interface{}            `json:"current_total_duties_set"`
	AdminGraphqlAPIID      string                 `json:"admin_graphql_api_id"`
	ShippingLines          []ShippingLines        `json:"shipping_lines"`
	BillingAddress         BillingAddress         `json:"billing_address"`
	ShippingAddress        ShippingAddress        `json:"shipping_address"`
	Customer               Customer               `json:"customer"`
	PaymentDetails         PaymentDetails         `json:"payment_details,omitempty"`
}

// list multiple orders
func (store *Store) GetOrders(params ...Param) []Order {
	bytes := doRequest(store, "get", "/orders.json", params, nil)
	orders := Orders{}
	err := json.Unmarshal(bytes, &orders)
	handleError(err)
	return orders.Orders
}
