package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetModFile(url string) GoModFile {

	// urls of the form: https://raw.githubusercontent.com/JoeLanglands/golenoid/main/go.mod

	resp, err := http.Get("https://raw.githubusercontent.com/JoeLanglands/golenoid/main/go.mod")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	return parseGoModFile(body)
}
