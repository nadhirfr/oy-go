package checkout

import (
	"context"

	oygo "github.com/nadhirfr/oy-go"
)

// Create creates new invoice
func Create(data *CreateParams) (*oygo.CheckoutResult, *oygo.Error) {
	return CreateWithContext(context.Background(), data)
}

// CreateWithContext creates new invoice with context
func CreateWithContext(ctx context.Context, data *CreateParams) (*oygo.CheckoutResult, *oygo.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateWithContext(ctx, data)
}

// Get gets one invoice
func Get(data *GetParams) (*oygo.CheckoutGet, *oygo.Error) {
	return GetWithContext(context.Background(), data)
}

// GetWithContext gets one invoice with context
func GetWithContext(ctx context.Context, data *GetParams) (*oygo.CheckoutGet, *oygo.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetWithContext(ctx, data)
}

func getClient() (*Client, *oygo.Error) {
	return &Client{
		Opt:          &oygo.Opt,
		APIRequester: oygo.GetAPIRequester(),
	}, nil
}
