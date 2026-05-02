# kraken

![Version Badge](https://img.shields.io/badge/version-1.0.0-blue)
![Language](https://img.shields.io/badge/language-golang-green)

kraken はシンプルな SNS バックエンド API サービスです。

---

## 目次

- [必要な環境](#必要な環境)
- [使用パッケージ](#使用パッケージ)
- [プロジェクト構成](#プロジェクト構成)
- [セットアップ](#セットアップ)
- [起動方法](#起動方法)
- [API エンドポイント](#api-エンドポイント)

---

## 必要な環境

| ツール | バージョン |
|--------|-----------|
| Go     | 1.19 以上  |
| MySQL  | 8.0 以上   |
| Docker | 任意       |

---

## 使用パッケージ

| パッケージ | 用途 |
|-----------|------|
| [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) | MySQL ドライバ |
| [github.com/joho/godotenv](https://github.com/joho/godotenv) | `.env` ファイルの読み込み |
| [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | パスワードのハッシュ化 |

依存関係の管理は Go Modules (`go.mod` / `go.sum`) で行っています。

---

## プロジェクト構成

```
kraken/
├── cmd/
│   └── kraken/
│       └── main.go          # エントリーポイント
├── config/
│   ├── db/
│   │   └── db.go            # DB 接続
│   └── logging/
│       └── logging.go       # ログ設定
├── internal/
│   ├── route.go             # ルーティング
│   ├── help/
│   │   └── help.go          # ヘルプエンドポイント
│   ├── users/
│   │   └── register.go      # ユーザー登録
│   └── posts/
│       └── post.go          # 投稿
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── .env                     # 環境変数 (要作成)
```

---

## セットアップ

### 1. リポジトリのクローン

```bash
git clone https://github.com/your-org/kraken.git
cd kraken
```

### 2. 依存パッケージのインストール

```bash
go mod download
```

### 3. 環境変数の設定

プロジェクトルートに `.env` ファイルを作成し、以下を設定します。

```env
DB_USERNAME=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=kraken
```

### 4. データベースのセットアップ

MySQL に接続し、以下のテーブルを作成します。

```sql
CREATE TABLE user_accounts (
    id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at    DATETIME NOT NULL,
    updated_at    DATETIME NOT NULL
);

CREATE TABLE posts (
    id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_account_id BIGINT UNSIGNED NOT NULL,
    content         TEXT NOT NULL,
    created_at      DATETIME NOT NULL,
    updated_at      DATETIME NOT NULL
);
```

---

## 起動方法

### ローカル起動

```bash
go run cmd/kraken/main.go
```

サーバーは `localhost:8080` で起動します。

### Docker を使った起動

```bash
docker-compose up --build
```

サーバーはホストの `9000` ポートで公開されます (`localhost:9000`)。

---

## API エンドポイント

ベースパス: `/api/v1`

| メソッド | エンドポイント          | 説明             |
|---------|------------------------|-----------------|
| GET     | `/api/v1/help`         | ヘルプの表示      |
| POST    | `/api/v1/users/register` | ユーザー登録    |
| POST    | `/api/v1/post/post`    | 投稿の作成        |

### GET `/api/v1/help`

```bash
curl http://localhost:8080/api/v1/help
```

### POST `/api/v1/users/register`

```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "secret", "name": "Taro"}'
```

### POST `/api/v1/post/post`

```bash
curl -X POST http://localhost:8080/api/v1/post/post \
  -H "Content-Type: application/json" \
  -d '{"content": "Hello, kraken!"}'
```
