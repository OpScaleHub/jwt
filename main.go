package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	tokenString := os.Args[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// Print Headers
		fmt.Println("      [#] HEADER:ALGORITHM & TOKEN TYPE")

		if _, ok := token.Header["alg"]; ok {
			fmt.Printf("alg:   [ %s ]\n", token.Header["alg"])
		}

		if _, ok := token.Header["kid"]; ok {
			fmt.Printf("kid:   [ %s ]\n", token.Header["kid"])
		}

		if _, ok := token.Header["typ"]; ok {
			fmt.Printf("typ:   [ %s ]\n", token.Header["typ"])
		}

		// Print Token DATA
		fmt.Println("      [#] PAYLOAD:DATA")
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
		// Verify SigSignature
		fmt.Println("      [#] VERIFY SIGNATURE")
		//fmt.Printf("Signature Verified\n")
		fmt.Printf("#ToDo\n")

		// for k, v := range claims {
		// 	if k == "exp" || k == "iat" {
		// 		epochTime := int64(v.(float64))
		// 		fmt.Println(time.Unix(epochTime, 0))
		// 		fmt.Printf("%s: [%d]\n", k, epochTime)
		// 	} else {
		// 		fmt.Printf("%s: [%s]\n", k, v)
		// 	}

		// }
	}
}
