package main

import (
	. "github.com/saschagrunert/demo"
	"github.com/andoniaf/otel-101-demo/demo-script/pkg"
	"github.com/urfave/cli/v2"
)

func main() {
	d := New()
	d.Name = "Opentelemetry 101 demo"
	d.HideVersion = true
	d.Usage = "Run Otel 101 demo"
	d.Authors = []*cli.Author{{
		Name: "Andoni Alonso", Email: "andonialonsof@gmail.com",
	}}
	d.Setup(pkg.Setup)
	d.Cleanup(pkg.Cleanup)

	d.Add(pkg.RunDemoSimple(), "demo_simple", "Run simple Otel demo")
	d.Add(pkg.RunOtelDemo(), "demo_otel_demo", "Run Opentelemetry Demo demo")
	d.Run()
}
