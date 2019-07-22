package zeit

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

// ListSecrets return a list of  secrets
func (s *Service) ListSecrets(name, value string) (interface{}, error) {
	var List ListSecrets

	request, err := http.NewRequest("GET", SecretsURL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", s.Token)

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

// AddSecrets return a list of  secrets
func (s *Service) AddSecrets(name, value string) (interface{}, error) {
	var (
		Secret AddSecrets
		r      http.Request
	)

	r.ParseForm()
	r.Form.Add("name", name)
	r.Form.Add("value", value)
	bodystr := strings.TrimSpace(r.Form.Encode())

	request, err := http.NewRequest("POST", SecretsURL, strings.NewReader(bodystr))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", s.Token)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &Secret)
	if err != nil {
		return nil, err
	}

	return Secret, nil
}

// UpdateName edit the name of a user's secret
func (s *Service) UpdateName(name string) (interface{}, error) {
	var Name UpdateName

	request, err := http.NewRequest("PATCH", SecretsURL+"/"+name, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", s.Token)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &Name)
	if err != nil {
		return nil, err
	}

	return Name, nil
}

// DeleteSecret  deletes the user's secret defined in the URL.
func (s *Service) DeleteSecret(name string) (interface{}, error) {
	var Secret DeleteSecret

	request, err := http.NewRequest("DELETE", SecretsURL+"/"+name, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", s.Token)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &Secret)
	if err != nil {
		return nil, err
	}

	return Secret, nil
}
