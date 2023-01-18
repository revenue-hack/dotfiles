# convert - 動画変換関数の処理フロー

```mermaid
sequenceDiagram
    autonumber
    actor User as User
    participant S3_origin as S3<br>(動画アップロード用)
    participant S3_deliv as S3<br>(配信用)
    participant EB as EventBridge
    participant Lambda as Lambda<br>動画変換関数
    participant DB as Aurora
    participant MC as Elemental<br>MediaConvert

    User->>S3_origin: 動画(mp4)をアップロード
    EB-->>S3_origin: PutObjectをwatch
    EB->>Lambda: 起動

    Lambda->>Lambda: 起動ログ出力

    alt 不正な起動パラメーター
      Lambda-->EB: 終了
    end

    Lambda->>Lambda: PUTされたパスをパース
    alt 不正な形式のパス
      Lambda-->EB: 終了
    end

    Lambda->>DB: xxx_moviewsから動画情報を検索
    alt 変換処理起動済み
      Lambda-->EB: 終了
    end

    Lambda->>DB: 動画の変換ステータスを更新（アップロード完了）
    alt DB更新に失敗
      Lambda-->EB: 終了
    end

    Lambda->>Lambda: Hash値を生成
    Lambda->>MC: 動画変換ジョブを作成
    alt ジョブ作成に失敗
      Lambda->>DB: 動画の変換ステータスを更新（ジョブ作成エラー）
      Lambda-->EB: 終了
    end

    Lambda->>Lambda: 終了ログ出力
    Lambda-->EB: 終了

    Note over MC: 非同期処理
    MC->>S3_deliv: Apple HLSに変換したファイルをアップロード
```

### 起動ログ

以下の項目を出力する

- `detail.id`
- `detail.object.key` いる？
- `detail.object.size` いる？

### 起動パラメーター

https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/userguide/ev-events.html

**起動パラメーターの検証**

- `detail.reason` が `PutObject` であること
- `detail.object.key` に値が存在すること

### PUTされたパスをパース

パスのフォーマットは `users/{顧客コード}/{動画ID}/ファイル名.mp4`

### 変換処理起動済み

動画の変換ステータスが `アップロード完了` 移行のステータスの場合は処理をスキップ

変換処理が複数回呼ばれると不要な動画が作成されてコスト増につながるため。

### 動画変換ジョブを作成

SDK: https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/mediaconvert

※ジョブ作成を実行する際に、起動パラメータをログ出力すること

**MediaConvertの起動パラメータ**

- 変換フォーマット: Apple HLS
- 出力先のパス: `users/{顧客コード}/movies/{動画ID}/{hash値}.m3u8`
  - TODO: インデックスファイル名指定出来る？できなかったらどういう形式になるかは確認しておく
- コーデック: AVC
- フレームレート: 30fps
- 解像度: HD（解像度が720 ~ 1,080）くらい？
- TODO: その他パラメータは要調査
- Tags
  - Key=Customer
    Value=`{顧客コード}`

### 終了ログ

以下の項目を出力する

- `detail.id`
- 動画変換ジョブのID
