  <nav class="menu">

      <p class="menu-label">
        <img class="navbar-item navbar-brand-logo" src="{{with getenv "NAVBAR_IMAGE"}}{{.}}{{else}}/admin/theme/navbar.png{{end}}" alt="CloudCMS Admin">
      </p>

    <p class="menu-label">
      General
    </p>
    <ul class="menu-list">
      {{ if eq .Name "dashboard" }}
      <li><a class="is-active" href="/admin/dashboard"><span class="icon is-small"><i class="fa fa-home"></i></span> Dashboard</a></li>
      {{ else }}
      <li><a href="/admin/dashboard"><span class="icon is-small"><i class="fa fa-home"></i></span> Dashboard</a></li>
      {{ end }}

      {{ $roundcube := getenv "ROUNDCUBE_HOST" }}
      {{ if $roundcube }}
      {{ if eq .Name "mail" }}
      <li><a class="is-active" href="/admin/mail"><span class="icon is-small"><i class="fa fa-envelope"></i></span> Mail</a></li>
      {{ else }}
      <li><a href="/admin/mail"><span class="icon is-small"><i class="fa fa-envelope"></i></span> Mail</a></li>
      {{ end }}
      {{ end }}
    </ul>


    <p class="menu-label">
      Contents
    </p>
    <ul class="menu-list">
      {{ $contentType := .Name }}
      
      {{ range $name, $fields := .Schema }}
      {{ if eq $contentType $name }}
      <li><a class="is-active" href="/admin/content/{{$name}}"><span class="icon is-small"><i class="fa fa-cogs"></i></span> {{title $name}}</a></li>
      {{ else }}
      <li><a href="/admin/content/{{$name}}"><span class="icon is-small"><i class="fa fa-cogs"></i></span> {{title $name}}</a></li>
      {{ end }}
      {{ end }}
    </ul>


    <p class="menu-label">
      Administration
    </p>
    <ul class="menu-list">

      <li><a href="/admin/logout"><span class="icon is-small"><i class="fa fa-power-off"></i></span> Logout</a></li>
      <li></li>
      <li></li>

    </ul>


  </nav>
