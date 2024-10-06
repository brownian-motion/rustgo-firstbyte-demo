package main

import "fmt"
import . "github.com/brownian-motion/rustgo-firstbyte-demo/internal"

func main() {
	print_first_byte("Some letters")
	print_first_byte("")
}

func print_first_byte(s string) {
	b := FirstByte(s)

	fmt.Printf("The first byte of \"%s\" is '%s' (%x)\n", s, string(b), byte(b))
}
