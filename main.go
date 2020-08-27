package main

import (
	"encoding/json"
	"net/http"
)

type Response struct { // テンプレート展開用のデータ構造
	Ip string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	page := Response{r.RemoteAddr}

	res, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}

func main() {
	http.HandleFunc("/", viewHandler) // hello
	_ = http.ListenAndServe(":80", nil)
}
