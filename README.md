# microblog

個人用のマイクロブログです。主な使用技術は、Golang（chi ＋ wire ＋ oapi-codegen ＋ bun）・GAE・PostgreSQL（supabase）です。

# インストール

```sh
make go-install
```

[参考] ソースコードをもとに依存関係を整理

```sh
go mod tidy
```

## ローカルでの起動

```sh
make serve

# コメント一覧取得
curl http://localhost:8080/comments

# コメント登録
curl -X POST -H "Content-Type: application/json" -d '{"text" : "あいうえお"}' http://localhost:8080/comments
```

- GET http://localhost:8080/comments
- POST http://localhost:8080/comments

## GAE にデプロイ

```sh
make deploy
```
