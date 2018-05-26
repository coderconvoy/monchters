package moncards

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/coderconvoy/lazyf"
	"github.com/coderconvoy/lz2"
	"github.com/coderconvoy/msvg"
	"github.com/coderconvoy/qrunner"
)

func Main(conf *lz2.Config) {
	cardlocs := conf.FlagDef("c", "cards.lz", "location of card list", "cards")
	out := conf.FlagDef("o", "out/cards", "base filename for output", "out")
	lpath := conf.FlagDef("lp", "", "Relative location of links in output", "link-path")
	if conf.Help(false, "pname moncards [-flag value]") {
		return
	}

	fpath := path.Dir(conf.Location)

	items := []lazyf.LZ{}
	for _, v := range strings.Split(cardlocs, ",") {
		v = strings.TrimSpace(v)
		floc := path.Join(fpath, v)
		if len(v) == 0 {
			continue
		}
		if v[0] == '/' {
			floc = v[1:]
		}

		adds, _, err := lazyf.GetConfig(floc)
		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, adds...)
	}

	cards := []msvg.Card{}
	//Count Card Types
	countmap := make(map[string]int)
	for _, v := range items {
		tp := v.PStringD("None", "tp")
		countmap[tp]++
		bc := NewBasic(v, lpath)
		bc2 := bc
		bc2.Bg = "#99ff99"
		cards = append(cards, bc, bc2)

	}

	for k, v := range countmap {
		fmt.Printf("%s:%d\n", k, v)
	}

	cards = msvg.SpreadCards(cards)

	outp := path.Join(fpath, out)
	//make pages
	pgnum := 0

	unitable := []string{}

	writePage := func(fpath string, cc []msvg.Card) {
		fmt.Printf("Making:%s%d\n", fpath, pgnum)
		pg := msvg.PageA4(35, 5, msvg.SpreadCF(cc))
		outf := fmt.Sprintf("%s%d.svg", fpath, pgnum)
		ioutil.WriteFile(outf, pg.Bytes(), 0777)
		outpdf := fmt.Sprintf("%s%d.pdf", fpath, pgnum)
		qrunner.Run("inkscape", outf, "--export-pdf="+outpdf)

		unitable = append(unitable, outpdf)
	}

	for i := 0; i < len(cards); i += 35 {
		writePage(outp, cards[i:])
		pgnum++
	}
	unitable = append(unitable, outp+"-all.pdf")

	qrunner.Run("pdfunite", unitable...)

}
