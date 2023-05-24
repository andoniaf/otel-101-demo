package pkg

import (
	. "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func Cleanup(ctx *cli.Context) error {
	_ = Ensure(
		"docker-compose -f docker-compose.yml down &&",
		"docker-compose -f opentelemetry-demo/docker-compose.yml down",
	)
	return nil
}
