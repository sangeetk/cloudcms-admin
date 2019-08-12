{{ template "layout/layout.tpl" . }}

{{ define "contents" }}

  {{ template "partial/flash.tpl" .}}

  <iframe src="https://matomo.urantiatech.com/index.php?module=Widgetize&action=iframe&moduleToWidgetize=Dashboard&actionToWidgetize=index&idSite=2&period=week&date=yesterday&token_auth={{ getenv "MATOMO_TOKEN" }}" frameborder="0" marginheight="0" marginwidth="0" width="100%" height="100%"></iframe>

{{ end }}