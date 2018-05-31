package cli

import (
	//"rpis/config"
	"fmt"
)

type LoadDB struct {}

func (l LoadDB) Run(args ...string) {
	fmt.Print(args)
	fmt.Print("Hey You ran load db")
}

func (l LoadDB) Usage() {
	fmt.Printf("rpis [OPTIONS] loaddb [help]")
}
