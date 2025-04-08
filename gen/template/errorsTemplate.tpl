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
func Error{{ .CamelValue }}WithMessage(ctx context.Context, format string, params ...any) *errors.Error {
	return errors.NewWithContext(ctx, {{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf(format, params...))
}

func Error{{ .CamelValue }}WithContext(ctx context.Context, params ...any) *errors.Error {
    if len(params) == 0 {
        return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), "")
    }
	return errors.NewWithContext(ctx, {{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), params)
}

{{- end }}
