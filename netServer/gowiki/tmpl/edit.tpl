<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>INDEX</title>
</head>
<h1>Editing {{.Title}}</h1>

<form action="/save/{{.Title}}" method="post">
    <div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
    <div><input type="submit" value="Save"></div>
</form>
</html>