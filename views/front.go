package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func FrontPage() g.Node {
	return Page(
		"Truffle",
		"/",
		H1(g.Text(`Truffle Wiki`)),
		P(g.Text(`Do you want to learn more about BFT consensus?`)),
	)
}
