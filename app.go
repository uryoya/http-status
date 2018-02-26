package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin/json"

	"github.com/uryoya/http-status/config"
)

type StatusResponse struct {
	StatusCode int
	Message    string
}

func main() {
	/**
	 * URL: /[status code]
	 * にマッチしたHTTP Status Code のレスポンスを返す
	 */
	reStatus := regexp.MustCompile(`^/([0-9]{3})`)

	errResp, err := json.Marshal(StatusResponse{500, "Internal Server Error"})
	if err != nil {
		panic(err)
	}
	notFoundResp, err := json.Marshal(StatusResponse{404, "Not Found"})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		matches := reStatus.FindStringSubmatch(r.URL.Path)
		if len(matches) < 2 {
			w.WriteHeader(http.NotFound) // マッチしなかった場合は404 Not Found
			return
		}
		fmt.Println(matches)
		switch matches[1] {
		case "200":
			w.WriteHeader(http.StatusOK)
			resp, err := json.Marshal(StatusResponse{200, "OK"})
			if err != nil {
				w.Write(errResp)
			} else {
				w.Write(resp)
			}
		case "201":
			w.WriteHeader(http.StatusCreated)
			resp, err := json.Marshal(StatusResponse{201, "Created"})
			if err != nil {
				w.Write(errResp)
			} else {
				w.Write(resp)
			}
		case "202":
			w.WriteHeader(http.StatusAccepted)
			resp, err := json.Marshal(StatusResponse{203, "Accepted"})
			if err != nil {
				w.Write(errResp)
			} else {
				w.Write(resp)
			}
		case "400":
			w.WriteHeader(400)
		case "500":
			w.WriteHeader(500)
		default:
			w.WriteHeader(404)
		}
	})

	http.ListenAndServe(config.Server.Port, nil)
}
