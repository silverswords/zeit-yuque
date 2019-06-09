package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// List -
var List struct {
	Data []Books `json:"data"`
}

// BookList -
func BookList(w http.ResponseWriter, r *http.Request) {
	var book = make(map[string]string)

	repoid := r.FormValue("RepoID")
	host := "https://www.yuque.com"
	url := host + "/api/v2/repos/" + repoid + "/docs/"

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

	err = json.Unmarshal(body, &List)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR:%s", err)
		return
	}

	if List.Data == nil {
		log.Println("Body is nil")
		fmt.Fprintf(w, "ERROR:%s", "Body is nil")
		return
	}

	for _, v := range List.Data {
		bookID := strconv.FormatInt(v.ID, 10)
		book[v.Title] = bookID
	}

	fmt.Fprintf(w, "%v", book)
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
