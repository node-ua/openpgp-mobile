package openpgp

import (
	"strings"

	"github.com/keybase/go-crypto/openpgp"
)

func (o *FastOpenPGP) GetKeyId(privateKey string) (string, error) {
	rings, err := openpgp.ReadArmoredKeyRing(strings.NewReader(privateKey))
	if err != nil {
		return "", err
	}

	return rings[0].PrivateKey.KeyIdString(), nil
}
