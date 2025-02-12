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

## Custom Redpanda Connect Build

To optimize the image size, you can create a custom Redpanda Connect image that includes only the essential components for your data pipelines.

If you've developed custom Redpanda Connect components, such as inputs, outputs, or processors, you will need a custom build to integrate them into your pipelines.

### 1. Clone the GitHub Repository

https://github.com/redpanda-data/redpanda-connect-plugin-example

### 2. Update Dependencies

Install [Go](https://go.dev/) and run the following commands:

```bash
go get github.com/redpanda-data/connect/v4@latest
go mod tidy
```

### 3. Edit `main.go`

```go
package main

import (
	"context"

	"github.com/redpanda-data/benthos/v4/public/service"

	_ "github.com/lib/pq"
	_ "github.com/redpanda-data/connect/v4/public/components/io"
	_ "github.com/redpanda-data/connect/v4/public/components/nats"
	_ "github.com/redpanda-data/connect/v4/public/components/prometheus"
	_ "github.com/redpanda-data/connect/v4/public/components/pure"
	_ "github.com/redpanda-data/connect/v4/public/components/sql/base"
)

func main() {
	service.RunCLI(context.Background())
}
```

### 4. Docker Build

```bash
docker build . -t redpanda-connect
```
