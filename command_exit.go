package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config , arg []string) error {
	if len(arg) != 0 {
		fmt.Println("Exit command doesn't take any arguments")
		return nil
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}