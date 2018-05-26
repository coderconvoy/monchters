package moncards

import (
	"fmt"
	"path"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/coderconvoy/lazyf"
)

type basicCard struct {
	Name   string
	Type   string
	Bg     string
	Path   string
	N      int
	Folder string
}

func (bc basicCard) Count() int {
	return bc.N //print 2 of each
}

func NewBasic(lz lazyf.LZ, path string) basicCard {
	return basicCard{
		Name:   lz.Name,
		Path:   path,
		Bg:     "#ffbbbb",
		Type:   strings.ToLower(lz.PStringD("brawn", "type", "Type")),
		N:      lz.PIntD(2, "ex0", "Num"),
		Folder: lz.PStringD("", "Folder", "Fol"),
	}
}

func (bc basicCard) Svg(cw, ch int, g *svg.SVG) {
	ms := cw
	if ms > ch {
		ms = ch
	}
	lcname := strings.ToLower(bc.Name)
	//Background
	g.Rect(0, 0, cw, ch, fmt.Sprintf("stroke:black;fill:%s;stroke-width:%dpx", bc.Bg, ms/50))
	pfolder := "creatures"
	if bc.Folder != "" {
		pfolder = bc.Folder
	}
	g.Image(cw/10, cw/20, cw*8/10, ch*8/10, path.Join(bc.Path, pfolder, lcname+".svg"))
	g.Image(5, 5, cw/5, cw/5, path.Join(bc.Path, "elems", bc.Type+".svg"))
	g.Text(cw/2, ch*9/10, bc.Name, fmt.Sprintf("stroke:none;fill:black;text-anchor:middle;font-family:Arial;font-size:%d", ch/14))

}
