{{ template "layout/layout.tpl" . }}

{{ define "contents" }}
  
  {{ template "partial/flash.tpl" .}}

<a href="/admin/content/{{.Name}}/editor" class="button is-primary">Add {{title .Name}}</a>
<br><br>
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

<nav class="pagination is-centered" role="navigation" aria-label="pagination">
    <a class="pagination-previous">Previous</a>
    <a class="pagination-next">Next page</a>
    <ul class="pagination-list">
      <li><a class="pagination-link" aria-label="Goto page 1">1</a></li>
      <li><span class="pagination-ellipsis">&hellip;</span></li>
      <li><a class="pagination-link" aria-label="Goto page 45">45</a></li>
      <li><a class="pagination-link is-current" aria-label="Page 46" aria-current="page">46</a></li>
      <li><a class="pagination-link" aria-label="Goto page 47">47</a></li>
      <li><span class="pagination-ellipsis">&hellip;</span></li>
      <li><a class="pagination-link" aria-label="Goto page 86">86</a></li>
    </ul>
  </nav>

{{ end }}