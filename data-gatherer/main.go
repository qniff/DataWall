package main

import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"DataWall/data-gatherer/config"
)

type Device struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        int     `json:"z"`
	UserType int     `json:"userType"`
	Hash     string  `json:"hash"`
}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func serve(w http.ResponseWriter, h *http.Request) {

	tokenSource := &TokenSource{
		AccessToken: config.Get().Token,
	}
	resp, _ := oauth2.NewClient(oauth2.NoContext, tokenSource).Get(config.Get().Url)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(body))
}

func getInfo(t time.Time) {
	tokenSource := &TokenSource{
		AccessToken: config.Get().Token,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	resp, err := oauthClient.Get(config.Get().Url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	jsondata := string(body)

	var devices []Device // stores json in struct

	err = json.Unmarshal([]byte(jsondata), &devices)

	fmt.Printf("%+v", devices)
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8000", nil)

	doEvery(20000*time.Millisecond, getInfo)
}
