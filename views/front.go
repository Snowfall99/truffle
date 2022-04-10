package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"truffle.io/model"
)

func FrontPage(bfts []model.BFT) g.Node {
	body := []g.Node{}
	body = append(body, H1(g.Text(`Truffle Wiki`)))
	body = append(body, P(g.Text(`Do you want to learn more about BFT consensus?`)))
	for _, bft := range bfts {
		body = append(body, A(Href(bft.Name), g.Text(bft.Name), Br()))
	}
	return Page(
		"Truffle",
		"/",
		body...,
	)
}
