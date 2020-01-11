//
// token.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package auth

import (
	"github.com/markkurossi/go-libs/tlv"
	"golang.org/x/crypto/ed25519"
)

const (
	TOKEN_VALUES tlv.Type = iota
	TOKEN_SIGNATURE
	T_TENANT_ID
	T_CLIENT_ID
	T_CREATED
	T_SCOPE
)

const (
	SCOPE_ADMIN tlv.Type = iota
)

func VerifyToken(token []byte, pub ed25519.PublicKey) bool {
	return false
}
