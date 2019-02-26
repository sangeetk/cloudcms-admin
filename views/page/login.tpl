<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Admin Login</title>
    <!-- Bulma Version 0.7.2-->
    <link rel="stylesheet" type="text/css" href="/badmin/css/bulma.min.css">
    <style type="text/css">
        html,body {
          font-family: 'Open Sans', serif;
          font-size: 14px;
          font-weight: 300;
        }
        .hero.is-success {
            {{ $bgcolor := getenv "BACKGROUND_COLOR" }}
          background: {{if $bgcolor}}{{$bgcolor}}{{else}}#00D1B2{{end}};
        }
        .hero .nav, .hero.is-success .nav {
          -webkit-box-shadow: none;
          box-shadow: none;
        }
        .box {
          margin-top: 5rem;
        }
        .avatar {
          margin-top: -70px;
          padding-bottom: 20px;
        }
        .avatar img {
          padding: 5px;
          background: #fff;
          border-radius: 50%;
          -webkit-box-shadow: 0 2px 3px rgba(10,10,10,.1), 0 0 0 1px rgba(10,10,10,.1);
          box-shadow: 0 2px 3px rgba(10,10,10,.1), 0 0 0 1px rgba(10,10,10,.1);
        }
        input {
          font-weight: 300;
        }
        p {
          font-weight: 700;
        }
        p.subtitle {
          padding-top: 1rem;
        }
        .formbg {
          background-color: {{getenv "BACKGROUND_COLOR"}}; 
          color: #ffffff;"
        }
    </style>
</head>

<body>
    <section class="hero is-success is-fullheight">
        <div class="hero-body">
            <div class="container has-text-centered">
                <div class="column is-4 is-offset-4">
                    <h3 class="title">{{ .Title }}</h3>
                    {{ with .Error }}
                    <p class="subtitle">{{ . }}</p>
                    {{ end }}
                    <div class="box">
                        <figure class="avatar">
                            {{ $logo := getenv "LOGO"}}
                            <img src="{{if $logo}}{{$logo}}{{else}}/badmin/hunabku.png{{end}}" height="128" width="128">
                        </figure>
                        <form method="POST" action="/admin">
                            <div class="field">
                                <div class="control">
                                    <input class="input is-large" name="username" type="username" placeholder="Username" autofocus="">
                                </div>
                            </div>

                            <div class="field">
                                <div class="control">
                                    <input class="input is-large" name="password" type="password" placeholder="Password" >
                                </div>
                            </div>
                            <div class="field">
                                <label class="checkbox">
                                    <input type="checkbox">
                                      Remember me
                                    </label>
                            </div>
                            <button class="button is-block is-warning is-large is-fullwidth">Login</button>
                        </form>
                    </div>
                    <p class="">
                        Powered by <a href="https://www.urantiatech.com">UrantiaTech</a> 
                    </p>
                </div>
            </div>
        </div>
    </section>
    <script async type="text/javascript" src="/badmin/js/login.js"></script>
</body>

</html>

