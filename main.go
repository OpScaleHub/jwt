package main

import (
	"bufio"
	"flag"
	"jwt/jwtExpose"
	"log"
	"os"
	"strings"
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

	jwtExpose.Expose(tokenString)
}
