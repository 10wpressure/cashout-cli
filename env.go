package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getAnInputFromUser(t string) string {
	var key string
	fmt.Printf("Enter your %s key: ", t)
	_, err := fmt.Scan(&key)
	if err != nil {
		panic("could not a key: " + err.Error())
	}
	return key
}

func saveKeysToEnv(keys ...string) {
	envContent := ""
	var et envType
	for i, key := range keys {
		fmt.Println(key)
		et = envType(i)
		name := et.String()
		envContent += fmt.Sprintf("%s=%s\n", name, key)
	}
	err := os.WriteFile(".env", []byte(envContent), 0644)
	if err != nil {
		log.Fatal("Error saving keys to .env file:", err)
	}
}

func loadKeysFromEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
		createEmptyEnvFile()
	}
	privateKey = os.Getenv("PRIVATE_KEY")
	publicKey = os.Getenv("PUBLIC_KEY")
	messages = os.Getenv("MESSAGES")
}

func loadKeysFromUserInput() {
	if privateKey == "" {
		privateKey = getAnInputFromUser("private")
	}
	if publicKey == "" {
		publicKey = getAnInputFromUser("public")
	}
	if messages == "" {
		messages = getAnInputFromUser("messages")
	}

	saveKeysToEnv(privateKey, publicKey, messages)
}

func createEmptyEnvFile() {
	err := os.WriteFile(".env", []byte{}, 0644)
	if err != nil {
		log.Fatal("Error creating .env file:", err)
	}
}
