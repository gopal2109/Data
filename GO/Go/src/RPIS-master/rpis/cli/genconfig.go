package cli

import (
	"fmt"
	"os"
)

var sample_conf = `
# RPIS Sample config

[mongo]
host='localhost'
name='nis'

[http]
hostaddress=':9000'

[logging]
file="-" # Use "-" for stdout
level="debug" # debug, warn, error, fatal, panic

[application]
timezone="US/Eastern"
dateformat="2006-01-01"
`

type GenConfig struct {}

func (gc GenConfig) Run(args ...string) {
	filename := "rpis.conf"
	
	if len(args) > 0 {
		filename = args[0]
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	
	_, err = file.WriteString(sample_conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Created Successfully")
}

func (gc GenConfig) Usage() {
	fmt.Printf("rpis genconfig [path_to_file_to_be_written]")
}
