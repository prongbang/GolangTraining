package utils


import (
	"encoding/json"
	"log"
)

func Type2JsonString(inf interface{}) string {
	b, err := json.Marshal(inf)
	if err != nil {
		log.Fatal(b)
	}
	return string(b)
}

func Type2JsonByte(inf interface{}) []byte {
	b, err := json.Marshal(inf)
	if err != nil {
		log.Fatal(b)
	}
	return b
}
