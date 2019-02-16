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

  <link rel="stylesheet" href="/badmin/css/bulma.min.css">
  <link rel="stylesheet" href="/badmin/fontawesome/css/font-awesome.min.css">

  <link rel="stylesheet" href="/badmin/dist/main.css">
  <link rel="stylesheet" type="text/css" href="/badmin/trix/trix.css">
  <script type="text/javascript" src="/badmin/trix/trix.js"></script>
</head>

<body>

  {{ template "partial/header.tpl" }}
  
  <div class="wrapper">
    <div class="columns">
      <aside class="column is-2 aside">
        {{ template "partial/navigation.tpl" }}
      </aside>

      <main class="column main">
        <nav class="breadcrumb is-small" aria-label="breadcrumbs">
          <ul>
            <li><a href="#">Home</a></li>
            <li class="is-active"><a href="#" aria-current="page">{{ .Title }}</a></li>
          </ul>
        </nav>

        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <div class="title has-text-primary"> {{ .Title }}</div>
            </div>
          </div>
          <div class="level-right">
            <div class="level-item">
              <a type="button" class="button is-small" href="/">
                Visit Site
              </a>
            </div>
          </div>
        </div>

        {{ block "contents" . }} {{ end }}
        
      </main>

    </div>
  </div>

  <script src="/badmin/dist/build.js"></script>

</body>

</html>
