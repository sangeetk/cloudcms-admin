	<nav class="pagination is-centered" role="navigation" aria-label="pagination">
		{{ $pagers := pager 10 .CurrentPage .Total .Size .Query }}

		{{ range $p := $pagers }}
			{{ if eq $p.Title "Prev"}}
		    <a class="pagination-previous" href="{{$p.URI}}">Previous</a>
		    {{ else if eq $p.Title "Next"}}
		    <a class="pagination-next" href="{{$p.URI}}">Next page</a>
		    {{ end }}
	    {{ end }}

	    <ul class="pagination-list">
		{{ range $p := $pagers }}
			{{ if eq $p.Title "..."}}
			<li><span class="pagination-ellipsis">&hellip;</span></li>
			{{ else if eq $p.Title "Prev" }}
			{{ else if eq $p.Title "Next" }}
			{{ else if $p.Active }}
			<li><a class="pagination-link is-current" aria-label="{{$p.Title}}">{{$p.Title}}</a></li>
			{{ else }}
			<li><a class="pagination-link" aria-label="{{$p.Title}}" href="{{$p.URI}}">{{$p.Title}}</a></li>
			{{ end }}
		{{ end }}
	    </ul>
	</nav>