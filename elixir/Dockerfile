FROM elixir:1.16-slim
WORKDIR /app
COPY mix.exs ./
RUN mix deps.get
COPY . .
RUN mix compile
CMD ["elixir", "-e", "App.hello(); Process.sleep(:infinity)"]
