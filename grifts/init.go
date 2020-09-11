package grifts

import (
	"github.com/deibyssoca/ds_backend_go/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
