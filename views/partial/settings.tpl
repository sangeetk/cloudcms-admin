{{ template "layout/layout.tpl" . }}

{{ define "contents" }}
	<form method="POST" action="/admin/settings">
		<input id="x" type="hidden" name="body" value="{{with .Page}}{{.Body}}{{end}}">
		<trix-editor class="trix-content" input="x"></trix-editor>
		<br>
		<input type="submit" class="button is-primary" value="Save">
	</form>
{{ end }}