<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Admin Login</title>
    <!-- Bulma Version 0.7.2-->
    <link rel="stylesheet" type="text/css" href="/badmin/css/bulma.min.css">
    <link rel="stylesheet" type="text/css" href="/badmin/css/login.css">
</head>

<body>
    <section class="hero is-success is-fullheight">
        <div class="hero-body">
            <div class="container has-text-centered">
                <div class="column is-4 is-offset-4">
                    <h3 class="title has-text-grey">{{ .Title }}</h3>
                    {{ with .Error }}
                    <p class="subtitle has-text-grey">{{ . }}</p>
                    {{ end }}
                    <div class="box">
                        <figure class="avatar">
                            <img src="/badmin/hunabku.png">
                        </figure>
                        <form method="POST" action="/admin">
                            <div class="field">
                                <div class="control">
                                    <input class="input is-large" name="username" type="username" placeholder="Username" autofocus="">
                                </div>
                            </div>

                            <div class="field">
                                <div class="control">
                                    <input class="input is-large" name="password" type="password" placeholder="Password">
                                </div>
                            </div>
                            <div class="field">
                                <label class="checkbox">
                                    <input type="checkbox">
                                      Remember me
                                    </label>
                            </div>
                            <button class="button is-block is-info is-large is-fullwidth">Login</button>
                        </form>
                    </div>
                    <p class="has-text-grey">
                        Powered by <a href="https://www.urantiatech.com">UrantiaTech</a> 
                    </p>
                </div>
            </div>
        </div>
    </section>
    <script async type="text/javascript" src="/badmin/js/login.js"></script>
</body>

</html>

