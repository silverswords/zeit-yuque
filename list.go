package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	core "github.com/silverswords/clouds"
	pub "github.com/silverswords/zeit-yuque/public"
)

// List -
var List struct {
	Data []Books `json:"data"`
}

// BookList -
func BookList(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			RepoID string `zeit:"required"`
		}
	)

	c := pub.NewContext(w, r)
	err := c.BindJSON(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusBadRequest, pub.H{"status": http.StatusBadRequest})
		return
	}

	err = core.Validate(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusNotAcceptable, pub.H{"status": http.StatusNotAcceptable})
		return
	}

	url := fmt.Sprintf(pub.ListURL, yuque.RepoID)

	body, err := c.CallAPI(url)
	if err != nil {
		c.WriteJSON(http.StatusRequestTimeout, pub.H{"status": http.StatusRequestTimeout})
		return
	}

	err = json.Unmarshal(body, &List)
	if err != nil {
		c.WriteJSON(http.StatusForbidden, pub.H{"status": http.StatusForbidden})
		return
	}

	c.WriteJSON(http.StatusOK, pub.H{"status": http.StatusOK, "List": string(body)})
}

// Books -
type Books struct {
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
