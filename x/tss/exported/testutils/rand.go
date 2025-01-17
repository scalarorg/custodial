package testutils

import (
	"github.com/scalarorg/scalar-core/testutils/rand"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
)

// RandKeyID creates a random valid key ID
func RandKeyID() tss.KeyID {
	keyID := tss.KeyID(rand.NormalizedStrBetween(tss.KeyIDLengthMin, tss.KeyIDLengthMax+1))
	if err := keyID.Validate(); err != nil {
		panic(err)
	}
	return keyID
}
