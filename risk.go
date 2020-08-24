package goshopify

import "fmt"

// RiskService is an interface for interfacing with the order risk endpoints
// of the Shopify API.
// https://shopify.dev/docs/admin-api/rest/reference/orders/order-risk
type RiskService interface {
	List(int64, interface{}) ([]Risk, error)
	Get(int64, int64, interface{}) (*Risk, error)
	Create(int64, Risk) (*Risk, error)
	Update(int64, Risk) (*Risk, error)
	Delete(int64, int64) error
}

// Risk represents a Shopify order risk.
type Risk struct {
	ID              int64  `json:"id,omitempty"`
	OrderID         int64  `json:"order_id,omitempty"`
	CheckoutID      *int64 `json:"checkout_id,omitempty"`
	Source          string `json:"source,omitempty"`
	Score           string `json:"score,omitempty"`
	Recommendation  string `json:"recommendation,omitempty"`
	Display         bool   `json:"display,omitempty"`
	CauseCancel     bool   `json:"cause_cancel,omitempty"`
	Message         string `json:"message,omitempty"`
	MerchantMessage string `json:"merchant_message,omitempty"`
}

// RiskOp handles communication with the order risk
// related methods of the Shopify API.
type RiskOp struct {
	client *Client
}

// RiskResource represents the result from the risks/X.json endpoint
type RiskResource struct {
	Risk *Risk `json:"risk"`
}

// RisksResource represents the result from the risks.json endpoint
type RisksResource struct {
	Risks []Risk `json:"risks"`
}

func (s *RiskOp) List(orderID int64, options interface{}) ([]Risk, error) {
	path := fmt.Sprintf("%s/%d/risks.json", ordersBasePath, orderID)
	resource := new(RisksResource)
	err := s.client.Get(path, resource, options)
	return resource.Risks, err
}

func (s *RiskOp) Get(orderID int64, riskID int64, options interface{}) (*Risk, error) {
	path := fmt.Sprintf("%s/%d/risks/%d.json", ordersBasePath, orderID, riskID)
	resource := new(RiskResource)
	err := s.client.Get(path, resource, options)
	return resource.Risk, err
}

func (s *RiskOp) Create(orderID int64, risk Risk) (*Risk, error) {
	path := fmt.Sprintf("%s/%d/risks.json", ordersBasePath, orderID)
	wrappedData := RiskResource{Risk: &risk}
	resource := new(RiskResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Risk, err
}

func (s *RiskOp) Update(orderID int64, risk Risk) (*Risk, error) {
	path := fmt.Sprintf("%s/%d/risks/%d.json", ordersBasePath, orderID, risk.ID)
	wrappedData := RiskResource{Risk: &risk}
	resource := new(RiskResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Risk, err
}

func (s *RiskOp) Delete(orderID int64, riskID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d/risks/%d.json", ordersBasePath, orderID, riskID))
}
