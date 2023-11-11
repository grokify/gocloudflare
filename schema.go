package gocloudflare

type Resource struct {
	Single    string
	Plural    string
	Functions []string
}

func Resources() []Resource {
	return []Resource{
		{
			Single:    "Ruleset",
			Plural:    "Rulesets",
			Functions: []string{"Create", "Get", "List", "Update"},
		},
	}
}
