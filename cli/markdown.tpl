{{ if getLayout -}}
---
title: {{ .CommandPath | title }}
layout: {{ getLayout }}
---
{{ end -}}
## {{.CommandPath}}

{{.Short}}

{{ if .Long -}}
### Synopsis

{{.Long}}

{{ end -}}
{{- if .Runnable -}}

```
{{.UseLine}}
```

{{ end -}}

{{ if .Example -}}
### Examples

```
{{.Example}}
```

{{ end -}}
{{ if .HasAvailableLocalFlags -}}
### Options

```
{{ .NonInheritedFlags.FlagUsages }}
```

{{ end -}}
{{ if .HasAvailableInheritedFlags -}}
### Options inherited from parent commands

```
{{ .InheritedFlags.FlagUsages }}
```

{{ end -}}
{{ $seeAlso := . | getSeeAlso -}}
{{ if $seeAlso -}}
### SEE ALSO

{{ range $index, $cmd := $seeAlso -}}
{{ $link := $cmd.CommandPath | replace " " "-" | printf "%s.md" -}}
* [{{$cmd.CommandPath}}]({{ $link }}) - {{ $cmd.Short }}
{{ end -}}
{{ end -}}
{{ if not (. | isAutoGenDisabled) }}
###### Auto generated by spf13/cobra on {{ "2-Jan-2006" | formatGeneratedAt }}
{{ end -}}
