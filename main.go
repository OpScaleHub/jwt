package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/exp/slices"
)

func main() {

	// Parse input Args
	flag.Parse()
	var tokenString string

	// from stdin/pipe
	if flag.NArg() == 0 {
		reader := bufio.NewReader(os.Stdin)
		var err error
		tokenString, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalln("failed to read input")
		}
		// Clean unexpected spaces before and after inputs
		tokenString = strings.TrimSpace(tokenString)

	} else {
		if flag.NArg() > 1 { // Get more that 1 and only 1 arg as input with is not as expected
			log.Fatalln("takes at most one input")
		}
		tokenString = os.Args[1]
	}

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
