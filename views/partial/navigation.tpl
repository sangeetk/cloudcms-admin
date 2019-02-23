  <nav class="menu">

    <p class="menu-label">
      General
    </p>
    <ul class="menu-list">
      <li><a class="is-active" href="/admin/dashboard"><span class="icon is-small"><i class="fa fa-tachometer"></i></span> Dashboard</a></li>
    </ul>

    <p class="menu-label">
      Config
    </p>
    <ul class="menu-list">
      <li><a href="/admin/about"><span class="icon is-small"><i class="fa fa-list"></i></span> Variables</a></li>
    </ul>

    <p class="menu-label">
      CMS Contents
    </p>
    <ul class="menu-list">
      {{ range $name, $fields := .Schema }}
      <li><a href="/admin/content/{{$name}}"><span class="icon is-small"><i class="fa fa-cogs"></i></span> {{title $name}}</a></li>
      {{ end }}
    </ul>

    <p class="menu-label">
      Global Services
    </p>
    <ul class="menu-list">
      <li><a href="/admin/transcripts"><span class="icon is-small"><i class="fa fa-file"></i></span> Transcripts</a></li>
      <li><a href="/admin/transcripts"><span class="icon is-small"><i class="fa fa-book"></i></span> Urantia Book</a></li>
      <li><a href="/admin/blogs"><span class="icon is-small"><i class="fa fa-book"></i></span> UB Dictionary</a></li>
    </ul>


    <!-- p class="menu-label">
      Users
    </p>
    <ul class="menu-list">
      <li><a class=""><i class="fa fa-users"></i> Users</a></li>
      <li><a class=""><i class="fa fa-user-plus"></i> Add User</a></li>
    </ul -->

    <p class="menu-label">
      Administration
    </p>
    <ul class="menu-list">

      <li>
        <a class=""><i class="fa fa-cog"></i> Settings</a>
      </li>

      <li><a href="/admin/logout"><span class="icon is-small"><i class="fa fa-power-off"></i></span> Logout</a></li>
      <li></li>
      <li></li>

    </ul>


  </nav>
