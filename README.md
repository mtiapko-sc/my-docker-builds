# my-docker-builds

Test repo for the Docker Build Agent — 10 language directories, each with a deliberately broken Dockerfile. The agent watches for failing builds and opens fix PRs automatically.

## Structure

Each directory contains a minimal app and a broken Dockerfile:

| Directory | Language  | Broken because |
|-----------|-----------|----------------|
| `python/` | Python    | `pip install` runs before `COPY requirements.txt` |
| `nodejs/` | Node.js   | `npm ci` requires `package-lock.json` which doesn't exist |
| `php/`    | PHP       | `apt-get install` without `apt-get update` |
| `rust/`   | Rust      | `cargo build --locked` without `Cargo.lock` |
| `go/`     | Go        | `COPY go.sum` references a file that doesn't exist |
| `java/`   | Java      | `apt-get install` without `apt-get update` |
| `ruby/`   | Ruby      | `nokogiri` gem requires `build-essential` (not installed) |
| `dotnet/` | .NET      | `apt-get install libgdiplus` without `apt-get update` |
| `elixir/` | Elixir    | `mix deps.get` runs without installing Hex first |
| `gradle/` | Gradle    | `./gradlew` runs without `chmod +x` |

## Triggering a build

Go to **Actions → Docker Build → Run workflow**, pick a language from the dropdown. The build will fail with a realistic error.

By default (on push to `main`), the workflow builds `rust`.

## Running the agent

The agent lives in [ai-agent-capstone-proj](../ai-agent-capstone-proj). Start it:

```bash
cd ai-agent-capstone-proj
docker compose up -d   # start Qdrant
uv run python src/agent.py
```

The agent polls every 15 minutes. When it finds a failing build on `main`, it:
1. Fetches the logs and broken Dockerfile
2. Diagnoses the root cause using NVIDIA NIM
3. Generates a minimal fix via RAG + LLM
4. Opens a PR with the fix

## Running the eval suite

Before running the agent against this repo, verify it handles all 10 languages:

```bash
cd ai-agent-capstone-proj
uv run pytest tests/ -v --timeout=600
```
