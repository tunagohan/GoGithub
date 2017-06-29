# GoGithub
GoでGithubの情報を取得する

## GithubAPIのgo-githubを利用してみる

Githubのアクセストークンを取得する
[Sing in to Github](https://github.com/settings/tokens)
`Generate new token`をしてアクセストークンをコピーしてください。

## チュートリアル

`RepoList.go`の中にある
"...Your Access Token..."
の部分を先ほどコピーしたアクセストークンに置き換えて
`go run RepoList.go`してください。

自分のリポジトリのListが全て表示されると思います。

実際はJSON形式で取得しているので
上手く整形できる方お願いします。

## 上手く行かない方
```
go get github.com/google/go-github
```
をして見てください。
