package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"rpis/config"
	"rpis/daemon"
	commands "rpis/cli"
	"rpis/api"
)

var (
	version *bool
	config_file *string
	stdout = os.Stdout
	cli commands.Cli
)

const (
	buildVersion = "1.0.0"
)

func DisplayHeader() {
	fmt.Println("Rackspace Inc.")
}


func main() {
	version = flag.Bool("version", false, "--version for current release version")
	config_file = flag.String("config", "./rpis.conf", "--config: configuration path, defaults to current directory")

	flag.Parse()
	
	if *version {
		fmt.Fprintf(stdout, "Current version is %s", buildVersion)
		return
	}

	conf := config.LoadConfig(*config_file)

	config.Conf = conf

	cli = commands.NewCli()

	flag.Usage = func() {
		DisplayHeader()
		fmt.Fprintf(stdout, "Usgae: rpis [[OPTIONS] COMMANDS [args..]]\n    rpis [ --help | --version]\n\n")
		flag.CommandLine.SetOutput(stdout)
		flag.PrintDefaults()
		
	}

	/* Commands if we got any */
	if flag.NArg() > 0 {
		cmd := flag.Args()[0]
		
		rpisCli := commands.NewCli()
		rpisCli.Command(strings.ToLower(cmd), flag.Args()[1:]...)
		return
	}

	fmt.Println("Starting Daemons")

	daemons := daemon.NewDaemon()
	daemons.Start()

	fmt.Println("Starting API Server")
	api.APIServer()
	
}
