package bae

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type BaeShorten struct {
	Error    int    `json:"error"`
	Id       string `json:"id"`
	ShortUrl string `json:"shortUrl"`
	Msg      string `json:"message"`
}

func Shorten(u string) string {
	godotenv.Load(".env")
	API_KEY := os.Getenv("API_KEY")
	BASE_URL := os.Getenv("BASE_URL")

	c := &http.Client{Timeout: time.Second * 10}
	body := bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"url":"%s"
		}`, u)))
	req, err := http.NewRequest("POST", BASE_URL+"add", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+API_KEY)
	req.Header.Add("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		log.Fatal("request failed")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("request error")
	}
	var shortendResponse BaeShorten
	if err := json.NewDecoder(res.Body).Decode(&shortendResponse); err != nil {
		fmt.Println(err)
	}
	if shortendResponse.Error == 1 {
		return shortendResponse.Msg
	}
	return shortendResponse.ShortUrl

}
