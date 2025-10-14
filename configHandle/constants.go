package confighandle

const Version = "v0.0.1"

const AppName = "var-saver"

// Templates for printing to CLI

const varTemplate = `*Variable*
  Name: {{ .Name }}
  Environment: {{ .Environment }}
`

const metaTemplate = `---------------------------------------
*Metadata*
Application: {{ .AppName }}
Version: {{ .Version }}
Variables Last Updated: {{ .LastUpdated }}
---------------------------------------
`

const projectTemplate = `
**Project: '{{ .Name }}'**
{{ .Description }}
Total Variables: {{ .Variables | len }}
Variables:
{{- range .Variables }}
*****
Variable:
    name: {{ .Name }}
    env: {{ .Environment }}
{{- end }}
=======================================
`
