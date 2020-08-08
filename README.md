# AwesomeMusic-api Server

![GithubAction](https://github.com/oshou/AwesomeMusic-api/workflows/develop/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/oshou/AwesomeMusic-api)

音楽サービス AwesomeMusic の API サーバ<br>
https://github.com/oshou/Portfolio

## 概要

- Onion Architecture を参考にしています。
- JSON 形式でレスポンスを返します。
- 環境変数はレポジトリ直下の.env を読み込みます。<br>
  ビルド時に指定の「.env.(環境名)」を「.env」としてコピー作成いたします。<br>
  ex) cp -rp .env.local -> .env
- デフォルトで Port 8080 で受付いたします。

## APIドキュメント
```bash
1. make gendoc
2. http://localhost:10080にアクセス
```

## Quick Start

```bash
go get -u github.com/oshou/AwesomeMusic-api
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
# APIサーバ起動
make run
```

## ディレクトリ構成

```
└── main.go                  # Entry Point
├── injector/                # DI
├── domain/                  # Domain Layer
│   └── model/               # Domain Model Layer
│   └── repository/          # Domain Service Layer
├── usecase/                 # Application Layer
├── ui/                      # UI Layer
│   └── http/                # http
│         └── router/        # http router
│         └── middleware/    # http middleware
│         └── handler/       # http handler
├── infrastructure/          # Infrastructure Layer
│   └── persistence/         # Persistence
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
