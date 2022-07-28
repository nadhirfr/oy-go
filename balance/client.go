package balance

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

// Get get balance
func (c *Client) Get() (*oygo.BalanceGet, *oygo.Error) {
	return c.GetWithContext(context.Background())
}

// GetWithContext gets one invoice with context
func (c *Client) GetWithContext(ctx context.Context) (*oygo.BalanceGet, *oygo.Error) {
	response := &oygo.BalanceGet{}
	header := &http.Header{
		"X-Api-Key":     []string{c.Opt.SecretKey},
		"X-Oy-Username": []string{c.Opt.Username},
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/api/balance", c.Opt.OyURL),
		header,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
