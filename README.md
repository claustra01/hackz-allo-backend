# 開発の準備
- 新しいディレクトリを作成
- `git init`
- `git remote add origin https://github.com/claustra01/hackz-allo-backend`
- `git pull origin master`
- `git checkout -b develop`
- `git pull origin develop`
- `git checkout -b <作業するブランチ名>`
- **masterブランチ/developブランチには絶対に直接pushしないこと**

# ローカルで実行
- `go run main.go`
- http://localhost:1323/ にアクセス

# 更新内容を統合
- commit/pushを完了させる
- **developブランチ**にプルリクエストを出す
- mergeする
- **大きな更新が一段落したタイミングで**developブランチをmasterブランチにmergeする
- **masterブランチに直接mergeしないこと**
