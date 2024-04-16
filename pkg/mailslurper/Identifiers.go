// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailslurper

import (
	"github.com/google/uuid"
)

/*
GenerateID creates a UUID ID for database records.
*/
func GenerateID() (string, error) {
	id, err := uuid.NewRandom()
	return id.String(), err
}
