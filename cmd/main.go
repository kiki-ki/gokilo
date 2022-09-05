package main

import (
	"flag"
	"fmt"
	"gokilo/editor"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: <filepath>")
		return
	}

	if err := run(flag.Arg(0)); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

func run(filePath string) error {
	e, err := editor.NewEditor(filePath)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", e)

	defer e.Termios.DisableRawMode()

	return nil
}
