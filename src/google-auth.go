package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func getClient(googleAuthConf googleAuth) (*http.Client, error) {
	jsonKey, err := ioutil.ReadFile(googleAuthConf.CredentialsPath)
	if err != nil {
		return &http.Client{}, err
	}

	config, err := google.ConfigFromJSON(jsonKey, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return &http.Client{}, err
	}

	tok, err := tokenFromFile(googleAuthConf.TokenPath)
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return &http.Client{}, err
		}
		if err = saveToken(googleAuthConf.TokenPath, tok); err != nil {
			return &http.Client{}, err
		}
	}
	return config.Client(context.Background(), tok), nil
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Println("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return &oauth2.Token{}, err
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return &oauth2.Token{}, err
	}
	return tok, nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return &oauth2.Token{}, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)

	return nil
}
