package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
)

var r *render.Engine

func init() {
	env := envy.Get("GO_ENV", "development")
	options := render.Options{
		DefaultContentType: "application/json",
	}
	if env == "test" {
		options.Helpers = render.Helpers{
			"myHelper": func() string {
				return "hello"
			},
		}
	}
	r = render.New(options)
}
