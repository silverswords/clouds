package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
