# AwesomeMusic-api サーバ

音楽サービス AwesomeMusic の API サーバ
https://github.com/oshou/Portfolio

## 概要

- JSON 形式でレスポンス
- 環境変数はレポジトリ直下の.env を読み込む
- Port8080 で受付

## 処理フロー

```plantuml
@startuml
box "interface" #LightBlue
  participant server
  participant controller
end box
box "usecase" #Yellow
  participant service
end box
box "domain" #Pink
  participant repository
  participant domain
end box

activate server
server -> server: APIサーバ起動
server -> server: 環境変数読み込み
server -> server: DBコネクション生成
server -> server: CQRSポリシー設定
[-> server :request
server -> controller: routing
controller -> controller: request解釈
controller -> service: service呼出
service -> repository: Modelメソッド呼出
repository -> domain:
domain -> repository:
repository -> service: 呼出結果返却
service -> controller: 呼出結果返却
controller -> controller: response生成
controller -> server: response
@enduml
```

```plantuml
@startuml
activate server
server -> server: APIサーバ起動
server -> server: 環境変数読み込み
server -> server: DBコネクション生成
server -> server: CQRSポリシー設定
[-> server :request
server -> controller: routing
controller -> controller: requestパラメータ解釈
controller -> service: serviceを実行
service -> service: ビジネスロジック前処理
service -> DB: modelメソッドを実行
DB -> service: 処理結果返却
service -> service: ビジネスロジック後処理
service -> controller: 処理結果返却
controller-> controller: responseデータ生成
controller-> server: response
server ->[: response
@enduml
```

### ディレクトリ構成

- Makefile
  - タスクランナー
- main.go
  - エントリポイント
- go.mod, go.sum
  - Go パッケージ管理
- server/
  - API サーバ起動、ルーティング、環境変数読み込み、CQRS 対応
- controller/
  - リクエストの Parse,レスポンス生成
- db/
  - DB 接続
- entity/
  - DB テーブル定義
- facade/
  - ビジネスロジック
- .env.xxx
  - 環境変数
  - .env が利用される、利用時は.env.xxx を.env として別名コピーが必要
- Dockerfile
  - 実行環境作成用
  - Multi-Stage Build でビルドコンテナ、起動用コンテナを分けている
- .circleci
  - CI 関連

## 操作方法

- ビルド
  - make build_local
- バイナリ実行
  - make run
