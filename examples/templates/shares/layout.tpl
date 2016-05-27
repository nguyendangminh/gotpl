<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta name="keyword" content="Go template helper">
    <meta name="description" content="Go template helper">
    <meta name="author" content="minhnd.com">
    <title>Go template helper</title>
    <link href="http://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.6/css/materialize.min.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  </head>
  <body>
    <div class="container">
        <div class="section">
            <div class="row">
                <div class="col s12 center cyan-text darken-4-text">
                  <h3 class="header">Go template helper</h3>
                </div>
            </div>
            {{$user := index . "user"}}
			{{template "content" $user}}
        </div>
    </div>
    <script type="text/javascript" src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.6/js/materialize.min.js"></script>
  </body>
</html>