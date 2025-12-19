# Go URL Shortener

## Run locally
go run cmd/server/main.go

## Create short URL
POST /api/shorten
{
  "url": "https://example.com"
}

## Redirect
GET /{shortCode}
