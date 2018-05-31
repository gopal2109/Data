package cli

type Command interface {
	Run(...string)
	Usage()
}

type Cmdfns map[string]Command

type Cli struct {
	commands Cmdfns
}

func (cli Cli) Command(cmd string, args ...string) {
	for _, a := range args {
		if a == "help" || a == "--help" || a == "-help" {
			for c, _ := range cli.commands {
				if c == cmd {
					(cli.commands[c]).Usage()
					return
				}
			}
		}
	}
	
	for c, _ := range cli.commands {
		if c == cmd {
			(cli.commands[c]).Run(args...)
			return
		}
	}
	panic("No Command Found")
}

func NewCli() Cli {
	cf := make(Cmdfns)

	cf["loaddb"] = LoadDB{}

	return Cli{commands: cf}
}

	
	
