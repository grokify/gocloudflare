package gocloudflare

import (
	cloudflare "github.com/cloudflare/cloudflare-go"
)

type API struct {
	*cloudflare.API
	Rulesets *RulesetsAPI
}

func NewAPI(api *cloudflare.API) *API {
	return &API{
		API:      api,
		Rulesets: &RulesetsAPI{API: api},
	}
}
