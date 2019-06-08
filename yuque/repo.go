package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Repo -
var Repo struct {
	Data []Book `json:"data"`
}

// Repostory -
func Repostory(w http.ResponseWriter, r *http.Request) {
	var book = make(map[string]int64)
	id := r.FormValue("RepoID")

	Token := r.Header

	host := "https://www.yuque.com"
	url := host + "/api/v2/groups/" + id + "/repos/"
	fmt.Println(url)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR:%s", err)
		return
	}

	request.Header.Add("X-Auth-Token", Token["X-Auth-Token"][0])

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

	err = json.Unmarshal(body, &Repo)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR:%s", err)
		return
	}

	if Repo.Data == nil {
		log.Println("Body is nil")
		fmt.Fprintf(w, "ERROR:%s", "Body is nil")
		return
	}

	for _, q := range Repo.Data {

		book[q.Name] = q.ID
	}
	fmt.Println(book)

	fmt.Fprintf(w, "%v", book)
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
