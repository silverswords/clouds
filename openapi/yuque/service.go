package yuque

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Service -
type Service struct {
	Token string
}

// NewService -
func NewService(token string) *Service {
	return &Service{
		Token: token,
	}
}

// List -
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

// Details -
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

// Repo -
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
