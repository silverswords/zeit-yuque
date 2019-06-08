package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Details -
var Details struct {
	Abilities Abilities `json:"abilities"`
	Data      Data      `json:"data"`
}

// BookDetail -
func BookDetail(w http.ResponseWriter, r *http.Request) {
	repoid := r.FormValue("RepoID")
	id := r.FormValue("ID")

	host := "https://www.yuque.com"
	url := host + "/api/v2/repos/" + repoid + "/docs/" + id + "?raw=0"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR:%s", err)
		return
	}

	Token := r.Header
	request.Header.Add("X-Auth-Token", Token["X-Auth-Token"][0])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR:%s", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR:%s", err)
		return
	}

	err = json.Unmarshal(body, &Details)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR:%s", err)
		return
	}

	if Details.Data.Body == "" {
		log.Println("Body is nil")
		fmt.Fprintf(w, "ERROR:%s", "Body is nil")
		return
	}

	fmt.Fprintf(w, "%s", Details.Data.Body)
	fmt.Fprintf(w, "****%s****", Details.Data.PublishedAt)
}

// Abilities -
type Abilities struct {
	Update  bool `json:"update"`
	Destroy bool `json:"destroy"`
}

// Data -
type Data struct {
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
