package generator

const template = `//go:generate mockgen -source={{.FileName}}.go -package=mock_{{.FileName}} -destination=mock_{{.FileName}}/mock.go
package {{.Package}}

type {{.TargetInterface}} interface {
{{range .Methods}}
	{{.Name}}({{range .Arguments}}{{.Name}} {{.Type}}{{if .HasNext}},{{end}}{{end}}) {{if .Returns}}({{range .Returns}}{{.Name}} {{.Type}}{{if .HasNext}},{{end}}{{end}}){{end}}
{{end}}
}
`
