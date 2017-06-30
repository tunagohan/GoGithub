# GoGithub
GoでGithubの情報を取得する

## GithubAPIのgo-githubを利用してみる

Githubのアクセストークンを取得する
[Sing in to Github](https://github.com/settings/tokens)

`Generate new token`をしてアクセストークンをコピーしてください。

## チュートリアル

`RepoList.go`
```
func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "{...Your Access Token...}"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "{...UserName...}", nil)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(repos)
}
```

- "{...Your Access Token...}"  
- "{...UserName...}"  
を自分のもの、もしくは他人のものに置き換えてください。

```
go run RepoList.go
```

指定したリポジトリのListが全て表示されると思います。

実際はJSON形式で取得しているので上手く整形できる方お願いします。

## 上手く行かない方

```
go get github.com/google/go-github
```

をしてみてください。
