package main

import (
	"os"

	"github.com/coderconvoy/cardmakers/monchters/elemtable"
	"github.com/coderconvoy/cardmakers/monchters/moncards"
	"github.com/coderconvoy/lz2"
)

func main() {
	//Help
	if lz2.Help(true,
		"First arg must be jobname, after that flags with '-' \nJobs :",
		"mon: Build the list of creature cards",
		"elem : Output The type Array",
		"",
		"-conf : config location",
	) {
		return
	}

	//begin

	conf, _ := lz2.LoadConfigArgs("conf", true, os.Args[2:], "cr-conf.lz")
	switch os.Args[1] {
	case "mon":
		moncards.Main(conf)
	case "elem":
		elemtable.Main(conf)
	}

}
