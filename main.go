package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Ip string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	ips := strings.Split(r.Header.Get("X-FORWARDED-FOR"), ", ")
	var page = Response{""}
	if len(ips[0]) > 0 {
		page = Response{ips[0]}
	} else {
		page = Response{r.RemoteAddr}
	}

	res, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}

func main() {
	http.HandleFunc("/", viewHandler) // hello
	envLoad()
	_ = http.ListenAndServe(":"+os.Getenv("port"), nil)
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
