package yuque

import (
	"encoding/json"
	"fmt"
	"net/http"

	core "github.com/silverswords/clouds/context"
	base "github.com/silverswords/zeit-yuque"
)

// Repostory -
func Repostory(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			GroupID string `zeit:"required"`
		}
	)
	c := core.NewContext(w, r)
	err := c.BindJSON(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusBadRequest, core.H{"status": http.StatusBadRequest})
		return
	}

	err = core.Validate(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusNotAcceptable, core.H{"status": http.StatusNotAcceptable})
		return
	}

	url := fmt.Sprintf(base.RepoURL, yuque.GroupID)

	body, err := c.CallAPI(url)
	if err != nil {
		c.WriteJSON(http.StatusRequestTimeout, core.H{"status": http.StatusRequestTimeout})
		return
	}

	err = json.Unmarshal(body, &Repo)
	if err != nil {
		c.WriteJSON(http.StatusForbidden, core.H{"status": http.StatusForbidden})
		return
	}

	c.WriteJSON(http.StatusOK, core.H{"status": http.StatusOK, "Repo": Repo})
}

// Repo -
var Repo struct {
	Data []Book `json:"data"`
}

// Book -
type Book struct {
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
