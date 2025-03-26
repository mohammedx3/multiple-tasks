{{/*
Generate the name of the application based on release name and chart name
*/}}
{{- define "shop-backend.name" -}}
{{- if .Values.nameOverride -}}
{{ .Values.nameOverride }}
{{- else -}}
{{ .Chart.Name }}
{{- end -}}
{{- end }}

{{/*
Generate the full name of the application, including release name
*/}}
{{- define "shop-backend.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{ .Values.fullnameOverride }}
{{- else -}}
{{ .Release.Name }}
{{- end -}}
{{- end }}

{{/*
Common labels for all resources
*/}}
{{- define "shop-backend.labels" -}}
app.kubernetes.io/name: {{ include "shop-backend.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}
