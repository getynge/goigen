package generator

const source = `//generated by goigen, do not edit manually
//go:generate mockgen -source={{.FileName}}.go -package={{.MockPackage}} -destination={{.MockDirectory}}/{{.FileName}}_mock.go
package {{.Package}}

type {{.TargetInterface}} interface {
{{range .Methods}}
	{{.Name}}({{range .Arguments}}{{.Name}} {{.Type}}{{if .HasNext}},{{end}}{{end}}) {{if .Returns}}({{range .Returns}}{{.Name}} {{.Type}}{{if .HasNext}},{{end}}{{end}}){{end}}
{{end}}
}
`
