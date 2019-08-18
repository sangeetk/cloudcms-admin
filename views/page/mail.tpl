{{ template "layout/layout.tpl" . }}

{{ define "contents" }}

  {{ template "partial/flash.tpl" .}}

  {{ $roundcube_host := getenv "ROUNDCUBE_HOST" }}
  {{ $roundcube_user := getenv "ROUNDCUBE_USER" }}
  {{ $roundcube_pass := getenv "ROUNDCUBE_PASS" }}

  {{ if $roundcube_host }}
  	<ul>
  		<li><strong>Webmail:  </strong><a href="{{$roundcube_host}}" target="_blank">{{$roundcube_host}}</a></li>
  		<li><strong>Username: </strong>{{$roundcube_user}}</li>
  		<li><strong>Password: </strong>{{$roundcube_pass}}</li>
  	</ul>
  <iframe src="{{$roundcube_host}}" frameborder="0" marginheight="0" marginwidth="0" width="100%" height="100%"></iframe>
  {{ end }}

{{ end }}