<!DOCTYPE html>
<html lang="en">

<head>
  <meta name="robots" content="index, follow" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=7">

  <meta name="language" content="en-EN" />
  <meta name="author" content="Urantia Tech" />
  <title>CloudCMS Admin - {{ .Title }}</title>

  <link rel="stylesheet" href="/admin/theme/css/bulma.min.css">
  <link rel="stylesheet" href="/admin/theme/fontawesome/css/font-awesome.min.css">

  <link rel="stylesheet" href="/admin/theme/dist/main.css">
  <link rel="stylesheet" type="text/css" href="/admin/theme/trix/trix.css">
  <script type="text/javascript" src="/admin/theme/trix/trix.js"></script>
</head>

<body>

  <div class="wrapper">
    <div class="columns">
      <aside class="column is-2 aside">
        {{ template "partial/navigation.tpl" . }}
      </aside>

      <main class="column main">
        <nav class="breadcrumb is-small" aria-label="breadcrumbs">
          <ul>
            <li><a href="#">Home</a></li>
            <li class="is-active"><a href="#" aria-current="page">{{ .Title }}</a></li>
          </ul>
        </nav>

        <form method="GET" action="/admin/content/{{.Name}}">
        <div class="level">
          
          <div class="level-left">
            <div class="level-item">
              <div class="title has-text-primary"> 
                {{.Title}}
                &nbsp;&nbsp;&nbsp;
              </div>
            </div>
          </div>

        {{ $contentType := .Name }}
        {{ $query := .Query }}
        <div class="field-body">
          <div class="field is-expanded has-addons">
            <p class="control is-expanded">
              <input class="input" type="text" name="q" value="{{with $query}}{{.}}{{end}}" placeholder="Search {{$contentType}} ...">
            </p>
            <p class="control">
              <input class="button is-primary" type="submit" value="Search">&nbsp;&nbsp;&nbsp;
            </p>
          </div>
        </div>

          <div class="level-right">
            <div class="level-item">
              <div class="is-centered">
                <a class="button is-white">Language</a>
                {{ $currentLang := .LanguageCode }}
                {{ $uri := .URI }}
                {{ range $lang := .Languages }}
                  {{ if eq $lang $currentLang }}
                    <a class="button is-primary is-label" title="{{langCodeToName $lang}}">{{uppercase $lang}}</a>
                  {{ else }}
                    <a class="button is-primary is-outlined" href="/admin?lang={{$lang}}&dst={{$uri}}" title="{{langCodeToName $lang}}">{{uppercase $lang}}</a>
                  {{ end }}
                {{ end }}
                &nbsp;
              </div>
            </div>
          </div>
        </div>
        </form>

        <br>
        
        {{ block "contents" . }} {{ end }}
        
      </main>

    </div>
  </div>

  <script src="/admin/theme/dist/build.js"></script>
  <script defer src="/admin/theme/fontawesome/js/all.js"></script>

</body>

</html>
