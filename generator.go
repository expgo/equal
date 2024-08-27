package equal

import (
	"embed"
	"github.com/expgo/ag/api"
	"sort"
	"strings"
	"text/template"
)

//go:embed equal.tmpl
var equalTmpl embed.FS

type Generator struct {
	api.BaseGenerator[Equal]
}

func NewGenerator(allEquals []*Equal) *Generator {
	result := &Generator{}

	tmpl := template.New("equal")

	result.Tmpl = template.Must(tmpl.ParseFS(equalTmpl, "*.tmpl"))

	result.DataList = append([]*Equal(nil), allEquals...)

	sort.Slice(result.DataList, func(i, j int) bool {
		x := result.DataList[i]
		y := result.DataList[j]
		return strings.Compare(x.Name, y.Name) < 0
	})

	return result
}
