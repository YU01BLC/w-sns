# URI一覧

| URI | メソッド | 詳細 |
| - | - | - |
| /user | GET | ユーザ一覧取得 |
| /user/{$user_id} | GET | ユーザ詳細取得 |
| /user/{$user_id} | POST | ユーザ編集 |

# Go環境構築

## 前提

* [YU01BLC/w-sns](https://github.com/YU01BLC/w-sns)をクローン済み
* リポジトリ直下のディレクトリをルートディレクトリとする

## バージョン情報

| 名称 | バージョン | URL |
| - | - | - |
| Go言語 | v1.22 | [The Go Programming Language](https://go.dev/) |
| Echo | v4 | [Quick Start &#124; Echo](https://echo.labstack.com/docs/quick-start) |
| goqu | v9 | [doug-martin/goqu: SQL builder and query library for golang](https://github.com/doug-martin/goqu) |

## ローカルへのインストール

[Download and install - The Go Programming Language](https://go.dev/doc/install)

上記のサイトからインストールします。
ルートディレクトリから移動してgo mod tidyを実行してライブラリをインストール

```bash
cd src/backend
go mod tidy
```

.envの作成

```bash
cp .env.example .env
```


## Docker立て方

```bash
docker compose up -d --build
```

## 最後に

> [!NOTE]
> 最後のアクセス方法はプロジェクトができたら変更するよ

http://localhost:8080 にアクセスして、`Hello world`が表示されたらOK
