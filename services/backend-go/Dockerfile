FROM cgr.dev/chainguard/go:latest-dev AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN go build -o /out/backend-go ./

FROM cgr.dev/chainguard/static:latest
LABEL application_name="backend-go" \
      description="Pantalasa Cronos monorepo backend-go service" \
      owner="brandon@pantalasa.org" \
      source_uri="https://github.com/pantalasa-cronos/monorepo" \
      git_sha="unknown"
COPY --from=build /out/backend-go /backend-go
EXPOSE 8080
ENTRYPOINT ["/backend-go"]
