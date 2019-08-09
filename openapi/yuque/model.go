package yuque

import "time"

var (
	// UserURL API about user information
	UserURL = "https://www.yuque.com/api/v2/user"
	// GroupURL API about user's groups
	GroupURL = "https://www.yuque.com/api/v2/users/%s/groups"
	// UserRepoURL API about user's repos
	UserRepoURL = "https://www.yuque.com/api/v2/users/%s/repos"
	// DetailURL API about document detail contents
	DetailURL = "https://www.yuque.com/api/v2/repos/%s/docs/%s?raw=0"
	// ListURL API about a list of articles in the library
	ListURL = "https://www.yuque.com/api/v2/repos/%s/docs/"
	// RepoURL API about a list of repositories in the group
	RepoURL = "https://www.yuque.com/api/v2/groups/%s/repos/"
)

//Info JSON structure for the authorized user
type Info struct {
	Data struct {
		ID               int       `json:"id"`
		Type             string    `json:"type"`
		SpaceID          int       `json:"space_id"`
		AccountID        int       `json:"account_id"`
		Login            string    `json:"login"`
		Name             string    `json:"name"`
		AvatarURL        string    `json:"avatar_url"`
		LargeAvatarURL   string    `json:"large_avatar_url"`
		MediumAvatarURL  string    `json:"medium_avatar_url"`
		SmallAvatarURL   string    `json:"small_avatar_url"`
		BooksCount       int       `json:"books_count"`
		PublicBooksCount int       `json:"public_books_count"`
		FollowersCount   int       `json:"followers_count"`
		FollowingCount   int       `json:"following_count"`
		Public           int       `json:"public"`
		Description      string    `json:"description"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		Serializer       string    `json:"_serializer"`
	} `json:"data"`
}

// DocDetail JSON structure for a yuque document detail contents
type DocDetail struct {
	Abilities struct {
		Update  bool `json:"update"`
		Destroy bool `json:"destroy"`
	} `json:"abilities"`
	Data struct {
		ID     int    `json:"id"`
		Slug   string `json:"slug"`
		Title  string `json:"title"`
		BookID int    `json:"book_id"`
		Book   struct {
			ID               int       `json:"id"`
			Type             string    `json:"type"`
			Slug             string    `json:"slug"`
			Name             string    `json:"name"`
			UserID           int       `json:"user_id"`
			Description      string    `json:"description"`
			CreatorID        int       `json:"creator_id"`
			Public           int       `json:"public"`
			ItemsCount       int       `json:"items_count"`
			LikesCount       int       `json:"likes_count"`
			WatchesCount     int       `json:"watches_count"`
			ContentUpdatedAt time.Time `json:"content_updated_at"`
			UpdatedAt        time.Time `json:"updated_at"`
			CreatedAt        time.Time `json:"created_at"`
			Namespace        string    `json:"namespace"`
			User             struct {
				ID               int         `json:"id"`
				Type             string      `json:"type"`
				Login            string      `json:"login"`
				Name             string      `json:"name"`
				Description      interface{} `json:"description"`
				AvatarURL        string      `json:"avatar_url"`
				LargeAvatarURL   string      `json:"large_avatar_url"`
				MediumAvatarURL  string      `json:"medium_avatar_url"`
				SmallAvatarURL   string      `json:"small_avatar_url"`
				BooksCount       int         `json:"books_count"`
				PublicBooksCount int         `json:"public_books_count"`
				FollowersCount   int         `json:"followers_count"`
				FollowingCount   int         `json:"following_count"`
				CreatedAt        time.Time   `json:"created_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				Serializer       string      `json:"_serializer"`
			} `json:"user"`
			Serializer string `json:"_serializer"`
		} `json:"book"`
		UserID  int `json:"user_id"`
		Creator struct {
			ID               int       `json:"id"`
			Type             string    `json:"type"`
			Login            string    `json:"login"`
			Name             string    `json:"name"`
			Description      string    `json:"description"`
			AvatarURL        string    `json:"avatar_url"`
			LargeAvatarURL   string    `json:"large_avatar_url"`
			MediumAvatarURL  string    `json:"medium_avatar_url"`
			SmallAvatarURL   string    `json:"small_avatar_url"`
			BooksCount       int       `json:"books_count"`
			PublicBooksCount int       `json:"public_books_count"`
			FollowersCount   int       `json:"followers_count"`
			FollowingCount   int       `json:"following_count"`
			CreatedAt        time.Time `json:"created_at"`
			UpdatedAt        time.Time `json:"updated_at"`
			Serializer       string    `json:"_serializer"`
		} `json:"creator"`
		Format            string      `json:"format"`
		Body              string      `json:"body"`
		BodyDraft         string      `json:"body_draft"`
		BodyHTML          string      `json:"body_html"`
		BodyLake          string      `json:"body_lake"`
		Public            int         `json:"public"`
		Status            int         `json:"status"`
		LikesCount        int         `json:"likes_count"`
		CommentsCount     int         `json:"comments_count"`
		ContentUpdatedAt  time.Time   `json:"content_updated_at"`
		DeletedAt         interface{} `json:"deleted_at"`
		CreatedAt         time.Time   `json:"created_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
		PublishedAt       time.Time   `json:"published_at"`
		FirstPublishedAt  time.Time   `json:"first_published_at"`
		WordCount         int         `json:"word_count"`
		Cover             interface{} `json:"cover"`
		Description       string      `json:"description"`
		CustomDescription interface{} `json:"custom_description"`
		Serializer        string      `json:"_serializer"`
	} `json:"data"`
}

// BookDetail  JSON structure for a list of articles in the library
type BookDetail struct {
	Data []struct {
		ID                int         `json:"id"`
		Slug              string      `json:"slug"`
		Title             string      `json:"title"`
		Description       string      `json:"description"`
		UserID            int         `json:"user_id"`
		BookID            int         `json:"book_id"`
		Format            string      `json:"format"`
		Public            int         `json:"public"`
		Status            int         `json:"status"`
		LikesCount        int         `json:"likes_count"`
		CommentsCount     int         `json:"comments_count"`
		ContentUpdatedAt  time.Time   `json:"content_updated_at"`
		CreatedAt         time.Time   `json:"created_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
		PublishedAt       time.Time   `json:"published_at"`
		FirstPublishedAt  time.Time   `json:"first_published_at"`
		DraftVersion      int         `json:"draft_version"`
		LastEditorID      int         `json:"last_editor_id"`
		WordCount         int         `json:"word_count"`
		Cover             interface{} `json:"cover"`
		CustomDescription interface{} `json:"custom_description"`
		LastEditor        struct {
			ID              int       `json:"id"`
			Type            string    `json:"type"`
			Login           string    `json:"login"`
			Name            string    `json:"name"`
			Description     string    `json:"description"`
			AvatarURL       string    `json:"avatar_url"`
			LargeAvatarURL  string    `json:"large_avatar_url"`
			MediumAvatarURL string    `json:"medium_avatar_url"`
			SmallAvatarURL  string    `json:"small_avatar_url"`
			FollowersCount  int       `json:"followers_count"`
			FollowingCount  int       `json:"following_count"`
			CreatedAt       time.Time `json:"created_at"`
			UpdatedAt       time.Time `json:"updated_at"`
			Serializer      string    `json:"_serializer"`
		} `json:"last_editor"`
		Book       interface{} `json:"book"`
		Serializer string      `json:"_serializer"`
	} `json:"data"`
}

// Book JSON structure for a list of repositories in the group
type Book struct {
	Data []struct {
		ID               int       `json:"id"`
		Type             string    `json:"type"`
		Slug             string    `json:"slug"`
		Name             string    `json:"name"`
		UserID           int       `json:"user_id"`
		Description      string    `json:"description"`
		CreatorID        int       `json:"creator_id"`
		Public           int       `json:"public"`
		ItemsCount       int       `json:"items_count"`
		LikesCount       int       `json:"likes_count"`
		WatchesCount     int       `json:"watches_count"`
		ContentUpdatedAt time.Time `json:"content_updated_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		CreatedAt        time.Time `json:"created_at"`
		Namespace        string    `json:"namespace"`
		User             struct {
			ID               int       `json:"id"`
			Type             string    `json:"type"`
			Login            string    `json:"login"`
			Name             string    `json:"name"`
			Description      string    `json:"description"`
			AvatarURL        string    `json:"avatar_url"`
			LargeAvatarURL   string    `json:"large_avatar_url"`
			MediumAvatarURL  string    `json:"medium_avatar_url"`
			SmallAvatarURL   string    `json:"small_avatar_url"`
			BooksCount       int       `json:"books_count"`
			PublicBooksCount int       `json:"public_books_count"`
			FollowersCount   int       `json:"followers_count"`
			FollowingCount   int       `json:"following_count"`
			CreatedAt        time.Time `json:"created_at"`
			UpdatedAt        time.Time `json:"updated_at"`
			Serializer       string    `json:"_serializer"`
		} `json:"user"`
		Serializer string `json:"_serializer"`
	} `json:"data"`
}

//Groups JSON structure for  user's groups
type Groups struct {
	Data []struct {
		ID                int       `json:"id"`
		Login             string    `json:"login"`
		Name              string    `json:"name"`
		AvatarURL         string    `json:"avatar_url"`
		LargeAvatarURL    string    `json:"large_avatar_url"`
		MediumAvatarURL   string    `json:"medium_avatar_url"`
		SmallAvatarURL    string    `json:"small_avatar_url"`
		BooksCount        int       `json:"books_count"`
		PublicBooksCount  int       `json:"public_books_count"`
		TopicsCount       int       `json:"topics_count"`
		PublicTopicsCount int       `json:"public_topics_count"`
		MembersCount      int       `json:"members_count"`
		Public            int       `json:"public"`
		Description       string    `json:"description"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Serializer        string    `json:"_serializer"`
	} `json:"data"`
}

// UserRepos JSON structure for  user's repos
type UserRepos struct {
	Data []struct {
		ID               int       `json:"id"`
		Type             string    `json:"type"`
		Slug             string    `json:"slug"`
		Name             string    `json:"name"`
		UserID           int       `json:"user_id"`
		Description      string    `json:"description"`
		CreatorID        int       `json:"creator_id"`
		Public           int       `json:"public"`
		ItemsCount       int       `json:"items_count"`
		LikesCount       int       `json:"likes_count"`
		WatchesCount     int       `json:"watches_count"`
		ContentUpdatedAt time.Time `json:"content_updated_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		CreatedAt        time.Time `json:"created_at"`
		Namespace        string    `json:"namespace"`
		User             struct {
			ID               int       `json:"id"`
			Type             string    `json:"type"`
			Login            string    `json:"login"`
			Name             string    `json:"name"`
			Description      string    `json:"description"`
			AvatarURL        string    `json:"avatar_url"`
			LargeAvatarURL   string    `json:"large_avatar_url"`
			MediumAvatarURL  string    `json:"medium_avatar_url"`
			SmallAvatarURL   string    `json:"small_avatar_url"`
			BooksCount       int       `json:"books_count"`
			PublicBooksCount int       `json:"public_books_count"`
			FollowersCount   int       `json:"followers_count"`
			FollowingCount   int       `json:"following_count"`
			CreatedAt        time.Time `json:"created_at"`
			UpdatedAt        time.Time `json:"updated_at"`
			Serializer       string    `json:"_serializer"`
		} `json:"user"`
		Serializer string `json:"_serializer"`
	} `json:"data"`
}
