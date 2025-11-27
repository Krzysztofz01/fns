package main

import (
	"os"

	"github.com/Krzysztofz01/fns/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
