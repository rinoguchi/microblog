# microblog

個人用のマイクロブログです。仕様技術は、Golang（chi）・D1・GAE

## ソースコードをもとに依存関係の解決

```sh
go mod tidy
```

## ローカルでの起動

```sh
go run .

# コメント一覧取得
curl http://localhost:8080/comments

# コメント登録
curl POST -H "Content-Type: application/json" -d '{"text" : "dummy coment"}' http://localhost:8080/comments
```

- GET http://localhost:8080/comments
- POST http://localhost:8080/comments

## GAE にデプロイ

```sh
# デプロイ
gcloud app deploy

# トップページ表示
gcloud app browse
```
