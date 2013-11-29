<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<h2>Cookies</h2>
		<p id="cookie">
			<ol>
			{{range .}}
				<li>{{.}}</li>
			{{end}}
			</ol>
		</p>
		<form action="/">
			<input type="submit" value="Reload"/>
		</form>
	</body>
</html>
