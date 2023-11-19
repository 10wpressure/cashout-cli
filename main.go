package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"strings"
)

var privateKey, publicKey, messages string

func init() {
	loadKeysFromEnv()
	loadKeysFromUserInput()
}

func main() {
	m := strings.Split(messages, ",")
	var result []string
	for _, msg := range m {
		sig := ed25519.Sign(toUint8Array(privateKey), toUint8Array(msg))
		signatureStr := hex.EncodeToString(sig)
		//checked := checkSignature(signatureStr, publicKey, msg)
		//fmt.Println("checked: ", checked)
		result = append(result, signatureStr)
	}

	fmt.Println(`"x-signatures":"` + strings.Join(result, "") + `"`)
}
