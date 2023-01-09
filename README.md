# zerolog-loki-grafana

Testing Zerolog (logging library in Go) with Promtail, Loki and Grafana

This was shamelessly referenced from https://github.com/ruanbekker/docker-promtail-loki.

## Usage

```bash
docker-compose up -d
```

Then go to http://localhost:3000, click in Explore on the left sidebar and then type any of following queries:

```text
{container="go-zerolog"}|json|level="debug"|message="hello world"
{container="go-zerolog"}|json|level="warn"|message="hello world"
{container="go-zerolog"}|json|level="error"|message="hello world"
```

