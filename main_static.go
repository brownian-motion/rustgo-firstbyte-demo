package main

import "fmt"
import . "github.com/brownian-motion/rustgo-firstbyte-demo/internal"

func main() {
	PrintFirstByte("Some letters")
	PrintFirstByte("")
}

func PrintFirstByte(s string) {
	b := FirstByte(s)

	fmt.Printf("The first byte of \"%s\" is '%s' (%x)\n", s, string(b), byte(b))
}
