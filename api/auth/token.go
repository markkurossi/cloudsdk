//
// token.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package auth

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/markkurossi/go-libs/tlv"
)

func VerifyToken(token []byte, pub ed25519.PublicKey) bool {
	return false
}

func Authorize(w http.ResponseWriter, r *http.Request, realm string,
	signatureVerifier func(message, sig []byte) bool,
	tenant *Tenant) tlv.Values {

	var tokenB64 string

	auth := r.Header.Get("Authorization")
	if len(auth) > 0 {
		parts := strings.Split(auth, " ")
		if len(parts) != 2 {
			Error401f(w, realm, "invalid_authorization",
				"Invalid HTTP Authorization header")
			return nil
		}
		if parts[0] != "Bearer" {
			Error401f(w, realm, "invalid_authorization",
				"Bearer authorization required")
			return nil
		}
		tokenB64 = parts[1]
	} else {
		tokens, ok := r.URL.Query()["token"]
		if ok && len(tokens) > 0 {
			tokenB64 = tokens[0]
		} else {
			Error401(w, realm)
			return nil
		}
	}
	data, err := base64.RawURLEncoding.DecodeString(tokenB64)
	if err != nil {
		Error401f(w, realm, "invalid_token", "Base64 decoding failed: %s", err)
		return nil
	}
	token, err := tlv.Unmarshal(data)
	if err != nil {
		Error401f(w, realm, "invalid_token", "TLV decoding failed: %s", err)
		return nil
	}
	values, ok := token[TOKEN_VALUES].(tlv.Values)
	if !ok {
		Error401f(w, realm, "invalid_token", "No token values")
		return nil
	}
	signature, ok := token[TOKEN_SIGNATURE].([]byte)
	if !ok {
		Error401f(w, realm, "invalid_token", "No token signature")
		return nil
	}
	valuesData, err := values.Marshal()
	if err != nil {
		Error401f(w, realm, "invalid_token", "Values marshal failed: %s", err)
		return nil
	}
	if !signatureVerifier(valuesData, signature) {
		Error401f(w, realm, "invalid_token", "Signature verification failed")
		return nil
	}

	tenantID, ok := values[T_TENANT_ID].(string)
	if !ok {
		Error401f(w, realm, "invalid_token", "No tenant ID")
		return nil
	}
	if tenant != nil && tenantID != tenant.ID {
		Error401f(w, realm, "invalid_token", "Wrong authentication tenant")
		return nil
	}

	// XXX expiration

	return values
}

func Error401(w http.ResponseWriter, realm string) {
	w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Bearer realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}

func Error401f(w http.ResponseWriter, realm, error, desc string,
	a ...interface{}) {

	description := fmt.Sprintf(desc, a...)
	w.Header().Set("WWW-Authenticate",
		fmt.Sprintf(`Bearer realm="%s", error="%s", error_description="%s"`,
			realm, error, description))
	w.WriteHeader(http.StatusUnauthorized)
}
