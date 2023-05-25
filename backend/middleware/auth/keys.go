package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const keysFile = "keys.json"

type Keys struct {
	ApiKeys []string `json:"api_keys"`
}

func GenerateAndSaveKeys(numberOfKeys int) {
	keys := Keys{}
	for i := 0; i < numberOfKeys; i++ {
		keys.ApiKeys = append(keys.ApiKeys, GenerateKey())
	}
	keysJson, _ := json.Marshal(keys)
	err := ioutil.WriteFile(keysFile, keysJson, 0644)
	if err != nil {
		log.Fatalf("Failed to write keys to file: %v", err)
	}

	log.Println("Generated API keys:")
	for _, key := range keys.ApiKeys {
		log.Println(key)
	}
	log.Println("This is your ONLY chance to back up these keys. Do not lose them!")
}

func LoadKeys() []string {
	file, err := ioutil.ReadFile(keysFile)
	if err != nil {
		log.Fatalf("Failed to read keys from file: %v", err)
	}

	var keys Keys
	err = json.Unmarshal(file, &keys)
	if err != nil {
		log.Fatalf("Failed to unmarshal keys: %v", err)
	}
	return keys.ApiKeys
}
