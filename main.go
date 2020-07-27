package main

import (
	"ether/repl"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
)

var (
	interactive bool
	version     bool
)

// Init flag
func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] [<filename>]\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.BoolVar(&version, "v", false, "display version information")

	flag.BoolVar(&interactive, "i", false, "enable interactive mode")
}

// Launch main program
func main() {
	flag.Parse()

	if version {
		fmt.Printf("Ether %s \n", FullVersion())
		os.Exit(0)
	}

	user, err := user.Current()
	if err != nil {
		log.Fatalf("could not determine current user: %s", err)
	}
	args := flag.Args()
	opts := &repl.Options{
		Interactive: interactive,
	}
	repl := repl.New(user.Username, args, opts)
	repl.Run()
}
