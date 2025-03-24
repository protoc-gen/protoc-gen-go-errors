{{ range .Errors }}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Is{{.CamelValue}}(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == {{ .Name }}_{{ .Value }}.String() && e.Code == {{ .HTTPCode }}
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Error{{ .CamelValue }}() *errors.Error {
	 return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), "")
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Error{{ .CamelValue }}WithContext(ctx context.Context, data map[string]any) *errors.Error {
	return errors.NewWithContext({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), ctx, data)
}

{{- end }}
