package zeit

import "time"

var (
	// ListSecretsURL API about a list of secrets
	ListSecretsURL = "https://api.zeit.co/v2/now/secrets"
)

// ListSecrets JSON structure for a list of  zeit secret
type ListSecrets struct {
	Secrets []struct {
		UID     string    `json:"uid"`
		Name    string    `json:"name"`
		Created time.Time `json:"created"`
	} `json:"secrets"`
}
