// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var JWTIssuer string = "mailslurper"
var ErrInvalidToken error = fmt.Errorf("Invalid token")
var ErrTokenMissingClaims error = fmt.Errorf("Token is missing claims")
var ErrInvalidUser error = fmt.Errorf("Invalid user")
var ErrInvalidIssuer error = fmt.Errorf("Invalid issuer")

type Claims struct {
	jwt.RegisteredClaims `json:"claims"`
	User                 string `json:"user"`
}
