//
// basic.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package auth

import (
	"encoding/base64"
	"net/url"
	"strings"
)

func BasicAuth(username, password string) string {
	data := url.QueryEscape(username) + ":" + url.QueryEscape(password)
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func ParseBasicAuth(authorization string) (username, password string, ok bool) {
	decoded, err := base64.StdEncoding.DecodeString(authorization)
	if err != nil {
		return
	}
	str := string(decoded)
	idx := strings.IndexByte(str, ':')
	if idx < 0 {
		return
	}
	u, err := url.QueryUnescape(str[:idx])
	if err != nil {
		return
	}
	p, err := url.QueryUnescape(str[idx+1:])
	if err != nil {
		return
	}
	return u, p, true
}
