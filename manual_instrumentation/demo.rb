require 'sinatra'
require_relative './config'

tracer = OpenTelemetry.tracer_provider.tracer('manual_demo')

get '/' do
  tracer.in_span("say_hello") do |span|
    'Hello manual world!'
  end
end

get '/attributes' do
  tracer.in_span("with_attributes") do |span|
    span.set_attribute("app.demo.example", "manual")

    'Hello manual world! With attributes'
  end
end
