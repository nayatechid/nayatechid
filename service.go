package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
)

var Users = []map[string]string{
	{"hashnode": "born2ngopi", "naya": "chandra"},
	{"hashnode": "hadihammurabi", "naya": "hadihammurabi"},
	{"hashnode": "nurkholis", "naya": "kholis"},
	{"hashnode": "hdinjos", "naya": "fullsnackdev"},
	{"hashnode": "yogahermawan", "naya": "yogahermawan"},
}

func GQL(query string) ([]byte, error) {
	cleanQuery := strings.ReplaceAll(query, "\n", "")
	cleanQuery = strings.ReplaceAll(cleanQuery, "\t", "")
	cleanQuery = strings.ReplaceAll(cleanQuery, " ", "")

	reqData := fmt.Sprintf(`{"query":"%s","variables":{}}`, cleanQuery)
	reqBody := strings.NewReader(reqData)
	resp, err := http.Post("https://api.hashnode.com/", "application/json", reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

type Post struct {
	Title      string `json:"title"`
	Brief      string `json:"brief"`
	Slug       string `json:"slug"`
	CoverImage string `json:"coverImage"`
	DateAdded  string `json:"dateAdded"`

	UsernameHashnode string
	UsernameNaya     string
}

type Publication struct {
	Posts []Post `json:"posts"`
}

type User struct {
	Publication Publication `json:"publication"`
}

type Data struct {
	User User `json:"user"`
}

type DtoGetPosts struct {
	Data Data `json:"data"`
}

func GetPostsAndMapUsername(usernameHashnode string, usernameNaya string) ([]Post, error) {
	resp, err := GQL(fmt.Sprintf(`
	query {
		user(username: \"%s\") {
			publication {
				posts {
					title,
					brief,
					slug,
					coverImage,
					dateAdded,
				},
			},
		},
	}
	`, usernameHashnode))
	if err != nil {
		return nil, err
	}

	respToDto := DtoGetPosts{}
	err = json.Unmarshal(resp, &respToDto)
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0)
	for _, p := range respToDto.Data.User.Publication.Posts {
		p.UsernameHashnode = usernameHashnode
		p.UsernameNaya = usernameNaya
		posts = append(posts, p)
	}

	return posts, nil
}

func SortPostsByDate(posts []Post) []Post {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].DateAdded > posts[j].DateAdded
	})

	return posts
}

func GetAllUsersPosts() ([]Post, error) {
	posts := make([]Post, 0)
	var wg sync.WaitGroup
	for _, u := range Users {
		wg.Add(1)
		go worderGetAllUsersPosts(u["hashnode"], u["naya"], &posts, &wg)
	}

	wg.Wait()
	return SortPostsByDate(posts), nil
}

func worderGetAllUsersPosts(usernameHashnode, usernameNaya string, posts *[]Post, wg *sync.WaitGroup) {
	defer wg.Done()
	postsPerUser, err := GetPostsAndMapUsername(usernameHashnode, usernameNaya)
	if err != nil {
		return
	}
	*posts = append(*posts, postsPerUser...)
}
