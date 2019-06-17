package yuque

var (
	// DetailURL -
	DetailURL = "https://www.yuque.com/api/v2/repos/%s/docs/%s?raw=0"
	// ListURL -
	ListURL = "https://www.yuque.com/api/v2/repos/%s/docs/"
	// RepoURL -
	RepoURL = "https://www.yuque.com/api/v2/groups/%s/repos/"
)

// DetailBody -
var DetailBody struct {
	Abilities Abilities  `json:"abilities"`
	Data      DetailData `json:"detaildata"`
}

// Abilities -
type Abilities struct {
	Update  bool `json:"update"`
	Destroy bool `json:"destroy"`
}

// DetailData -
type DetailData struct {
	ID     int64  `json:"id"`
	Slug   string `json:"slug"`
	Title  string `json:"title"`
	BookID int64  `json:"book_id"`
	Book   struct {
		ID               int64  `json:"id"`
		Type             string `json:"type"`
		Slug             string `json:"slug"`
		Name             string `json:"name"`
		UserID           int64  `json:"user_id"`
		Description      string `json:"description"`
		CreatorID        int64  `json:"creator_id"`
		Public           int    `json:"public"`
		ItemsCount       int    `json:"items_count"`
		LikesCount       int    `json:"likes_count"`
		WatchesCount     int    `json:"watches_count"`
		ContentUpdatedAt string `json:"content_updated_at"`
		UpdatedAt        string `json:"updated_at"`
		CreatedAt        string `json:"created_at"`
		Namespace        string `json:"namespace"`
		User             struct {
			ID               int64  `json:"id"`
			Type             string `json:"type"`
			Login            string `json:"login"`
			Name             string `json:"name"`
			Description      string `json:"description"`
			AvatarURL        string `json:"avatar_url"`
			LargeAvatarURL   string `json:"large_avatar_url"`
			MediumAvatarURL  string `json:"medium_avatar_url"`
			SmallAvatarURL   string `json:"small_avatar_url"`
			BooksCount       int    `json:"books_count"`
			PublicBooksCount int    `json:"public_books_count"`
			FollowersCount   int    `json:"followers_count"`
			FollowingCount   int    `json:"following_count"`
			CreatedAt        string `json:"created_at"`
			UpdatedAt        string `json:"updated_at"`
			Serializer       string `json:"_serializer"`
		} `json:"user"`
		Serializer string `json:"_serializer"`
	} `json:"book"`
	UserID  int64 `json:"user_id"`
	Creator struct {
		ID               int64  `json:"id"`
		Type             string `json:"type"`
		Login            string `json:"login"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		AvatarURL        string `json:"avatar_url"`
		LargeAvatarURL   string `json:"large_avatar_url"`
		MediumAvatarURL  string `json:"medium_avatar_url"`
		SmallAvatarURL   string `json:"small_avatar_url"`
		BooksCount       int    `json:"books_count"`
		PublicBooksCount int    `json:"public_books_count"`
		FollowersCount   int    `json:"followers_count"`
		FollowingCount   int    `json:"following_count"`
		CreatedAt        string `json:"created_at"`
		UpdatedAt        string `json:"updated_at"`
		Serializer       string `json:"_serializer"`
	} `json:"creator"`
	Format            string `json:"format"`
	Body              string `json:"body"`
	BodyDraft         string `json:"body_draft"`
	BodyHTML          string `json:"body_html"`
	BodyLake          string `json:"body_lake"`
	Public            int    `json:"public"`
	Status            int    `json:"status"`
	LikesCount        int    `json:"likes_count"`
	CommentsCount     int    `json:"comments_count"`
	ContentUpdatedAt  string `json:"content_updated_at"`
	DeletedAt         string `json:"deleted_at"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	PublishedAt       string `json:"published_at"`
	FirstPublishedAt  string `json:"first_published_at"`
	WordCount         int    `json:"word_count"`
	Cover             string `json:"cover"`
	Description       string `json:"description"`
	CustomDescription string `json:"custom_description"`
	Serializer        string `json:"_serializer"`
}

// ListBody -
var ListBody struct {
	Data []ListDate `json:"data"`
}

// ListDate -
type ListDate struct {
	ID               int64  `json:"id"`
	Slug             string `json:"slug"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	UserID           int64  `json:"user_id"`
	BookID           int64  `json:"book_id"`
	Format           string `json:"format"`
	Public           int    `json:"public"`
	Status           int    `json:"status"`
	LikesCount       int    `json:"likes_count"`
	CommentsCount    int    `json:"comments_count"`
	ContentUpdated   string `json:"content_updated_at"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	PublishedAt      string `json:"published_at"`
	FirstPubilshedAt string `json:"first_published_at"`
	DraftVersion     int    `json:"draft_version"`
	LastEditorID     int32  `json:"last_editor_id"`
	WordCount        int    `json:"word_count"`
	Cover            string `json:"cover"`
	Custom           string `json:"custom_description"`
	LastEditor       struct {
		ID              int64  `json:"id"`
		Type            string `json:"type"`
		Login           string `json:"login"`
		Name            string `json:"name"`
		Description     string `json:"description"`
		AvatarURL       string `json:"avatar_url"`
		LargeAvatarURL  string `json:"large_avatar_url"`
		MediumAvatarURL string `json:"medium_avatar_url"`
		SmallAvatarURL  string `json:"small_avatar_url"`
		FollowersCount  int    `json:"followers_count"`
		FollowingCount  int    `json:"following_count"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
		Serializer      string `json:"_serializer"`
	} `json:"last_editor"`
	Book       string `json:"book"`
	Serializer string `json:"_serializer"`
}

// RepoBody -
var RepoBody struct {
	Data []RepoDate `json:"data"`
}

// RepoDate -
type RepoDate struct {
	ID               int64  `json:"id"`
	Type             string `json:"type"`
	Slug             string `json:"slug"`
	Name             string `json:"name"`
	UserID           int64  `json:"user_id"`
	Description      string `json:"description"`
	CreatorID        int64  `json:"creator_id"`
	Public           int    `json:"public"`
	ItemsCount       int    `json:"items_count"`
	LikesCount       int    `json:"likes_count"`
	WatchesCount     int    `json:"watches_count"`
	ContentUpdatedAt string `json:"content_updated_at"`
	UpdatedAt        string `json:"updated_at"`
	CreatedAt        string `json:"created_at"`
	Namespace        string `json:"namespace"`
	User             struct {
		ID               int64  `json:"id"`
		Type             string `json:"type"`
		Login            string `json:"login"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		AvatarURL        string `json:"avatar_url"`
		LargeAvatarURL   string `json:"large_avatar_url"`
		MediumAvatarURL  string `json:"medium_avatar_url"`
		SmallAvatarURL   string `json:"small_avatar_url"`
		BooksCount       int    `json:"books_count"`
		PublicBooksCount int    `json:"public_books_count"`
		FollowersCount   int    `json:"followers_count"`
		FollowingCount   int    `json:"following_count"`
		CreatedAt        string `json:"created_at"`
		UpdatedAt        string `json:"updated_at"`
		Serializer       string `json:"_serializer"`
	} `json:"user"`
	Serializer string `json:"_serializer"`
}
