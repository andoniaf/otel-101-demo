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
		"(It wasn't hard to be honest...)",
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
		"and now the traces for that service:",
	), S(
		"open \"http://$(docker-compose -f opentelemetry-demo/docker-compose.yml port jaeger 16686)/jaeger/ui/search?limit=20&service=emailservice\"",
	))

	r.Step(S(
		"We could also check all the traces from the 'paymentservice'",
		"using VISA cards:",
	), S(
		"open \"http://$(docker-compose -f opentelemetry-demo/docker-compose.yml port jaeger 16686)/jaeger/ui/search?limit=20&lookback=1h&service=paymentservice&tags=%7B%22app.payment.card_type%22%3A%22visa%22%7D\"",
	))

	r.Step(S(
		"Or the traces from the 'frontend', that were okey (200)",
		"and took more than 350ms:",
	), S(
		"open \"http://$(docker-compose -f opentelemetry-demo/docker-compose.yml port jaeger 16686)/jaeger/ui/search?limit=20&lookback=1h&minDuration=350ms&service=frontend\"",
	))

	r.Step(S(
		"And that's it for the Opentelemetry Demo demo!",
	), nil)

	return r
}
