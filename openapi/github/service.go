package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	gogithub "github.com/google/go-github/github"
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

// Service encapsulate authenticated token
type Service struct {
	Token string
}

// NewService create Service for external use
func NewService(token string) *Service {
	return &Service{
		Token: token,
	}
}

// Contributor return contributors list with additions, deletions, and commit counts
func (s *Service) Contributor(Owner string, Repo string) (interface{}, error) {
	var Contributors RepoContributor
	url := fmt.Sprintf(GitHubRepoContributorURL, Owner, Repo)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "token "+s.Token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &Contributors)
	if err != nil {
		return nil, err
	}

	return Contributors, nil
}

// RepoTrending return an array of trending repositories.
func RepoTrending(language string, datarange string) (interface{}, error) {
	var RepoTrendModel RepoTrend
	url := fmt.Sprintf(GitHubRepoTrendURL, language, datarange)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &RepoTrendModel)
	if err != nil {
		return nil, err
	}

	return RepoTrendModel, nil
}

// DeveloperTrending return an array of trending developers.
func DeveloperTrending(language string, datarange string) (interface{}, error) {
	var DeveloperTrendModel DeveloperTrend
	url := fmt.Sprintf(GitHubDeveloperTrendURL, language, datarange)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &DeveloperTrendModel)
	if err != nil {
		return nil, err
	}

	return DeveloperTrendModel, nil
}
