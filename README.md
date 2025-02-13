# wackylighthouse
Simulate Lighthouse responses for testing

## Usage

Request the page you need.

For example, to test a valid response:
```bash
npx lighthouse https://wackylighthouse.com
```

To test a response with no FCP:
```bash
npx lighthouse https://wackylighthouse.com/no-fcp.html
```

You can even run it through Pagespeed Insights:
```bash
curl https://www.googleapis.com/pagespeedonline/v5/runPagespeed\?url\=https://wackylighthouse.com
```

## Development

### Run the server

Start a local server, because Lighthouse can only run on HTTP/HTTPS.

```bash
go run cmd/main.go
```

### Run Lighthouse

```bash
npx lighthouse http://localhost:8080
```

## Production

Serve the [`./static`](./static) directory.
