package main

import (
	"fmt"
	"github.com/thejunghare/searchfiles/cmd"
)

func main() {
	checkError(cmd.rootCmd.Execute())
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Something went wrong", err)
	}
}
