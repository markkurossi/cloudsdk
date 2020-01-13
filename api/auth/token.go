//
// token.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package auth

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/ed25519"
)

func VerifyToken(token []byte, pub ed25519.PublicKey) bool {
	return false
}

func Authorize(w http.ResponseWriter, r *http.Request, realm string,
	tokenVerifier func(token string) bool) bool {

	auth := r.Header.Get("Authorization")
	if len(auth) == 0 {
		Error401(w, realm)
		return false
	}
	// XXX

	return true
}

func Error401(w http.ResponseWriter, realm string) {
	w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Bearer realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}
