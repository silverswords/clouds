package zeit

import "time"

var (
	// SecretsURL API about zeit  secrets
	SecretsURL = "https://api.zeit.co/v2/now/secrets"
)

// ListSecrets JSON structure for a list of  zeit secrets
type ListSecrets struct {
	Secrets []struct {
		UID     string    `json:"uid"`
		Name    string    `json:"name"`
		Created time.Time `json:"created"`
	} `json:"secrets"`
}

// AddSecrets JSON structure for creating  a zeit secret
type AddSecrets struct {
	UID     string    `json:"uid"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	UserID  string    `json:"userId"`
	Value   struct {
		Type string `json:"type"`
		Data []int  `json:"data"`
	} `json:"value"`
}

// UpdateName JSON structure for edit the name of a user's secret
type UpdateName struct {
	UID     string    `json:"uid"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	OldName string    `json:"oldName"`
}

// DeleteSecret JSON structure for delete the name of a user's secret
type DeleteSecret struct {
	UID     string    `json:"uid"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}
