# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A minimal Go static file server that serves the current working directory over HTTP. The entire application lives in `main.go` — there are no other packages, dependencies, or tests. Module path: `github.com/shogoki/http-server`.

## Commands

```bash
go build ./...          # build
go vet ./...            # lint
go run . -host 0.0.0.0 -port 3000   # run, serving the current directory (defaults: localhost:8080)
```

Install for end users: `go install github.com/shogoki/http-server`

## Architecture

`main.go` parses `-host` (default `localhost`), `-port` (default 8080), and `-version` flags, then starts `http.ListenAndServe` with `http.FileServer(http.Dir("."))` as the handler. It serves whatever directory it is launched from.

## Releases

Versions are git tags (`vX.Y.Z`) on `main`; the same version string is kept in the `version` constant in `main.go` (printed by `-version`). When cutting a release, bump the constant and tag the merge commit on `main`.
