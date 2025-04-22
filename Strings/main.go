package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// This code is about to take an argument from the user given input(first is the old string, and then swap it
	// to the new string (e.g our text have string cook, then our input is cook eat, so the string 'cook' will swap by eat.
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough ergs")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}

}
