{{ template "layout/layout.tpl" . }}

{{ define "contents" }}
  
  {{ template "partial/flash.tpl" .}}

<a href="/admin/content/{{.Name}}/editor" class="button is-primary">Add {{title .Name}}</a>
<br><br>
<h3>Displaying {{.First}} - {{.Last}} of total {{.Total}}</h3>
<br>

<table class="table is-bordered is-hoverable is-fullwidth">
  <thead>
    <tr>
      <th>ID</th>
      <th>Lang</th>
      <th>Title</th>
      <th>Date</th>
      <th>Status</th>
    </tr>
  </thead>
  <tbody>

  {{ $name := .Name }}
  {{ range $item := .List }}
    <tr>
      <th>{{ $item.id }}</th>
      <td>{{ $item.language }}</td>
      <td><a href="/admin/content/{{$name}}/editor?slug={{$item.slug}}">{{ $item.title }}</a></td>
      <td>{{ unixTimeToDateString $item.created_at}}</td>
      <td>{{ status $item.status }}</td>
    </tr>
  {{ end }}
  </tbody>
</table>

<br>

{{ template "partial/pagination.tpl" . }}

<br>
<br>
{{ end }}