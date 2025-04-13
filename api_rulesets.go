package gocloudflare

import (
	"context"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

const (
	// WAF Managed Rules - https://developers.cloudflare.com/waf/managed-rules/
	RulesetIDCloudflareManaged                 = "efb7b8c949ac4650a09736fc376e9aee"
	RulesetIDCloudflareOWASPCore               = "4814384a9e5d4991b9815dcfc25d2f1f"
	RulesetIDCloudflareExposedCredentialsCheck = "c2e184081120413c86c3ab7e14069605"
	RulesetIDCloudflareFreeManaged             = "77454fe2d30c4220b5701f6fdfb893ba"
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

/*
	RulesetPhaseHTTPRequestFirewallCustom    RulesetPhase = "http_request_firewall_custom"
	RulesetPhaseHTTPRequestFirewallManaged   RulesetPhase = "http_request_firewall_managed"
	RulesetPhaseHTTPRequestLateTransform     RulesetPhase = "http_request_late_transform"
*/
