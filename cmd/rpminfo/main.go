// 指定した パッケージ名 または ファイル名 を rpm パッケージとみなして解析し、結果を出力します.
//
// usage:
//   $ rpminfo openssl-1.1.1c-2.el8.x86_64
//   or
//   $ rpminfo openssl-1.1.1c-2.el8.x86_64.rpm
//
// どちらも結果は以下です.
//   [name] openssl  [version] 1.1.1c        [rel] 2.el8     [arch] x86_64
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/devlights/rpminfo/pkg/rpm"
)

var (
	mainLog = log.New(os.Stdout, "", 0)
	dbgLog  = log.New(ioutil.Discard, "[debug] ", 0)
)

var (
	debug    bool
	newline  bool
	field    string
	filename string
)

func main() {
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.BoolVar(&newline, "newline", false, "output with new line")
	flag.StringVar(&field, "field", "", "specify field value (name/version/rel/arch/epoch)")
	flag.Parse()

	if flag.NArg() == 0 {
		mainLog.Println("usage: rpminfo [option] rpm-pkg-name or filename")
		return
	}

	filename = flag.Arg(0)

	if debug {
		dbgLog.SetOutput(os.Stdout)
	}

	dbgLog.Println(filename)
	os.Exit(run())
}

func run() int {
	r := rpm.Parse(filename)

	if newline {
		r.SetOutputPattern(rpm.RpmOutputNewLine)
	}

	switch field {
	case "":
		mainLog.Println(r)
	default:
		mainLog.Println(r.Get(field))
	}

	return 0
}
