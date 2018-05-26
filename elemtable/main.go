package elemtable

import (
	"fmt"
	"log"

	"github.com/coderconvoy/lz2"
)

func Main(conf *lz2.Config) {
	fname := conf.FlagDef("f", "elems.lz", "Location of Elem Config File", "elem.file")
	lpath := conf.FlagDef("lp", "", "Relative Link Path for html", "elem.link-path")

	if conf.Help(false, "Prints table as html ") {
		return
	}

	elems, err := LoadElems(fname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ElemTable(elems, lpath))

}
