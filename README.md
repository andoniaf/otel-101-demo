# Opentelemetry 101 Demo
> Demo materials to show Opentelemetry basic concepts.

---
## Usage
```
$ make
Available rules:

build               Build images
build-script        Build go binary script
down                Stop demo
down-otel           Stop otel-demo
up                  Run demo containers
up-otel-demo        Run compose for otel-demo
```

### Demo script
```
$ ./demo-script/otel-101-demo --help
NAME:
   otel-101-demo - Run Opentelemetry 101 demo

USAGE:
   otel-101-demo [global options] command [command options] [arguments...]

AUTHOR:
   Andoni Alonso <andonialonsof@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --all, -l                     run all demos (default: false)
   --auto, -a                    run the demo in automatic mode, where every step gets executed automatically (default: false)
   --auto-timeout auto, -t auto  the timeout to be waited when auto is enabled (default: 3s)
   --continuously, -c            run the demos continuously without any end (default: false)
   --hide-descriptions, -d       hide descriptions between the steps (default: false)
   --immediate, -i               immediately output without the typewriter animation (default: false)
   --skip-steps value, -s value  skip the amount of initial steps within the demo (default: 0)
   -0, --demo_simple             Run simple Otel demo (default: false)
   -1, --demo_otel_demo          Run Opentelemetry Demo demo (default: false)
   --help, -h                    show help (default: false)
```

## Slides
Slides from the ["Hacknight: Opentelemetry 101"](https://www.meetup.com/hacknight-valencia/events/293587428) are available [here](./).