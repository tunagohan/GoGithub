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

type Payload struct {
	Action  string
	Comment []Comment
	Number  int
}

type Comment struct {
	body      string
	CreatedAT string
}

type Data struct {
	Type      string  `json:"type"`
	PayLoad   Payload `json:"payload"`
	CreatedAt string  `json:"created_at"`
}

func main() {
	fmt.Println("ユーザー名を入力してください")
	fmt.Print("YourName: ")
	comments := PullRequestComments()
	fmt.Println(comments)
}

func PullRequestComments() (pullreqtotal int) {
	ownerName := getUserName() + "/events"
	url := "https://api.github.com/users/" + ownerName
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	var d []Data
	dec.Decode(&d)
	var count = 0
	for _, json := range d {
		// 日付文字列パース
		var ts = strings.Replace(json.CreatedAt, "T", " ", 1)
		ts = strings.Replace(ts, "Z", " UTC", 1)
		// JST置換
		t, _ := time.Parse("2006-01-02 15:04:05 MST", ts)
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		tJst := t.In(jst)
		// 日付比較
		nowJst := time.Now()
		if tJst.Format("2006-01-02") == nowJst.Format("2006-01-02") &&
			json.Type == "PushEvent" {
			count += len(json.PayLoad.Comment)
		}
	}
	return count
}

func getUserName() (stringReturned string) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	stringReturned = sc.Text()
	return
}
