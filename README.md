# 詳解Go言語Webアプリケーション開発 コードノート

このリポジトリは、「[詳解Go言語Webアプリケーション開発](https://www.amazon.co.jp/%E8%A9%B3%E8%A7%A3Go%E8%A8%80%E8%AA%9EWeb%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E9%96%8B%E7%99%BA-%E6%B8%85%E6%B0%B4-%E9%99%BD%E4%B8%80%E9%83%8E/dp/4863543727)」の学習メモとして作成されたコードが格納されています。

## 使用される技術:

- GitHub Actions: テスト
- MySQL: データベース
- Redis: JWTトークンの格納
- Docker: プロジェクトの構築

以下の表は主な機能とそれに対応するURIをまとめたものです:

| 機能           | HTTPメソッド | URI        |
| -------------- | ----------- | ---------- |
| ユーザー登録   | POST        | /register  |
| ユーザーログイン | POST        | /login     |
| ToDoタスクのログイン | GET         | /tasks     |
| ToDoタスクの一覧表示 | POST        | /tasks     |
| 管理者権限     | GET         | /admin     |

## 使用方法:

プロジェクトはmakeコマンドを用いて操作できます。以下は利用可能なコマンド一覧です:

```make
make build            # Dockerイメージをデプロイのためにビルドします。
make build-local      # Dockerイメージをローカル開発のためにビルドします。
make up               # Docker compose upをホットリロードとともに実行します。
make down             # Docker compose downを実行します。
make logs             # Docker compose logsを確認します。
make ps               # コンテナの状態を確認します。
make test             # テストを実行します。
make dry-migrate      # マイグレーションを試行します。
make migrate          # マイグレーションを実行します。
make generate         # コードを生成します。
make help             # オプションを表示します。
```