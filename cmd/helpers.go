package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

func checkTokenFlag() {
	if *tokenFlag == "" {
		fmt.Println("Token not set.")
		os.Exit(1)
	}
}

func checkKeyFlag() {
	if *keyFlag == "" {
		fmt.Println("Key not set.")
		os.Exit(1)
	}
}

func writeTOML(key, token string) error {
	var inputs = Config{
		key,
		token,
	}

	var buffer bytes.Buffer

	encoder := toml.NewEncoder(&buffer)

	err := encoder.Encode(inputs)

	if nil != err {
		return err
	}

	return ioutil.WriteFile("config.toml", buffer.Bytes(), 0644)
}

func readTOML(filename string) {
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
