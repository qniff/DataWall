package main

import (
	"encoding/json"
	"fmt"
  "time"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
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
var personalAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6ImdyQWk2cnJRU0JiVVItY01ZOHpRTHE2aGdVQSIsImtpZCI6ImdyQWk2cnJRU0JiVVItY01ZOHpRTHE2aGdVQSJ9.eyJpc3MiOiJodHRwczovL2lkZW50aXR5LmZoaWN0Lm5sIiwiYXVkIjoiaHR0cHM6Ly9pZGVudGl0eS5maGljdC5ubC9yZXNvdXJjZXMiLCJleHAiOjE1MDY1MTY1ODMsIm5iZiI6MTUwNjUwOTM4MywiY2xpZW50X2lkIjoiYXBpLWNsaWVudCIsInVybjpubC5maGljdDp0cnVzdGVkX2NsaWVudCI6InRydWUiLCJzY29wZSI6WyJvcGVuaWQiLCJwcm9maWxlIiwiZW1haWwiLCJmaGljdCIsImZoaWN0X3BlcnNvbmFsIiwiZmhpY3RfbG9jYXRpb24iXSwic3ViIjoiNWRlY2Y0YTgtNDhiZC00N2VjLTk4YjQtYTM5YzU1Y2RhMmRiIiwiYXV0aF90aW1lIjoxNTA2NTA5MzgzLCJpZHAiOiJmaGljdC1zc28iLCJyb2xlIjpbInVzZXIiLCJzdHVkZW50Il0sInVwbiI6IkkzNTg3MjJAZmhpY3QubmwiLCJuYW1lIjoiVGFybmV2LEhyaXN0aXlhbiBILkUuIiwiZW1haWwiOiJocmlzdGl5YW4udGFybmV2QHN0dWRlbnQuZm9udHlzLm5sIiwidXJuOm5sLmZoaWN0OnNjaGVkdWxlIjoiY2xhc3N8RGVsdGEgLyBFaTNTMSAvIFNNNDEiLCJmb250eXNfdXBuIjoiMzU4NzIyQHN0dWRlbnQuZm9udHlzLm5sIiwiYW1yIjpbImV4dGVybmFsIl19.MnxzAWV0aPFRke9rUqvQvJLCwF0-15952B9JXym26gB0v-vOglGuHSeWofOCfoHLMf291MVRQRj0WC9YUB0oPrqMHIk9lOuZFkJobwC682qTLUwSfK-171AifTuOW5itjggYEtKwZrZIG97kdAEwtqN_tGmFXTIqH_s4AneHCQDioJk4yZBNt9Mtb4BxvmfwwwKjGHUVvCvg7hsAav1QyIMnoqY3a8bLRwSOMJvAsbf_WC3HGma4APYfW8jS_DoZUqOLmZ6pRJgBcLbZYMFm9UGkaC1LUhndqk4Zr5YSEC451cZdUjgAdOLM_0_nSFIlsdf2po5JbYdi9n27chberQ"




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
func serve(w http.ResponseWriter, h *http.Request){
	
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	resp, _ := oauth2.NewClient(oauth2.NoContext, tokenSource).Get(url)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(body))
}
func getInfo(t time.Time){
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
// func doFancyStuff(t time.Time){
// 	http.HandleFunc("/", serve)
// 	http.ListenAndServe(":8000",nil)
	
// }

func main() {
  http.HandleFunc("/", serve)
	http.ListenAndServe(":8000",nil)
	//getInfo(time.Now());
	//doEvery(20000*time.Millisecond, getInfo)
	//http.HandleFunc("/", serve)
	//doFancyStuff(time.Now())
	//doEvery(20000*time.Millisecond, doFancyStuff)
	//http.ListenAndServe(":8000",nil)
	
	
}
