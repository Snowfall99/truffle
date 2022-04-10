package views

import (
	"path"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"truffle.io/model"
)

func NormalPage(bft model.BFT) g.Node {
	return Page(
		bft.Name,
		path.Join("/", bft.Name),
		H1(g.Text(bft.Name)),
		A(Href(bft.Link), g.Text(`Link`), Br()),
	)
}
