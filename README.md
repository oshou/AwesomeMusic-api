# AwesomeMusic API

![GithubAction](https://github.com/oshou/AwesomeMusic-api/workflows/develop/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/oshou/AwesomeMusic-api)

音楽サービス AwesomeMusic の API サーバ<br>
https://github.com/oshou/Portfolio

## 概要

- Onion Architecture を参考にしています。
- ポートはデフォルト 8080 で受付
- 環境変数: レポジトリ直下の.env を読み込みます。<br>
  ビルド時に指定の「.env.(環境名)」を「.env」としてコピー作成します。<br>
  ex) cp -rp .env.local -> .env

## API ドキュメント

```bash
$ make apidoc
$ open http://localhost:10080
```

## ER 図

```bash
$ make dbdoc
```

## ディレクトリ構成

```
└─ .circleci                # CI config
└─ cmd/                     # CI config
│     └─ main.go           # Entry Point
└─ config/                  #
│     └─ config.go         # Config
└─ db/                      #
│     └─ db.go             # DB Connection
└─ deployments/             # Dockerfiles
│     └─ api/              #
│     └─ apitest/          #
│     └─ postgres/         #
├─ docs/                    # DI
├─ domain/                  # Domain Layer
│     └─ model/            # Domain Model Layer
│     └─ repository/       # Domain Service Layer
├─ infra/                   # Infrastructure Layer
│     └─ persistence/      # Persistence
└─ log/                      #
│     └─ logger.go             # DB Connection
└─ test/                      #
│     └─ test.http             # DB Connection
│     └─ test_api.tavern.yaml  # DB Connection
├─ ui/                      # UI Layer
│     └─ http/                # http
│           └─ middleware/    # http middleware
│           └─ handler/       # http handler
├─ usecase/                 # Application Layer
└─ Makefile                 # Task Runner
└─ go.mod                   # Go Package management
└─ go.sum                   # Go Package management
└─ .env                     # Config Environment
└─ Dockerfile               # Docker Container
```

## Quick Start

```bash
$ go get -u github.com/oshou/AwesomeMusic-api

# 環境変数を指定(.env.localに設定例を記載)
$ vim .env.local

# APIサーバ起動
$ make run
```

## 処理フロー

![sequence](https://github.com/oshou/AwesomeMusic-api/blob/img/out/docs/sequence/sequence.png)

## クラス構成(のようなもの)

![class](https://github.com/oshou/AwesomeMusic-api/blob/img/out/docs/class/class.png)
