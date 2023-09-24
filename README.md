# interactio-go-kata
REST API endpoint that allows users to create events.

## Code Structure
```
cmd             - CLI logic
docs            - generated Swagger documentation files
internal
  |_ api        - HTTP request and response handling
  |_ repository - entity storage management logic
  |_ service    - domain logic
```

## Usage
```
REST API endpoint that allows users to create events.

Usage:
  server [flags]

Flags:
      --address string       address to launch the HTTP server on (default "localhost:8080")
      --docs                 include documentation routes (default true)
  -h, --help                 help for server
      --sqlite-path string   path to SQLite DB (default "app.sqlite")
```

Build the binary from source and launch the server:
```
make build && make run
```

## Documentation
Swagger Web UI documentation is available on route: http://localhost:8080/docs/
