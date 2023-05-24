package pkg

import (
	. "github.com/saschagrunert/demo"
)

func RunOtelDemo() *Run {
	r := NewRun(
		"opentelemetry-demo demo",
		"(yeah, I know, not the best title...)",
	)

	r.Step(S(
		"Opentelemetry provides a demo environment more",
		"complete than the one I created.",
		"(It wasn't hard...)",
		"",
		"Here they have an architecture overview:",
	), S(
		"docker-compose -f opentelemetry-demo/docker-compose.yml up -d &&",
		"open https://opentelemetry.io/docs/demo/architecture/",
	))

	r.Step(S(
		"Let's check their Ruby service example:",
	), S(
		"open https://github.com/open-telemetry/opentelemetry-demo/blob/main/src/emailservice/email_server.rb",
	))

	r.Step(S(
		"and now the traces:",
	), S(
		"open \"http://$(docker-compose -f opentelemetry-demo/docker-compose.yml port jaeger 16686)\"",
	))

	return r
}
