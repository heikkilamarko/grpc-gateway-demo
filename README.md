# gRPC Gateway Demo

## Start

```bash
docker compose build
```

```bash
docker compose up -d
```

## Stop

```bash
docker compose down -v
```

## Example Requests

| Demo       | Protocol | Example Requests     |
| ---------- | -------- | -------------------- |
| calculator | HTTP     | calculator_http.http |
| calculator | NATS     | calculator_nats.sh   |

## Grafana

http://localhost:8000/grafana

| Username | Password |
| -------- | -------- |
| admin    | S3c_r3t! |
