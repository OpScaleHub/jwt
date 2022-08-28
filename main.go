package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {

	flag.Parse()
	var tokenString string

	if flag.NArg() == 0 { // from stdin/pipe
		reader := bufio.NewReader(os.Stdin)
		var err error
		tokenString, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalln("failed to read input")
		}
		tokenString = strings.TrimSpace(tokenString) // otherwise, we would have a blank line
	} else { // from argument
		if flag.NArg() > 1 {
			log.Fatalln("takes at most one input")
		}
		tokenString = os.Args[1]
	}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println("---------------------------------")
		// Print Headers
		fmt.Println("[#] HEADER:ALGORITHM & TOKEN TYPE")
		if _, ok := token.Header["alg"]; ok {
			fmt.Printf("alg:   [ %s ]\n", token.Header["alg"])
		}

		if _, ok := token.Header["kid"]; ok {
			fmt.Printf("kid:   [ %s ]\n", token.Header["kid"])
		}

		if _, ok := token.Header["typ"]; ok {
			fmt.Printf("typ:   [ %s ]\n", token.Header["typ"])
		}
		fmt.Println("\n---------------------------------")
		// Print Token DATA
		fmt.Println("[#] PAYLOAD:DATA")
		if _, ok := claims["iss"]; ok {
			fmt.Printf("iss:   [ %s ]\n", claims["iss"])
		}
		decodedSub, err := base64.StdEncoding.DecodeString(claims["sub"].(string))
		if err != nil {
			fmt.Printf("sub:   [ %s ]\n", claims["sub"])
		} else {

			fmt.Printf("sub:   [ %s ]\n", decodedSub[1:])
		}

		if _, ok := claims["aud"]; ok {
			fmt.Printf("aud:   [ %s ]\n", claims["aud"])
		}

		if _, ok := claims["exp"]; ok {
			fmt.Printf("exp:   [ %s ]\n", time.Unix(int64(claims["exp"].(float64)), 0))
		}
		if _, ok := claims["iat"]; ok {
			fmt.Printf("iat:   [ %s ]\n", time.Unix(int64(claims["iat"].(float64)), 0))
		}

		if _, ok := claims["nonce"]; ok {
			fmt.Printf("nonce: [ %s ]\n", claims["nonce"])
		}
		fmt.Println("\n----------------")
		// Verify SigSignature
		fmt.Println("[#] VERIFY SIGNATURE")
		if ok := token.Valid; ok {
			fmt.Printf("Signature Verified.\n")
		} else {
			fmt.Printf("Invalid Signature.\n")
		}
		fmt.Println("--------------------")

	}
}
