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
	TOKEN_VALUES tlv.Type = iota
	TOKEN_SIGNATURE
	T_TENANT_ID
	T_CLIENT_ID
	T_NOT_BEFORE
	T_NOT_AFTER
	T_SCOPE
)

const (
	SCOPE_ADMIN tlv.Type = iota
)

var (
	TokenSymtab = tlv.Symtab{
		TOKEN_VALUES: tlv.Symbol{
			Name: "values",
			Child: tlv.Symtab{
				T_TENANT_ID: tlv.Symbol{
					Name: "tenant_id",
				},
				T_CLIENT_ID: tlv.Symbol{
					Name: "client_id",
				},
				T_NOT_BEFORE: tlv.Symbol{
					Name: "not_before",
				},
				T_NOT_AFTER: tlv.Symbol{
					Name: "not_after",
				},
				T_SCOPE: tlv.Symbol{
					Name: "scope",
					Child: tlv.Symtab{
						SCOPE_ADMIN: tlv.Symbol{
							Name: "admin",
						},
					},
				},
			},
		},
		TOKEN_SIGNATURE: tlv.Symbol{
			Name: "signature",
		},
	}
)

const (
	KEY_CLIENT_ID_SECRET    = "client-id-secret"
	KEY_TOKEN_SIGNATURE_KEY = "token-signature-key"
)

const (
	ASSET_AUTH_PUBKEY = "auth-pubkey"
)
