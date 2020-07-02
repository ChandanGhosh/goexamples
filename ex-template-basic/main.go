package main

import "fmt"

func main() {
	name := "Chandan Ghosh"
	tmpl := `
<!DOCTYPE html>
<html>
<head>
<title>Sample basic template</title>
</head>
<body>
<h1>` + name + `</h1>
</body>
</html>
`
	fmt.Print(tmpl)
}
