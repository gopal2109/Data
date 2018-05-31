package cli

import "fmt"

type Command interface {
	Run(...string)
	Usage()
}

type Cmdfns map[string]Command

type Cli struct {
	Commands Cmdfns
}

func (cli Cli) Command(cmd string, args ...string) {
	for _, a := range args {
		if a == "help" || a == "--help" || a == "-help" {
			for c, _ := range cli.Commands {
				if c == cmd {
					(cli.Commands[c]).Usage()
					return
				}
			}
		}
	}

	for c, _ := range cli.Commands {
		if c == cmd {
			(cli.Commands[c]).Run(args...)
			return
		}
	}
	fmt.Printf("Command Not Found")
}

func NewCli() Cli {
	cf := make(Cmdfns, 0)

	cf["loaddb"] = LoadDB{}
	cf["genconfig"] = GenConfig{}
	cf["initthresholds"] = InitThresholds{}

	return Cli{Commands: cf}
}

	
	
