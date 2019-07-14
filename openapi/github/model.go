package github

var (
	// GitHubRepoTrendURL GitHub reposisories trending API
	GitHubRepoTrendURL = "https://github-trending-api.now.sh/repositories?language=%s&since=%s"
	// GitHubDeveloperTrendURL GitHub developer trending API
	GitHubDeveloperTrendURL = "https://github-trending-api.now.sh/developers?language=%s&since=%s"
)

// RepoTrend GitHub repositories trending JSON structure
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

// DeveloperTrend GitHub developers trending JSON structure
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
