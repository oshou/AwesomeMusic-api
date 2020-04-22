# AwesomeMusic-api Server

![GithubAction](https://github.com/oshou/AwesomeMusic-api/workflows/develop/badge.svg)

音楽サービス AwesomeMusic の API サーバ<br>
https://github.com/oshou/Portfolio

## 概要

- JSON 形式でレスポンスを返します。
- 環境変数はレポジトリ直下の.env を読み込みます。
  ビルド時に指定の.env.(環境名)のファイルを.env としてコピー作成いたします。
  ex) cp -rp .env.local -> .env
- デフォルトで Port 8080 で受付いたします。
- Onion Architecture を参考にしています。

## Quick Start

```bash
# 環境変数指定(.env.local)
# - DB_DRIVER
# - DB_HOST
# - DB_PORT
# - DB_USER
# - DB_PASSWORD
# - DB_NAME
# - DB_OPTION
# - API_SERVER_PORT
vim .env.local
# ビルド(ローカル)
make build_local
# ビルド済バイナリ実行
make run
```

## ディレクトリ構成

```
└── main.go                  # Entry Point
├── injector/                # DI
├── domain/                  # Domain Layer
│   └── model/               # Domain Model Layer
│   └── repository/          # Domain Service Layer
├── service/                 # Application Layer
├── ui/                      # UI Layer
│   └── http/                # http
│         └── router/        # http router
│         └── middleware/    # http middleware
│         └── handler/       # http handler
├── infrastructure/          # Infrastructure Layer
│   └── datastore/           # Datastore
├── db/                      # DB Connetion
└── Makefile                 # Task Runner
└── go.mod                   # Go Package management
└── go.sum                   # Go Package management
└── .env                     # Config Environment
└── Dockerfile               # Docker Container
```

## 処理フロー

![sequence](https://github.com/oshou/AwesomeMusic-api/blob/img/out/docs/sequence/sequence.png)

## クラス構成(のようなもの)

![class](https://github.com/oshou/AwesomeMusic-api/blob/img/out/docs/class/class.png)
