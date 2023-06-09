package pkg

import (
	. "github.com/saschagrunert/demo"
)

func RunDemoSimple() *Run {
	r := NewRun(
		"DEMO TIME!",
		"Let's check the code...",
	)

	r.Step(S(
		"I've prepared two simple containers:",
		"",
		"(this is cool right, it feels like Matrix...)",
		"(Well, let's focus)",
		"",
		"",
		"- One with auto instrumentation",
		"- One with manual instrumentation",
	), S(
		"docker-compose up -d &&",
		"open https://github.com/andoniaf/otel-101-demo",
	))

	r.Step(S(
		"We can see in the logs which instrumentation we have enabled,",
		"in the automatic instrumentation example service:",
	), S(
		"docker-compose logs automatic | grep Instrumentation",
	))

	r.Step(S(
		"and in the manual instrumentation service:",
	), S(
		"docker-compose logs manual | grep Instrumentation ",
	))

	r.Step(S(
		"Let's do some requests to both containers",
		"and check how we see those traces in Jaeger:",
	), S(
		"curl localhost:4567; echo &&",
		"curl localhost:4568; echo &&",
		"open http://localhost:16686",
	))

	r.Step(S(
		"Okey, let's do now some requests to the",
		"/attributes endpoint:",
	), S(
		"curl localhost:4567/attributes; echo &&",
		"curl localhost:4568/attributes; echo &&",
		"open http://localhost:16686",
	))

	r.Step(S(
		"And that's it for the simple demo!",
	), nil)

	return r
}
