package scheduleddisbursement

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
func (c *Client) Create(data *CreateParams) (*oygo.ScheduledDisbursementResult, *oygo.Error) {
	return c.CreateWithContext(context.Background(), data)
}

// CreateWithContext creates new invoice with context
func (c *Client) CreateWithContext(ctx context.Context, data *CreateParams) (*oygo.ScheduledDisbursementResult, *oygo.Error) {

	response := &oygo.ScheduledDisbursementResult{}
	header := &http.Header{
		"X-Api-Key":     []string{c.Opt.SecretKey},
		"X-Oy-Username": []string{c.Opt.Username},
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/api/scheduled-remit", c.Opt.OyURL),
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
func (c *Client) Get(data *GetParams) (*oygo.ScheduledDisbursementGet, *oygo.Error) {
	return c.GetWithContext(context.Background(), data)
}

// GetWithContext gets one invoice with context
func (c *Client) GetWithContext(ctx context.Context, data *GetParams) (*oygo.ScheduledDisbursementGet, *oygo.Error) {
	response := &oygo.ScheduledDisbursementGet{}
	header := &http.Header{
		"X-Api-Key":     []string{c.Opt.SecretKey},
		"X-Oy-Username": []string{c.Opt.Username},
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/api/scheduled-remit", c.Opt.OyURL),
		header,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
