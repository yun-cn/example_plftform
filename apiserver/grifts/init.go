package grifts

import (
	"github.com/yanshiyason/noonde_platform/apiserver/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
