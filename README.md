# microblog

個人用のマイクロブログです。仕様技術は、Golang（chi）・D1・GAE

## ソースコードをもとに依存関係の解決

```sh
go mod tidy
```

## ローカルでの起動

```sh
go run .
```

## GAE にデプロイ

```sh
# デプロイ
gcloud app deploy

# トップページ表示
gcloud app browse
```
