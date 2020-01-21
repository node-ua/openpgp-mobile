package openpgp

import (
	"encoding/json"
	"testing"
)

func TestFastOpenPGP_EncryptMultipleSign(t *testing.T) {
	options := &Options{
		Email:      "sample@sample.com",
		Name:       "Test2",
		Comment:    "sample",
		Passphrase: "123123",
		KeyOptions: &KeyOptions{
			CompressionLevel: 9,
			RSABits:          2048,
			Cipher:           "aes256",
			Compression:      "zlib",
			Hash:             "sha512",
		},
	}

	openPGP := NewFastOpenPGP()
	_output2, _ := openPGP.Generate(options)
	_output3, _ := openPGP.Generate(options)

	publicKey2 := _output2.PublicKey
	publicKey3 := _output3.PublicKey
	keys := []string{publicKey2, publicKey3}
	k, _ := json.Marshal(keys)
	output, err := openPGP.EncryptWithMultipleKeys(inputMessage, string(k))
	if err != nil {
		t.Fatal(err)
	}

	t.Log("output:", output)
}
