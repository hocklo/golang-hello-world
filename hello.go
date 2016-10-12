package main

import (
	"fmt"
	"github.com/hocklo/stringutil"
)

func main() {
	fmt.Printf("hello, world\n")
	// System output using the library stringutil
	fmt.Printf(stringutil.Reverse("!oG ,olleH")+"\n")
}

