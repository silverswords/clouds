package github

import (
	"net/http"

	gogithub "github.com/google/go-github/v27/github"
)

// APIClient encapsulate github.Client
type APIClient struct {
	Client *gogithub.Client
}

// NewAPIClient creates GitHubClient
func NewAPIClient(g *http.Client) *APIClient {
	client := gogithub.NewClient(g)
	return &APIClient{
		Client: client,
	}
}
