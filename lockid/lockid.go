// Package lockid helps generate a unique id from strings supplied.
package lockid

import (
	"crypto/sha256"
	"errors"
	"strings"

	"github.com/google/uuid"
)

//GenerateLockID Generates a UUID from the identifier
func GenerateLockID(identifier string) (string, error) {

	if len(strings.TrimSpace(identifier)) == 0 {
		return "", errors.New("Invalid identifier")
	}

	hashsum := sha256.Sum256([]byte(identifier))
	lockid, err := uuid.FromBytes(hashsum[0:16])
	if err != nil {
		//we should never reach this, but ¯\_(ツ)_/¯
		return "", err
	}
	return lockid.String(), nil
}
