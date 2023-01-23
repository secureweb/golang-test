package structure

import (
	"embed"
	"html/template"
)

var (
	//go:embed resources
	res   embed.FS
	pages = map[string]string{
		"/": "resources/index.gohtml",
	}
)

func Test() (*template.Template, error) {
	page := pages["/"]
	tpl, err := template.ParseFS(res, page)
	if err != nil {
		return nil, err
	}
	return tpl, err
}
