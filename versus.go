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

type Actor struct {
	ID         int
	Login      string
	GravatarID string
	URL        string
	AvatarURL  string
}

type Repo struct {
	ID   int
	Name string
	URL  string `json:"url"`
}

type Payload struct {
	PushID       int
	Size         int
	DistinctSize int
	Ref          string
	Head         string
	Before       string
	Commits      []Commits
}

type Commits struct {
	SHA      string
	Message  string
	Distinct string
	URL      string
}

type Data struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Actor     Actor   `json:"actor"`
	Repo      Repo    `json:"repo"`
	PayLoad   Payload `json:"payload"`
	Public    string  `json:"public"`
	CreatedAt string  `json:"created_at"`
}

func main() {
	fmt.Println("ユーザー名を入力してください")
	fmt.Print("Player1: ")
	player1 := commitCount()
	fmt.Print("Player2: ")
	player2 := commitCount()
	if player1 > player2 {
		fmt.Println("Win : +", player1-player2)
		fmt.Printf("MyCommits:%d\nPlayer2 Commits:%d", player1, player2)
	}
	if player1 < player2 {
		fmt.Println("player1 Lose : -", player2-player1)
		fmt.Printf("MyCommits:%d\nPlayer2 Commits:%d", player1, player2)
	}
	if player1 == player2 {
		fmt.Println("DROW")
		fmt.Printf("MyCommits:%d\nPlayer2 Commits:%d", player1, player2)
	}

}

//commitCount is counts the number of commits .
func commitCount() (commitTotal int) {
	player := getWord() + "/events"
	url := "https://api.github.com/users/" + player
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
			count += len(json.PayLoad.Commits)
		}
	}
	return count
}

// getWord is type text .
func getWord() (stringReturned string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringReturned = scanner.Text()
	return
}
