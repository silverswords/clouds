package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RepoTrending -
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

// DevelTrending -
func DevelTrending(language string, datarange string) (interface{}, error) {
	var DevelTrendModel DevelTrend
	url := fmt.Sprintf(GitHubDevelTrendURL, language, datarange)

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

	err = json.Unmarshal(body, &DevelTrendModel)
	if err != nil {
		return nil, err
	}

	return DevelTrendModel, nil
}
