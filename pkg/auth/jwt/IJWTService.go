// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type IJWTService interface {
	CreateToken(authSecret, user string) (string, error)
	DecryptToken(token string) (string, error)
	EncryptToken(token string) (string, error)
	GetUserFromToken(token *jwt.Token) string
	Parse(tokenFromHeader, authSecret string) (*jwt.Token, error)
	IsTokenValid(token *jwt.Token) error
}
