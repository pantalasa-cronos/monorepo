FROM cgr.dev/chainguard/python:latest-dev AS build
WORKDIR /app
COPY requirements.txt ./
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt
COPY main.py ./

FROM cgr.dev/chainguard/python:latest
LABEL application_name="api-python" \
      description="Pantalasa Cronos monorepo api-python service" \
      owner="brandon@pantalasa.org" \
      source_uri="https://github.com/pantalasa-cronos/monorepo" \
      git_sha="unknown"
WORKDIR /app
COPY --from=build /install /usr/local
COPY --from=build /app/main.py /app/main.py
EXPOSE 8081
ENTRYPOINT ["python", "/app/main.py"]
