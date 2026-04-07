FROM golang:1.25-bookworm AS builder

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN curl -sLo /usr/local/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 \
    && chmod +x /usr/local/bin/tailwindcss

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN templ generate
RUN tailwindcss -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css
RUN go build -ldflags="-w -s" -o out ./cmd/api

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/out .

EXPOSE 8080
CMD ["./out"]
