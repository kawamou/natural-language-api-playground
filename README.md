# natural-language-api-playground
## セットアップ
- 自身のGCPプロジェクトでNatural Language APIを有効化
- 認証情報を追加からサービスアカウントを作成
- リポジトリ直下にサービスアカウント`svc.json`（ロール付与必要なし）を配置
    - https://cloud.google.com/natural-language/docs/setup?hl=ja#role
- `docker compose up app`
## 実行
### ポジネガ判定
```sh
❯ curl "localhost:8080/analyze-sentiment/こんにちは、私の名前は太郎です。最高の一日。"
{"positive_nagative":"positive"}
❯ curl "localhost:8080/analyze-sentiment/こんにちは、私の名前は太郎です。辛い一日。"
{"positive_nagative":"negative"}
```

### 単語分析
```sh
❯ curl "localhost:8080/analyze-entity/こんにちは、私の名前は太郎です。楽しい一日。バッキンガム宮殿。Macbookカメラ"
{"entities":[{"Name":"こんにち","Type":"OTHER","Metadata":null},{"Name":"名前","Type":"OTHER","Metadata":null},{"Name":"太郎","Type":"PERSON","Metadata":null},{"Name":"バッキンガム","Type":"OTHER","Metadata":null},{"Name":"宮殿","Type":"PERSON","Metadata":null},{"Name":"Macbook","Type":"CONSUMER_GOOD","Metadata":null},{"Name":"カメラ","Type":"CONSUMER_GOOD","Metadata":null},{"Name":"一","Type":"NUMBER","Metadata":null}]}
```

### 単語感情分析
日本語はほとんど対応していなさそう？🤔
```sh
❯ curl "localhost:8080/analyze-entity-sentiment/zoo,park,disney,star"
{"entities":[{"Name":"zoo","Type":"OTHER","Metadata":null,"Sentiment":{"Magnitude":0.3,"Score":0.3}},{"Name":"park","Type":"LOCATION","Metadata":null,"Sentiment":{"Magnitude":0.4,"Score":0.4}},{"Name":"star","Type":"PERSON","Metadata":null,"Sentiment":{"Magnitude":0.3,"Score":0.3}},{"Name":"disney","Type":"ORGANIZATION","Metadata":null,"Sentiment":{"Magnitude":0.5,"Score":0.5}}]}
```

### 柔和化
```sh
❯ curl "localhost:8080/nyuwaize/こんにちは、私の名前は太郎です。最高の一日。"
{"message":"こんにちは🥰私の名前は太郎です😆最高の一日😆"}
```
