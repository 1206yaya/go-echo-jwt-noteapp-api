# NoteApp API Server

このプロジェクトは echo を使用して JWT 認証を実装したノートアプリケーションの API サーバーです。

[react-jwt-noteapp-web](https://github.com/1206yaya/react-jwt-noteapp-web) と連携して使用してください。

## データベース

データベースは PostgreSQL を使用しています

## テーブル

- users
- notes

## Features

- [x] docker でデータベースを作成
- [x] gorm を使用してデータベースに接続する
- [x] User Sign up 機能
- [x] User Looing / Logout 機能
- [x] Note CRUD 機能
- [ ] ozzo-validation を使用してバリデーションを実装する
