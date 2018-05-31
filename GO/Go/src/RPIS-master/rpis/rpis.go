package main

import (
	"flag"
	"fmt"
	"os"
	"rpis/api"
	"rpis/backend"
	commands "rpis/cli"
	"rpis/config"
	"rpis/daemon"
	"strings"
)

var (
	version     *bool
	config_file *string
	stdout      = os.Stdout
	cli         commands.Cli
	log         = backend.Log
)

const (
	buildVersion = "1.0.0"
)

func DisplayHeader() {
	fmt.Fprintf(stdout, "Rackspace Inc.\n\n")
}

func main() {
	version = flag.Bool("version", false, "--version for current release version")
	config_file = flag.String("config", "./rpis.conf", "--config: configuration path, defaults to current directory")

	flag.Parse()

	if *version {
		fmt.Fprintf(stdout, "Current version is %s", buildVersion)
		return
	}

	cli = commands.NewCli()
	
	err, conf := config.LoadConfig(*config_file)
	if err != nil {
		fmt.Println(err.Error(), "\n", "Run 'rpis genconfig rpis.conf'")
	}
	
	config.Conf = conf

	flag.Usage = func() {
		DisplayHeader()
		fmt.Fprintf(stdout, "Usage: rpis [[OPTIONS] COMMANDS [args..]]\n    rpis [ --help | --version]\n\n")
		flag.CommandLine.SetOutput(stdout)
		flag.PrintDefaults()

		fmt.Fprintf(stdout, "Avaliable Commands:\n")
		for k, _ := range cli.Commands {
			fmt.Fprintf(stdout, "\t"+k+"\n")
		}

	}

	/* Commands if we got any */
	if flag.NArg() > 0 {
		cmd := flag.Args()[0]
		if cmd == "help" || cmd == "-help" || cmd == "--help" {
			flag.Usage()
			return
		}
		rpisCli := commands.NewCli()
		rpisCli.Command(strings.ToLower(cmd), flag.Args()[1:]...)
		return
	}

	log.Info("Starting Daemons")

	daemons := daemon.NewDaemon()
	daemons.Start()

	log.Info("Starting API Server")
	api.APIServer()

}
