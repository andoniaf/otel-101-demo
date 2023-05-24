require 'sinatra'
require_relative './config'

get '/' do
  'Hello automatic world!'
end

get '/attributes' do
  # get the current auto-instrumented span
  current_span = OpenTelemetry::Trace.current_span
  current_span.add_attributes({
    "app.demo.example" => "auto"
  })

  'Hello manual world! With attributes'
end
