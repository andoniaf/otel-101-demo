require 'opentelemetry/sdk'
require 'opentelemetry/exporter/otlp'
require 'opentelemetry/instrumentation/sinatra'

OpenTelemetry::SDK.configure do |c|
  c.service_name = 'manual_demo'
end
