# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A minimal Go static file server that serves the current working directory over HTTP. The entire application lives in `main.go` — there are no other packages, dependencies, or tests. Module path: `github.com/shogoki/http-server`.

## Commands

```bash
go build ./...          # build
go vet ./...            # lint
go run . -port 3000     # run, serving the current directory (default port: 8080)
```

Install for end users: `go install github.com/shogoki/http-server`

## Architecture

`main.go` parses a `-port` flag (default 8080), then starts `http.ListenAndServe` bound to `localhost` with `http.FileServer(http.Dir("."))` as the handler. The server intentionally binds to localhost only and serves whatever directory it is launched from.
