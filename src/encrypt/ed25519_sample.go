package encrypt

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func parsePrivateKey(apiSecret string) (crypto.PrivateKey, error) {
	block, _ := pem.Decode([]byte(apiSecret))
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func signPayload(payload string, privateKey crypto.PrivateKey) (string, error) {
	key, ok := privateKey.(crypto.Signer)
	if !ok {
		return "", fmt.Errorf("key does not implement crypto.Signer")
	}

	signature, err := key.Sign(nil, []byte(payload), crypto.Hash(0))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func SampleMain() {

	// ed25519 key
	skey := "-----BEGIN PRIVATE KEY-----\n\n-----END PRIVATE KEY-----"
	payload := "apiKey=&timestamp=1746605972672"
	parseKey, _ := parsePrivateKey(skey)
	sign, _ := signPayload(payload, parseKey)
	fmt.Printf("sign:%s\n", sign)
}
