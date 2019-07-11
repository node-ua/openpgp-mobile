package openpgp

import (
	"bytes"
	"golang.org/x/crypto/openpgp"
)

func (o *OpenPGP) Sign(message, publicKey, privateKey, passphrase string) (string, error) {

	entity, err := o.readSignKey(publicKey, privateKey, passphrase)
	if err != nil {
		return "", err
	}

	writer := new(bytes.Buffer)
	reader := bytes.NewReader([]byte(message))
	err = openpgp.ArmoredDetachSign(writer, entity, reader, nil)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}