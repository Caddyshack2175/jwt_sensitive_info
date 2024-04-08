package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var encodedString = os.Args[1]
	jwt := JWTParts(encodedString)
	for _, j := range jwt {
		if strings.Contains(strings.ToLower(j), "name") || strings.Contains(strings.ToLower(j), "email") || strings.Contains(strings.ToLower(j), "user") || strings.Contains(strings.ToLower(j), "usr") {
			fmt.Printf("\n\nA JSON Web Token (JWT) should not contain sensitive information. JWTs are encoded using Base64, which is not an encryption method, meaning anyone can decode and read the information within the token. Encryption is not a built-in feature of JWT, but you can use another standard called JSON Web Encryption (JWE) to encrypt the data in JWT. This provides an extra layer of security to protect sensitive data. The following JWT Token was found to contain sensitive information:\n\n...\n%s\n...\n\n", j)
		}
	}
}

func JWTParts(encodedString string) []string {
	jwtparts := strings.Split(encodedString, ".")
	bdata := []string{}
	for _, jwt := range jwtparts[:2] {
		decodedBytes, _ := base64.StdEncoding.DecodeString(jwt)
		decodedString := string(decodedBytes) + "}"
		p := map[string]any{}
		_ = json.Unmarshal([]byte(decodedString), &p)
		byteData, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)

		}
		bdata = append(bdata, string(byteData))
	}
	return bdata
}
