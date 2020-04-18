# AwesomeMusic-api Server

音楽サービス AwesomeMusic の API サーバ
https://github.com/oshou/Portfolio

## 概要

- JSON 形式でレスポンスを返します。
- 環境変数はレポジトリ直下の.env を読み込みます。
  ビルド時に指定の.env.(環境名)のファイルを.env としてコピー作成されます。
  ex) cp -rp .env.local -> .env
- Port8080 で受付いたします。
- OnionArchitecture を採用しています。

## Quick Start

```bash
# 環境変数編集(.env.local)
vim .env
# ビルド(ローカル)
make build_local
# ビルド済バイナリ実行
make run
```

## ディレクトリ構成

```
├── domain/
│   └── model/        # Domain Model Layer
│   └── repository/   # Domain Service Layer
├── usecase/          # Application Layer
├── presenter/        # UI Layer
├── injector/         # DI
├── db/               # DB Connetion
└── Makefile          # Task Runner
└── main.go           # Entry Point
└── go.mod            # Go Package management
└── go.sum            # Go Package management
└── .env              # Config Environment
└── Dockerfile        # Docker Container
```

## 処理フロー

![sequence](https://github.com/oshou/AwesomeMusic-api/blob/img/out/docs/sequence/sequence.png)

## クラス構成(のようなもの)

![class](https://github.com/oshou/AwesomeMusic-api/blob/img/out/docs/class/class.png)
