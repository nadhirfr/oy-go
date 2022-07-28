package accountinquiry

import (
	"context"

	oygo "github.com/nadhirfr/oy-go"
)

// Get gets one invoice
func Get(data *GetParams) (*oygo.AccountInquiryGet, *oygo.Error) {
	return GetWithContext(context.Background(), data)
}

// GetWithContext gets one invoice with context
func GetWithContext(ctx context.Context, data *GetParams) (*oygo.AccountInquiryGet, *oygo.Error) {
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
