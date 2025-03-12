# Go で OAuth アプリを作成する

## ローカル起動

```sh
# Authlete の認可サーバ (java-oauth-server) を起動
[java-oauth-server プロジェクト]
docker-compose up

# フロント
[front]
npm run dev
```

## 必要なもの

- 認可サーバ
  Authlete で作るか Go で作るか
- フロントエンド
  React なり Svelte なり Sveltekit なり
- API サーバ
  Go

## 作業

### 認可サーバの作成

- [Authlete のチュートリアル](https://www.authlete.com/ja/developers/tutorial/) を参考に認可サーバを作成

  - サービス名: go-oauth-demo
  - アプリ名: go-oauth-demo-client
    go-oauth-demo サービスのコンソール画面に入ると確認できる
    (ログイン ID: API キー, パスワード: API シークレット)

- 疎通確認

```sh
curl -s -X POST https://api.authlete.com/api/auth/authorization \
-u '<API Key e.g. 10723797812772>:<API Secret e.g. ekYoYTI84qZcpe6bXGzDwduQ1fGBYxJT8K8Tnwd7poc>' \
-H 'Content-Type: application/json' \
-d '{ "parameters": "redirect_uri=https://client.example.org/cb/example.com&response_type=code&client_id=<Client ID e.g. 12800697055611>" }'
```

### フロント

- [SvelteKit チュートリアル](https://svelte.jp/docs/kit/creating-a-project)
