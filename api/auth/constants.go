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
)

const (
	T_TENANT_ID tlv.Type = iota
	T_CLIENT_ID
	T_SCOPE
)

const (
	SCOPE_ADMIN tlv.Type = iota
)
