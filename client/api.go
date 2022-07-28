package client

import (
	"net/http"

	oygo "github.com/nadhirfr/oy-go"
	"github.com/nadhirfr/oy-go/accountinquiry"
	"github.com/nadhirfr/oy-go/balance"
	"github.com/nadhirfr/oy-go/checkout"
	"github.com/nadhirfr/oy-go/disbursement"
	"github.com/nadhirfr/oy-go/scheduleddisbursement"
)

// API is the oy client which contains all products
type API struct {
	opt                   oygo.Option
	apiRequester          oygo.APIRequester
	Checkout              *checkout.Client
	Balance               *balance.Client
	Disbursement          *disbursement.Client
	ScheduledDisbursement *scheduleddisbursement.Client
	AccountInquiry        *accountinquiry.Client
}

func (a *API) init() {
	a.Checkout = &checkout.Client{Opt: &a.opt, APIRequester: a.apiRequester}
	a.Balance = &balance.Client{Opt: &a.opt, APIRequester: a.apiRequester}
	a.Disbursement = &disbursement.Client{Opt: &a.opt, APIRequester: a.apiRequester}
	a.ScheduledDisbursement = &scheduleddisbursement.Client{Opt: &a.opt, APIRequester: a.apiRequester}
	a.AccountInquiry = &accountinquiry.Client{Opt: &a.opt, APIRequester: a.apiRequester}
}

// New creates a new oy API client
func New(secretKey, username string) *API {
	api := API{
		opt: oygo.Option{
			SecretKey: secretKey,
			OyURL:     "https://partner.oyindonesia.com",
			// OyURL:     "https://api-stg.oyindonesia.com",
			Username: username,
		},
		apiRequester: oygo.GetAPIRequester(),
	}
	api.init()

	return &api
}

// WithAPIRequester set custom APIRequester for Oy Client
// Can be chained with constructor like below:
// 		client.New(yourSecretKey).WithAPIRequester(yourCustomRequester)
func (a *API) WithAPIRequester(apiRequester oygo.APIRequester) *API {
	a.apiRequester = apiRequester
	a.init()
	return a
}

// WithCustomURL set custom Oy URL for Oy Client
// Can be chained with constructor like below:
// 		client.New(yourSecretKey).WithCustomURL(yourCustomURL)
func (a *API) WithCustomURL(oyURL string) *API {
	a.opt.OyURL = oyURL
	a.init()
	return a
}

// WithCustomHTTPClient set custom HTTP Client for default API Requester
// Can be chained with constructor like below:
// 		client.New(yourSecretKey).WithCustomHTTPClient(yourCustomHTTPClient)
func (a *API) WithCustomHTTPClient(client *http.Client) *API {
	a.apiRequester = &oygo.APIRequesterImplementation{
		HTTPClient: client,
	}
	a.init()
	return a
}
