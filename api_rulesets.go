package gocloudflare

import (
	"context"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

type RulesetsAPI struct {
	*cloudflare.API
}

func (api *RulesetsAPI) CreateRuleset(ctx context.Context, rc *cloudflare.ResourceContainer, params cloudflare.CreateRulesetParams) (cloudflare.Ruleset, error) {
	return api.API.CreateRuleset(ctx, rc, params)
}

func (api *RulesetsAPI) DeleteRuleset(ctx context.Context, rc *cloudflare.ResourceContainer, rulesetID string) error {
	return api.API.DeleteRuleset(ctx, rc, rulesetID)
}

func (api *RulesetsAPI) GetRuleset(ctx context.Context, rc *cloudflare.ResourceContainer, rulesetID string) (cloudflare.Ruleset, error) {
	return api.API.GetRuleset(ctx, rc, rulesetID)
}

func (api *RulesetsAPI) ListRulesets(ctx context.Context, rc *cloudflare.ResourceContainer, params cloudflare.ListRulesetsParams) ([]cloudflare.Ruleset, error) {
	return api.API.ListRulesets(ctx, rc, params)
}

func (api *RulesetsAPI) UpdateRuleset(ctx context.Context, rc *cloudflare.ResourceContainer, params cloudflare.UpdateRulesetParams) (cloudflare.Ruleset, error) {
	return api.API.UpdateRuleset(ctx, rc, params)
}
