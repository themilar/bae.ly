package bae

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/mitchellh/go-homedir"
)

type BaeShorten struct {
	Error    int    `json:"error"`
	Id       string `json:"id"`
	ShortUrl string `json:"shortUrl"`
	Msg      string `json:"message"`
}
type url struct {
	Id       int
	Longurl  string
	Shorturl string
}
type BaeList struct {
	Error int `json:"error"`
	Data  struct {
		Url []url `json:"urls"`
	} `json:"data"`
	Msg string `json:"message"`
}

var root, _ = homedir.Dir()

func Shorten(u string, exp int) string {
	godotenv.Load(".env")
	API_KEY := os.Getenv("API_KEY")
	BASE_URL := os.Getenv("BASE_URL")
	c := &http.Client{Timeout: time.Second * 10}
	expiry := time.Now().Add(time.Duration(exp) * time.Hour).UTC()
	body := bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"url":"%s"
		}`, u)))
	if exp != 0 {
		body = bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"url":"%s",
			"expiry":"%s"
		}`, u, expiry)))
	}
	if BASE_URL == "" {
		BASE_URL = "https://urlbae.com/api/url/"
	}
	req, err := http.NewRequest("POST", BASE_URL+"add", body)
	if err != nil {
		log.Fatal(err)
	}
	if API_KEY == "" {
		file, err := os.ReadFile(root + "/.baerc")
		if err != nil {
			log.Fatal("provide your api key using the \"auth\" command")
		}
		content := string(file)
		API_KEY = strings.Split(content, "\n")[1]
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
func Auth(token string) {
	var key string
	if file, err := os.ReadFile(root + "/.baerc"); err == nil {
		content := strings.Split(string(file), "\n")
		if len(content) > 1 && token != content[1] {
			key = token
			os.WriteFile(root+"/.baerc", []byte("API_KEY\n"+token), 0644)
		} else {
			key = content[1]
		}
		fmt.Println(key)
	}
	_ = os.WriteFile(root+"/.baerc", []byte("API_KEY\n"+token), 0644)

}
func List(limit int) string {
	godotenv.Load(".env")
	API_KEY := os.Getenv("API_KEY")
	BASE_URL := os.Getenv("BASE_URL")
	c := &http.Client{Timeout: time.Second * 10}
	if BASE_URL == "" {
		BASE_URL = "https://urlbae.com/api/urls?"
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(BASE_URL+"limit=%d&page=1&order=date", limit), nil)
	if err != nil {
		log.Fatal(err)
	}
	if API_KEY == "" {
		file, err := os.ReadFile(root + "/.baerc")
		if err != nil {
			log.Fatal("provide your api key using the \"auth\" command")
		}
		content := string(file)
		API_KEY = strings.Split(content, "\n")[1]
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
	var response BaeList
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		fmt.Println(err)
	}
	if response.Error == 1 {
		return response.Msg
	}
	return prettyPrintStruct(fmt.Sprintf("%v", response.Data))
}

func prettyPrintStruct(obj interface{}) string {
	bytes, _ := json.MarshalIndent(obj, "\t", "\t")
	return string(bytes)
}
