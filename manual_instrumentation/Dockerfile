FROM ruby:3.2-slim-bullseye

WORKDIR /app

RUN apt-get update && apt-get install build-essential g++ -y

COPY Gemfile /app

RUN bundle install

COPY demo.rb config.rb /app/

CMD [ "ruby","demo.rb","-o","0.0.0.0"]