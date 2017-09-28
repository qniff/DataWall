package main

import (
	"encoding/json"
	"fmt"
	"time"
	"golang.org/x/oauth2"
	"io/ioutil"
	//"net/http"
	//"bytes"
	//"reflect"
)

type Device struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        int     `json:"z"`
	UserType int     `json:"userType"`
	Hash     string  `json:"hash"`
}

/* Token below must be replaced with a valid one */
var personalAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6ImdyQWk2cnJRU0JiVVItY01ZOHpRTHE2aGdVQSIsImtpZCI6ImdyQWk2cnJRU0JiVVItY01ZOHpRTHE2aGdVQSJ9.eyJpc3MiOiJodHRwczovL2lkZW50aXR5LmZoaWN0Lm5sIiwiYXVkIjoiaHR0cHM6Ly9pZGVudGl0eS5maGljdC5ubC9yZXNvdXJjZXMiLCJleHAiOjE1MDYzNzY5MDAsIm5iZiI6MTUwNjM2OTcwMCwiY2xpZW50X2lkIjoiYXBpLWNsaWVudCIsInVybjpubC5maGljdDp0cnVzdGVkX2NsaWVudCI6InRydWUiLCJzY29wZSI6WyJvcGVuaWQiLCJwcm9maWxlIiwiZW1haWwiLCJmaGljdCIsImZoaWN0X3BlcnNvbmFsIiwiZmhpY3RfbG9jYXRpb24iXSwic3ViIjoiNWRlY2Y0YTgtNDhiZC00N2VjLTk4YjQtYTM5YzU1Y2RhMmRiIiwiYXV0aF90aW1lIjoxNTA2MzY5Njk5LCJpZHAiOiJmaGljdC1zc28iLCJyb2xlIjpbInVzZXIiLCJzdHVkZW50Il0sInVwbiI6IkkzNTg3MjJAZmhpY3QubmwiLCJuYW1lIjoiVGFybmV2LEhyaXN0aXlhbiBILkUuIiwiZW1haWwiOiJocmlzdGl5YW4udGFybmV2QHN0dWRlbnQuZm9udHlzLm5sIiwidXJuOm5sLmZoaWN0OnNjaGVkdWxlIjoiY2xhc3N8RGVsdGEgLyBFaTNTMSAvIFNNNDEiLCJmb250eXNfdXBuIjoiMzU4NzIyQHN0dWRlbnQuZm9udHlzLm5sIiwiYW1yIjpbImV4dGVybmFsIl19.fhAeEQ-40D5tGlNeF6slP9p8gQGFn3LQ8QOw1VCj8dlLjP41o3hQvXA4uO9aabQLaWkUxi3jnHce7yZhTMNaE2LDxJyIeSv1erElNBlNta4AXB_tcbE-H57yC3P2B0sP98PQrZUX_Uy7VtJbXwudWd6cFBhcJg9yEeZKUz3wszKqYaqx09ukxPo7taxAmVZKhqss_dI90jW0GHVMQEAy19hYbrWw2nWT_YFf-2fFjAkc4bbTuh4TI-j7PiZOOJCvkz7jA5gUvvYS2hmOfje4jHAq0uaGfgzoM93v27qQKmGnNlRfDC3RHWfK6K7rMWc9b-tYU0UYmNHO_Fuz4VcX0w"

var url = "https://api.fhict.nl/location/devices"

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

func getInfo(t time.Time) {
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	resp, err := oauthClient.Get(url)
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

	doEvery(20000*time.Millisecond, getInfo)

}
