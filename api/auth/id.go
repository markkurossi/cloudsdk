//
// id.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

type ID [8]byte

func NewID() (ID, error) {
	var id ID

	_, err := rand.Read(id[:])

	return id, err
}

func (id ID) String() string {
	return base64.RawURLEncoding.EncodeToString(id[:])
}

func ParseID(val string) (ID, error) {
	var id ID

	data, err := base64.RawURLEncoding.DecodeString(val)
	if err != nil {
		return id, err
	}
	if len(data) != len(id) {
		return id, fmt.Errorf("Invalid ID '%s'", val)
	}
	copy(id[:], data)

	return id, nil
}
