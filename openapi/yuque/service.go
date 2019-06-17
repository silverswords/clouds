package yuque

import (
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
func (s *Service) List(repoid string) (string, error) {
	url := fmt.Sprintf(ListURL, repoid)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("X-Auth-Token", s.Token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Details -
func (s *Service) Details(repoid, id string) (string, error) {
	url := fmt.Sprintf(ListURL, repoid, id)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("X-Auth-Token", s.Token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Repo -
func (s *Service) Repo(groupid string) (string, error) {
	url := fmt.Sprintf(ListURL, groupid)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("X-Auth-Token", s.Token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
