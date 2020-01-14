//
// oauth2.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type OAuth2Client struct {
	id            string
	secret        string
	TokenEndpoint string
	http          *http.Client
}

type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func NewOAuth2Client(id, secret, token string) *OAuth2Client {
	return &OAuth2Client{
		id:            id,
		secret:        secret,
		TokenEndpoint: token,
		http:          new(http.Client),
	}
}

func (client *OAuth2Client) GetToken() (*TokenResponse, error) {
	resp, err := client.http.PostForm(client.TokenEndpoint, url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {client.id},
		"client_secret": {client.secret},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	token := new(TokenResponse)
	err = json.Unmarshal(body, token)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`error="%s", error_description="%s"`,
			token.Error, token.ErrorDescription)
	}

	return token, nil
}
