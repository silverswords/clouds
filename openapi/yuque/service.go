package yuque

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

// List return a list of repositories in the group
func (s *Service) List(repoid string) (interface{}, error) {
	var List BookDetail

	url := fmt.Sprintf(ListURL, repoid)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("X-Auth-Token", s.Token)

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

	err = json.Unmarshal(body, &List)
	if err != nil {
		return nil, err
	}

	return List, nil
}

// Details rerurn a yuque document of detail contents
func (s *Service) Details(repoid, id string) (interface{}, error) {
	var Details DocDetail

	url := fmt.Sprintf(DetailURL, repoid, id)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("X-Auth-Token", s.Token)

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

	err = json.Unmarshal(body, &Details)
	if err != nil {
		return nil, err
	}

	return Details, nil
}

// Repo return a list of repositories in the group
func (s *Service) Repo(groupid string) (interface{}, error) {
	var Repo Book

	url := fmt.Sprintf(RepoURL, groupid)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("X-Auth-Token", s.Token)

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

	err = json.Unmarshal(body, &Repo)
	if err != nil {
		return nil, err
	}

	return Repo, nil
}
