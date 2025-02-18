package gen

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed template/errorsTemplate.tpl
var errorsTemplate string

type errorInfo struct {
	Name       string
	Value      string
	HTTPCode   int
	CamelValue string
	Comment    string
	HasComment bool
}

type errorWrapper struct {
	Errors []*errorInfo
}

func (e *errorWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(errorsTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}

//go:embed template/commonTemplate.tpl
var commonTemplate string

func executeCommon() string {
	buf := new(bytes.Buffer)
	commonTmpl, err := template.New("common").Parse(commonTemplate)
	if err != nil {
		panic(err)
	}
	if err := commonTmpl.Execute(buf, nil); err != nil {
		panic(err)
	}
	return buf.String()
}
