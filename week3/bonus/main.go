package main

import (
	"flag"
	"fmt"
)

func xorEncryptDecrypt(text, key string) []byte {
	result := make([]byte, len(text))
	for i := range text {
		result[i] = text[i] ^ key[i%len(key)]
	}
	return result
}

func main() {
	plaintext := flag.String("t", "Slizik", "text to cypher")
	key := flag.String("k", "key", "key used")
	flag.Parse()

	encrypted := xorEncryptDecrypt(*plaintext, *key)
	fmt.Printf("XOR Encrypted (hex): %x\n", encrypted)

	decrypted := xorEncryptDecrypt(string(encrypted), *key)
	fmt.Println("XOR Decrypted:", string(decrypted))
}
