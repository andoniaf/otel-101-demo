package pkg

import (
	. "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func Setup(ctx *cli.Context) error {
	_ = Ensure(
	)
	return nil
}
