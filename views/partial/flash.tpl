  {{ with .Flash }}
    {{ if .error }}
    <div class="notification is-danger">
    <button class="delete"></button> {{ .error }}
    {{ end }}
    {{ if .warning }}
    <div class="notification is-warning">
    <button class="delete"></button> {{ .warning }}
    {{ end }}
  </div>
  {{ end }}
