//
// token.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package auth

import (
	"golang.org/x/crypto/ed25519"
)

func VerifyToken(token []byte, pub ed25519.PublicKey) bool {
	return false
}
