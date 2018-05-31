package backend

import (
	"os"
	"rpis/config"
	"github.com/Sirupsen/logrus"
)

var (
	level string
	Log logrus.Logger
)

func init() {
	var fileout *os.File
	Log = *(logrus.New())
	
	if config.Conf.Logging.File != "" {
		fileout, err := os.OpenFile(config.Conf.Logging.File,
			os.O_RDWR | os.O_APPEND,
			0666)
		if err != nil {
			panic(err)
		}
		defer fileout.Close()
	} else {
		fileout = os.Stderr
	}
	
	Log.Out = fileout
}	

	
