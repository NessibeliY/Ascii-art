package internal

import (
	"errors"
	"hash/crc64"
	"os"
)

func ValidInput(input string) (bool, error) {
	for _, letter := range input {
		if letter > 127 {
			return false, errors.New("provide ascii chars only")
		}
	}

	return true, nil
}

func getInput() (string, error) {
	args := os.Args[1:]

	if len(args) != 1 {
		return "", errors.New("please provide one string")
	}

	return args[0], nil
}

func strToHash(bannerText []byte) uint64 {
	crc64Table := crc64.MakeTable(crc64.ECMA)
	hashedData := crc64.Checksum([]byte(bannerText), crc64Table)
	return hashedData
}
