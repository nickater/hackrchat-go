package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nickater/hackrchat-cli/src/model"
)

func LoginUser(credentials model.Credential) (string, error) {
	fmt.Println(credentials.Username)
	result, err := http.Get("https://catfact.ninja/fact")

	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()
	var catFactResult struct {
		Fact   string `json: "#fact"`
		Length int    `json: "#length"`
	}

	json.NewDecoder(result.Body).Decode(&catFactResult)
	//Convert the body to type string
	sb := fmt.Sprintf("Fact: %s, Length: %d", catFactResult.Fact, catFactResult.Length)

	return sb, nil
}

func RegisterUser(credentials model.Credential) (string, error) {
	fmt.Println(credentials.Username)
	result, err := http.Get("https://catfact.ninja/fact")

	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()
	var catFactResult struct {
		Fact   string `json: "#fact"`
		Length int    `json: "#length"`
	}

	json.NewDecoder(result.Body).Decode(&catFactResult)
	//Convert the body to type string
	sb := fmt.Sprintf("Fact: %s, Length: %d", catFactResult.Fact, catFactResult.Length)

	return sb, nil
}
