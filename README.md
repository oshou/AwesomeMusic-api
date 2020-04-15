# AwesomeMusic-api サーバ

音楽サービス AwesomeMusic の API サーバ
https://github.com/oshou/Portfolio

## 概要

- JSON 形式でレスポンス
- 環境変数はレポジトリ直下の.env を読み込む
- Port8080 で受付
- OnionArchitectureを採用

## 処理フロー
![sequence](https://user-images.githubusercontent.com/4841735/79293900-c2f08100-7f0f-11ea-9ddd-cfa521302759.png)

## クラス構成(のようなもの)
![class](https://user-images.githubusercontent.com/4841735/79293889-bb30dc80-7f0f-11ea-89d7-36a980dc11ef.png)

### ディレクトリ構成

- Makefile
  - タスクランナー
- main.go
  - エントリポイント
- go.mod, go.sum
  - Golang パッケージ管理
- db/
  - DB 接続
- interactor/
  - 簡易 DI
- domain/
  - model/
    - Domain Model 層
  - repository/
    - Domain Service 層
- usecase/
  - Application 層
- presenter/
  - UI 層
- infrastructure/
  - Infrastructure 層
- docs/
  - UML 等
- .env.xxx
  - 環境変数
  - .env が利用される、利用時は.env.xxx を.env として別名コピーが必要
- Dockerfile
  - 実行環境作成用
  - Multi-Stage Build でビルドコンテナ、起動用コンテナを分けている
- .circleci
  - CI 関連

## 操作方法

- リンター実行
  - make lint
- ビルド
  - make build_local
- exec binary
  - make run
