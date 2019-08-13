{{ template "layout/layout.tpl" . }}

{{ define "contents" }}

  {{ template "partial/flash.tpl" .}}

  {{ $siteid := get "MATOMO_SITE_ID" }}
  {{ $token := getenv "MATOMO_AUTH_TOKEN" }}
  {{ if $token }}
  <iframe src="https://matomo.urantiatech.com/index.php?module=Widgetize&action=iframe&moduleToWidgetize=Dashboard&actionToWidgetize=index&idSite={{ $siteid }}&period=week&date=yesterday&token_auth={{ $token }}" frameborder="0" marginheight="0" marginwidth="0" width="100%" height="100%"></iframe>
  {{ end }}

{{ end }}