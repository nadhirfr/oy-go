package checkout

import (
	"context"
	"fmt"
	"net/http"

	oygo "github.com/nadhirfr/oy-go"
)

// Client is the client used to invoke invoice API.
type Client struct {
	Opt          *oygo.Option
	APIRequester oygo.APIRequester
}

// Create creates new invoice
func (c *Client) Create(data *CreateParams) (*oygo.CheckoutResult, *oygo.Error) {
	return c.CreateWithContext(context.Background(), data)
}

// CreateWithContext creates new invoice with context
func (c *Client) CreateWithContext(ctx context.Context, data *CreateParams) (*oygo.CheckoutResult, *oygo.Error) {

	response := &oygo.CheckoutResult{}
	header := &http.Header{
		"X-Api-Key":     []string{c.Opt.SecretKey},
		"X-Oy-Username": []string{c.Opt.Username},
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/api/payment-checkout/create-v2", c.Opt.OyURL),
		header,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get gets one invoice
func (c *Client) Get(data *GetParams) (*oygo.CheckoutGet, *oygo.Error) {
	return c.GetWithContext(context.Background(), data)
}

// GetWithContext gets one invoice with context
func (c *Client) GetWithContext(ctx context.Context, data *GetParams) (*oygo.CheckoutGet, *oygo.Error) {
	response := &oygo.CheckoutGet{}
	header := &http.Header{
		"X-Api-Key":     []string{c.Opt.SecretKey},
		"X-Oy-Username": []string{c.Opt.Username},
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/api/payment-checkout/status?partner_tx_id=%s", c.Opt.OyURL, data.PartnerTxId),
		header,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
