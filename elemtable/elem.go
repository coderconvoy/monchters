package elemtable

import (
	"fmt"
	"path"
	"strings"

	"github.com/coderconvoy/lz2"
)

type Elem struct {
	Name    string
	Strong  []string
	Weak    []string
	Slow    []string
	Stop    []string
	Special string
}

func cSlice(lz lz2.LZ, prop string) []string {
	s := lz.PStringD("", prop)
	ss := strings.Split(s, ",")
	res := []string{}
	for _, v := range ss {
		v := strings.TrimSpace(v)
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func NewElem(lz lz2.LZ) Elem {
	return Elem{
		Name:    lz.Name,
		Strong:  cSlice(lz, "strong"),
		Weak:    cSlice(lz, "weak"),
		Slow:    cSlice(lz, "slow"),
		Stop:    cSlice(lz, "stop"),
		Special: lz.PStringD("", "special"),
	}
}

func LoadElems(fname string) ([]Elem, error) {
	res := []Elem{}
	cf, err := lz2.ReadFile(fname, true)
	if err != nil {
		return res, err
	}

	for _, v := range cf.LL {
		res = append(res, NewElem(v))
	}
	return res, nil
}

func htmlImages(ss []string, lpath string) string {
	res := ""
	for _, v := range ss {
		res += fmt.Sprintf(`<img src="%s.svg">`, path.Join(lpath, v))
	}

	return res
}

func ElemTable(elems []Elem, lpath string) string {
	res := "<table><tr>\n"
	res += "<th>Name</th><th>Strong</th><th>Weak</th><th>Slow</th><th>Stop</th><th>Special</th>"
	res += "</tr>\n"
	for _, v := range elems {
		res += v.Html(lpath)
	}
	res += "</table>"
	return res
}

func (e Elem) Html(lpath string) string {
	res := "<tr><th>" + htmlImages([]string{e.Name}, path.Join(lpath, "elems")) + "<br>" + e.Name + "</th>\n"
	res += "<td>" + htmlImages(e.Strong, path.Join(lpath, "elems")) + "</td>\n"
	res += "<td>" + htmlImages(e.Weak, path.Join(lpath, "elems")) + "</td>\n"
	res += "<td>" + htmlImages(e.Slow, path.Join(lpath, "tiles")) + "</td>\n"
	res += "<td>" + htmlImages(e.Stop, path.Join(lpath, "tiles")) + "</td>\n"
	res += "<td>" + e.Special + "</td>\n"

	res += "</tr>"
	return res
}
