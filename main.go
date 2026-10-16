package main

import (

	"flag"
	"fmt"
	"os"
	"net/url"
	"context"
	"time"
	"github.com/amad3us47/veylith/internal/engine"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Println("Usage: veylith <command>")
		fmt.Println("Commands: scan")
		os.Exit(1)
	}

	switch os.Args[1] {
	
		case "scan":
			scanCommand()
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
			os.Exit(1)
	}
}

func scanCommand(){

	flags := flag.NewFlagSet("scan",flag.ExitOnError)
	target := flags.String("u","", "Target URL")
	
	flags.Parse(os.Args[2:])

	if *target == "" {
	
		fmt.Println("Usage: veylith scan -u <URL>")
		os.Exit(1)
	}
	parsed, err := url.ParseRequestURI(*target)
	if err != nil {
	
		fmt.Println("Invalid URL")

	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
	
		fmt.Println("Invalid URL: scheme must be http or https")
		os.Exit(1)
	}

	if parsed.Host == "" {
	
		fmt.Println("Invalid URL: host is required")
		os.Exit(1)
	}

	fmt.Printf("Target received: %s\n", *target)
}

