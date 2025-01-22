package main

import (
	// ---------- 標準ライブラリ（ログ出力、HTTPサーバー、HTMLテンプレート）----------
	"log"
	"net/http"
	"html/template"
	// ----------------------------------------

	"github/Abiti0233/google_auth_test/auth"

	// URLのルーティングを柔軟に行うための外部ライブラリ
	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	// muxを使用して新しいルーターを作成
	r := mux.NewRouter()

	// GETリクエストで/(ルートURL)が呼ばれたらhomeHandler関数を実行
	r.HandleFunc("/", homeHandler).Methods("GET")

	// 下記も同じような感じでルーティングを設定
	r.HandleFunc("/login", auth.GoogleLoginHandler).Methods("GET")

	// Google側で認証後に呼び出されるコールバックURL。Googleの認証画面で許可を与えられた後、このURLにリダイレクトされる。
	r.HandleFunc("/auth/google/callback", auth.GoogleCallbackHandler).Methods("GET")
	r.HandleFunc("/success", successHandler).Methods("GET")

	log.Println("Starting server on :8081")

	// エラーがあったらlog.Fatalでログを出力して終了
	log.Fatal(http.ListenAndServe(":8081", r))
}

// ルートURLへGETリクエストがあったら呼ばれる関数
// w http.ResponseWriter: レスポンスをクライアントに送信するためのインターフェース
// r *http.Request: クライアントからのリクエスト情報を含む構造体へのポインタ
func homeHandler(w http.ResponseWriter, r *http.Request) {

	// home.htmlというテンプレートを使ってHTMLを生成し、ブラウザに返却する
	templates.ExecuteTemplate(w, "home.html", nil)
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	// ログイン成功時には、ユーザー情報を取得してsuccess.htmlに表示する
	userInfo := r.Context().Value("user")
	templates.ExecuteTemplate(w, "success.html", userInfo)
}
