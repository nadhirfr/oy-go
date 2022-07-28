package accountinquiry

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

// Get gets one invoice
func (c *Client) Get(data *GetParams) (*oygo.AccountInquiryGet, *oygo.Error) {
	return c.GetWithContext(context.Background(), data)
}

// GetWithContext gets one invoice with context
func (c *Client) GetWithContext(ctx context.Context, data *GetParams) (*oygo.AccountInquiryGet, *oygo.Error) {
	response := &oygo.AccountInquiryGet{}
	header := &http.Header{
		"X-Api-Key":     []string{c.Opt.SecretKey},
		"X-Oy-Username": []string{c.Opt.Username},
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/api/account-inquiry", c.Opt.OyURL),
		header,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
