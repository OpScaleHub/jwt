package jwtExpose

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/exp/slices"
)

// create add function
func Expose(tokenString string) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		fmt.Println("[#] HEADER:ALGORITHM & TOKEN TYPE")
		//fmt.Println(token.Header)
		for k, v := range token.Header {
			fmt.Printf("%s: [%s]\n", k, v)
		}
		fmt.Println("\n----------------")
		fmt.Println("[#] PAYLOAD:DATA")
		//fmt.Println(token.Claims)
		for k, v := range claims {
			timestamps := []string{"exp", "iat", "nbf", "naf"}
			if slices.Contains(timestamps, k) {
				fmt.Printf("%s: [ %s ]\n", k, time.Unix(int64(v.(float64)), 0))
			} else {
				fmt.Printf("%s: [ %s ]\n", k, v)
			}

		}
		fmt.Println("\n----------------")
		// Verify SigSignature
		fmt.Println("[#] VERIFY SIGNATURE")
		if ok := token.Valid; ok {
			fmt.Printf("Signature Verified.\n")
			fmt.Println(token.Signature)
		} else {
			fmt.Printf("Invalid Signature.\n")
		}
	}
}
