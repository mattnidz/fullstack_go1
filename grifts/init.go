package grifts

import (
	"github.com/fullstack_go1/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
