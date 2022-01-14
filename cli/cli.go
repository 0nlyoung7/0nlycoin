package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/0nlyoung7/0nlycoin/explorer"
	"github.com/0nlyoung7/0nlycoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 0nly coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000: Set the PORT of the server\n")
	fmt.Printf("-mode=rest: Choose between 'html' and 'rest'\n")
	os.Exit(0)
}

func Start() {

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
