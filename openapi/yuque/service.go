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

//UserInfo return information for the authorized user
func (s *Service) UserInfo() (interface{}, error) {
	var User Info

	request, err := http.NewRequest("GET", UserURL, nil)
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

	err = json.Unmarshal(body, &User)
	if err != nil {
		return nil, err
	}

	return User, nil
}

// List return a list of repositories in the group
func (s *Service) List(Repoid string) (interface{}, error) {
	var List BookDetail

	url := fmt.Sprintf(ListURL, Repoid)
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
func (s *Service) Details(RepoID, ID string) (interface{}, error) {
	var Details DocDetail

	url := fmt.Sprintf(DetailURL, RepoID, ID)
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
func (s *Service) Repo(GroupID string) (interface{}, error) {
	var Repo Book

	url := fmt.Sprintf(RepoURL, GroupID)
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

//UserGroups return information for the user's groups
func (s *Service) UserGroups(UserID string) (interface{}, error) {
	var Group Groups

	url := fmt.Sprintf(GroupURL, UserID)
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

	err = json.Unmarshal(body, &Group)
	if err != nil {
		return nil, err
	}

	return Group, nil
}

//UserRepos return information for the user's groups
func (s *Service) UserRepos(UserID string) (interface{}, error) {
	var Repos UserRepos

	url := fmt.Sprintf(UserRepoURL, UserID)
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

	err = json.Unmarshal(body, &Repos)
	if err != nil {
		return nil, err
	}

	return Repos, nil
}
