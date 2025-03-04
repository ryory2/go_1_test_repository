package main

import (
	"fmt"
	"log"
	"net/http"
)

// "/" にアクセスされた場合のハンドラー
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, ECS on port 8080!(updated11)")
}

// "/health" にアクセスされた場合のヘルスチェックハンドラー
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// 必要に応じてDB接続などの内部チェックを追加可能
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
