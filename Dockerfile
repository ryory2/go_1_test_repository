# ビルドステージ: Go の公式イメージを使用
FROM golang:1.20-alpine AS builder

WORKDIR /app

# モジュールファイルとソースコードのコピー
COPY go.mod .
COPY main.go .
RUN go mod tidy

# アプリケーションのビルド
RUN go build -o server

# 実行ステージ: 軽量な Alpine イメージを利用
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .

# ポート8080を公開
EXPOSE 8080

# アプリケーションの起動
CMD ["./server"]
