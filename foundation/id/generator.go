package id

import (
	"encoding/hex"
	"fmt"
	"os"
)

func GenerateID(length int) (string, error) {
	b, err := byteRand(length)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), err
}

func GenerateUUID() (string, error) {
	b, err := byteRand(16)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

func byteRand(length int) ([]byte, error) {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		return nil, err
	}
	b := make([]byte, length)
	_, err = f.Read(b)
	if err != nil {
		return nil, err
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	return b, nil
}
