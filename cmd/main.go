package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors" // github.com/rs/cors パッケージをインポート
)

// "/" にアクセスされた場合のハンドラー
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, ECS on port 8080!(updated14")
}

// "/health" にアクセスされた場合のヘルスチェックハンドラー
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// 必要に応じてDB接続などの内部チェックを追加可能
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {

	// CORS 設定
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:30010"},           // 許可するオリジン
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},           // 許可する HTTP メソッド
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"}, // 許可するヘッダー
		AllowCredentials: true,                                         // Cookie や Authorization ヘッダーを許可
		Debug:            true,                                         // デバッグモード (開発環境向け)
	})

	// ハンドラーを登録
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)

	// CORS ミドルウェアを適用
	handler := c.Handler(http.DefaultServeMux) // DefaultServeMux をラップ

	// サーバー起動
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
