package goshopify

import "fmt"

// RefundService is an interface for interfacing with the refunds endpoints of
// the Shopify API.
// See: https://shopify.dev/docs/admin-api/rest/reference/orders/refund
type RefundService interface {
	List(int64, interface{}) ([]Refund, error)
	Get(int64, int64, interface{}) (*Refund, error)
	Caculate(int64, Refund) (*Refund, error)
	Create(int64, Refund) (*Refund, error)
}

// RefundServiceOp handles communication with the refund related methods of the
// Shopify API.
type RefundServiceOp struct {
	client *Client
}

// RefundResource represents the result from the orders/X/refunds/Y.json endpoint
type RefundResource struct {
	Refund *Refund `json:"refund"`
}

// RefundsResource represents the result from the orders/X/refunds.json endpoint
type RefundsResource struct {
	Refunds []Refund `json:"refunds"`
}

// List refunds
func (s *RefundServiceOp) List(orderID int64, options interface{}) ([]Refund, error) {
	path := fmt.Sprintf("%s/%d/refunds.json", ordersBasePath, orderID)
	resource := new(RefundsResource)
	err := s.client.Get(path, resource, options)
	return resource.Refunds, err
}

// Get individual refund
func (s *RefundServiceOp) Get(orderID int64, refundID int64, options interface{}) (*Refund, error) {
	path := fmt.Sprintf("%s/%d/refunds/%d.json", ordersBasePath, orderID, refundID)
	resource := new(RefundResource)
	err := s.client.Get(path, resource, options)
	return resource.Refund, err
}

// Caculate a new refund
func (s *RefundServiceOp) Caculate(orderID int64, refund Refund) (*Refund, error) {
	path := fmt.Sprintf("%s/%d/refunds/caculate.json", ordersBasePath, orderID)
	wrappedData := RefundResource{Refund: &refund}
	resource := new(RefundResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Refund, err
}

// Create a new refund
func (s *RefundServiceOp) Create(orderID int64, refund Refund) (*Refund, error) {
	path := fmt.Sprintf("%s/%d/refunds.json", ordersBasePath, orderID)
	wrappedData := RefundResource{Refund: &refund}
	resource := new(RefundResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Refund, err
}
