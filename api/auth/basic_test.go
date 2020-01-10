//
// basic.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package auth

import (
	"testing"
)

func TestBasic(t *testing.T) {
	username := "mtr"
	password := "passw0rd"

	val := BasicAuth(username, password)
	u, p, ok := ParseBasicAuth(val)
	if !ok {
		t.Fatalf("Basic authentication encode/decode failed")
	}
	if u != username {
		t.Fatalf("Username mismatch")
	}
	if p != password {
		t.Fatalf("Password mismatch")
	}
}
