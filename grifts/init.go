package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/wbean1/apiwars/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
