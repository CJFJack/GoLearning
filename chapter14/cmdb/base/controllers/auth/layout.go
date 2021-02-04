package auth

type LayoutController struct {
	AuthorizationController
}

func (c *LayoutController) Prepare() {
	c.AuthorizationController.Prepare()
	c.Layout = "base/layouts/layout.html"
}
