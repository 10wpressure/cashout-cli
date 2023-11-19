package main

import (
	"crypto/ed25519"
	"encoding/hex"
)

func toUint8Array(str string) []byte {
	bytes, _ := hex.DecodeString(str)
	return bytes
}

func checkSignature(signature string, publicKey string, message string) bool {
	sigBytes := toUint8Array(signature)
	pubKeyBytes := toUint8Array(publicKey)
	msgBytes := toUint8Array(message)

	return ed25519.Verify(pubKeyBytes, msgBytes, sigBytes)
}
