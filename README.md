# natural-language-api-playground
## セットアップ
- リポジトリ直下にサービスアカウント`svc.json`（ロール付与必要なし）を配置
    - https://cloud.google.com/natural-language/docs/setup?hl=ja#role
- `docker compose up app`
## 実行
```sh
❯ curl "localhost:8080/nyuwaize/こんにちは、私の名前は太郎です。最高の一日。"
{"message":"こんにちは🥰私の名前は太郎です😆最高の一日😆"}
```