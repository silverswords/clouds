package github

var (
	// GitHubRepoTrendURL GitHub reposisories trending API
	GitHubRepoTrendURL = "https://github-trending-api.now.sh/repositories?language=%s&since=%s"
	// GitHubDeveloperTrendURL GitHub developer trending API
	GitHubDeveloperTrendURL = "https://github-trending-api.now.sh/developers?language=%s&since=%s"
	// GitHubRepoContributorURL GitHub reposisories contributor API
	GitHubRepoContributorURL = "https://api.github.com/repos/%s/%s/stats/contributors"
)

// RepoTrend JSON structure for GitHub repositories trending
type RepoTrend []struct {
	Author             string `json:"author"`
	Name               string `json:"name"`
	Avatar             string `json:"avatar"`
	URL                string `json:"url"`
	Description        string `json:"description"`
	Language           string `json:"language"`
	LanguageColor      string `json:"languageColor"`
	Stars              int    `json:"stars"`
	Forks              int    `json:"forks"`
	CurrentPeriodStars int    `json:"currentPeriodStars"`
	BuiltBy            []struct {
		Username string `json:"username"`
		Href     string `json:"href"`
		Avatar   string `json:"avatar"`
	} `json:"builtBy"`
}

// DeveloperTrend JSON structure for GitHub developers trending JSON structure
type DeveloperTrend []struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	Avatar   string `json:"avatar"`
	Repo     struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"repo"`
}

// RepoContributor JSON structure for getting contributors list with additions, deletions, and commit counts
type RepoContributor []struct {
	Total int `json:"total"`
	Weeks []struct {
		W int `json:"w"`
		A int `json:"a"`
		D int `json:"d"`
		C int `json:"c"`
	} `json:"weeks"`
	Author struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
}
