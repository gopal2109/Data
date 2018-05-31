Setup and Install
-----------------

* install go 1.6 from golang.org

*setting up work space*

```
$ mkdir rpis
$ cd rpis
$ mkdir src bin pkg
```

Export $GOPATH & $GOROOT, clone repository inside `src` directory. 

```
$ go get github.com/BurntSushi/toml
$ go get go get gopkg.in/mgo.v2
```

Building the source 
-------------------

If you are inside src directory just type in `go build`,
if you are in directory where `src` `pkg` `bin` are avaliable then `go install rpis` will build and move the binary into `bin`
(Clone the repo inside `src` directory)

Installing dependencies (External)
----------------------------------
```
go get (
	github.com/julienschmidt/httprouter
	labix.org/v2/mgo
	github.com/BurntSushi/toml
	github.com/Sirupsen/logrus
)
```

Run Run Run
-----------

```
$ ./rpis --help
$ ./rpis --version
$ ./rpis
```
