package openpgp

import (
	"strings"

	"github.com/keybase/go-crypto/openpgp"
)

func (o *FastOpenPGP) GetKeyId(key string) (string, error) {
	rings, err := openpgp.ReadArmoredKeyRing(strings.NewReader(key))
	if err != nil {
		return "", err
	}

	return rings[0].PrimaryKey.KeyIdString(), nil
}
