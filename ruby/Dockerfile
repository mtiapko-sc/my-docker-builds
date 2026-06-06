FROM ruby:3.3-slim
WORKDIR /app
COPY Gemfile ./
RUN apt-get update && apt-get install -y build-essential && bundle install
COPY . .
EXPOSE 8080
CMD ["ruby", "app.rb"]
