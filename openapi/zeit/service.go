package zeit

import (
	"encoding/json"
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

// ListSecrets return a list of  secrets
func (s *Service) ListSecrets(name, value string) (interface{}, error) {
	var List ListSecrets

	request, err := http.NewRequest("GET", ListSecretsURL, nil)
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
