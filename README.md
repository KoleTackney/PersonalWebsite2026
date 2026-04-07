# koletackney.dev

Personal portfolio and resume website, rebuilt with Go + Templ + HTMX + Tailwind CSS.

## Tech Stack

- **Go 1.25** — server with stdlib `net/http`
- **Templ** — type-safe HTML templating compiled to Go
- **HTMX** — dynamic interactions without a JS framework
- **Tailwind CSS v4** — utility-first CSS via standalone CLI

## Architecture

Single Go binary, server-side rendered. Templ components produce HTML, HTMX handles interactivity, Tailwind handles styling. Static assets (CSS, JS) are embedded into the binary via `go:embed`. No SPA, no JS bundler.

```
cmd/api/          — Entrypoint
cmd/web/          — Templates, handlers, embedded assets
internal/server/  — Server config, routing, middleware
referenceMaterial/ — Old site HTML, resume content
```

## Prerequisites

- Go 1.25+
- [templ CLI](https://templ.guide/) — installed automatically by `make build` if missing

## Development

```bash
make build    # templ generate → tailwind compile → go build
make run      # go run cmd/api/main.go
make watch    # live reload via Air
make test     # go test ./... -v
```

## Environment

Create a `.env` file:

```
PORT=8080
APP_ENV=local
```

## Deployment

Targeting [Railway](https://railway.app). Configuration TBD.

CI/CD via GitHub Actions — tests on push/PR, releases via GoReleaser on version tags.

## Content

Website content is derived from an external master profile document. See `referenceMaterial/` for additional reference files including the old site HTML and current resume.
