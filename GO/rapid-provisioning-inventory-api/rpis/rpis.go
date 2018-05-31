package main

import (
	"flag"
	"fmt"
	"os"
	"rpis/api"
	"rpis/config"

	log "github.com/Sirupsen/logrus"
)

var (
	version    *bool
	configFile *string
)

const (
	buildVersion = "v0.1"
)

func displayBanner() {
	fmt.Fprintf(os.Stdout, "Rackspace Inc.\n\n")
}

func main() {
	version = flag.Bool("version", false, "--version for current release version")
	configFile = flag.String("config", "./rpis.conf", "--config: configuration path, defaults to current directory")

	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "Current version is %s", buildVersion)
		return
	}

	config.Init(*configFile)
	conf := config.C()
	config.Watcher()

	flag.Usage = func() {
		displayBanner()
		fmt.Fprintf(os.Stdout, "Usage: rpis [[OPTIONS] COMMANDS [args..]]\n    rpis [ --help | --version]\n\n")
		flag.CommandLine.SetOutput(os.Stdout)
		flag.PrintDefaults()
	}

	/* Commands if we got any */
	if flag.NArg() > 0 {
		cmd := flag.Args()[0]
		if cmd == "help" || cmd == "-help" || cmd == "--help" {
			flag.Usage()
			return
		}
	}

	log.WithFields(log.Fields{
		"port":  conf.Server.Port,
		"debug": conf.Server.Debug,
	}).Info("Starting RPIS..")
	api.Start()
}
