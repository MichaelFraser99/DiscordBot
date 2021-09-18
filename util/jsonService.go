package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJson(file string) []byte {
	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}
