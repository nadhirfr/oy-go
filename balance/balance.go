package balance

import (
	"context"

	oygo "github.com/nadhirfr/oy-go"
)

// Get balance
func Get() (*oygo.BalanceGet, *oygo.Error) {
	return GetWithContext(context.Background())
}

// GetWithContext get balance with context
func GetWithContext(ctx context.Context) (*oygo.BalanceGet, *oygo.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetWithContext(ctx)
}

func getClient() (*Client, *oygo.Error) {
	return &Client{
		Opt:          &oygo.Opt,
		APIRequester: oygo.GetAPIRequester(),
	}, nil
}
