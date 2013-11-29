<!DOCTYPE html>
<html>
	<head></head>
	<body>
{{if .Username}}	
	<h2>Hello, {{.Username}}</h2>
	<p>
		You visited this page <b>{{.AmountVisited}}</b> times.
	</p>
	<form action="/">
		<input type="submit" value="Reload"/>
	</form>
	<form action="/logout">
		<input type="submit" value="Logout"/>
	</form>
{{else}}
	<h2>Account Login</h2>
	<form action="/login" method="post">
		<label>Username: 
			<input type="text" name="username"/>
		</label>
		<input type="submit" value="Login"/>
	</form>
{{end}}
	</body>
</html>
