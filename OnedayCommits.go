package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type actor struct {
	ID         int
	Login      string
	GravatarID string
	URL        string
	AvatarURL  string
}

type repo struct {
	ID   int
	Name string
	URL  string `json:"url"`
}

type payload struct {
	PushID       int
	Size         int
	DistinctSize int
	Ref          string
	Head         string
	Before       string
	Commits      []commits
}

type commits struct {
	SHA      string
	Message  string
	Distinct string
	URL      string
}

type data struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Actor     actor   `json:"actor"`
	Repo      repo    `json:"repo"`
	PayLoad   payload `json:"payload"`
	Public    string  `json:"public"`
	CreatedAt string  `json:"created_at"`
}

func main() {
	fmt.Print("UserName >>> ")
	ownername := getword() + "/events"
	url := "https://api.github.com/users/" + ownername
	req, _ := http.NewRequest("GET", url, nil)
	cl := new(http.Client)
	resp, _ := cl.Do(req)
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var d []data
	dec.Decode(&d)
	var count = 0
	for _, json := range d {
		// 日付の文字列をパースする
		var ts = strings.Replace(json.CreatedAt, "T", " ", 1)
		ts = strings.Replace(ts, "Z", " UTC", 1)
		// UTCをJSTに置換する
		t, _ := time.Parse("2006-01-02 15:04:05 MST", ts)
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		tJst := t.In(jst)
		// 今日の日付とパースした日付を比較する
		nowJst := time.Now()
		if tJst.Format("2006-01-02") == nowJst.Format("2006-01-02") &&
			json.Type == "PushEvent" {
			count += len(json.PayLoad.Commits)
		}
	}
	fmt.Println(count)
}

func getword() (stringReturned string) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	stringReturned = sc.Text()
	return
}
