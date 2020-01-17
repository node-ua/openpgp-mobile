package openpgp

import "testing"

func TestFastOpenPGP_GetKeyId(t *testing.T) {

	openPGP := NewFastOpenPGP()
	output, err := openPGP.GetKeyId(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("output:", output)
}
