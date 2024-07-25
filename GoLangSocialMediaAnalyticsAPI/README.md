# GoLang Social Media Analytics API

## Overview

A simple API for managing and analyzing social media posts. It allows adding posts, retrieving individual post stats, and aggregating stats across all posts.

## Endpoints

- `POST /add` - Adds a new post with metrics (likes, shares, comments).
- `GET /stats?id=<post_id>` - Retrieves the stats for a specific post.
- `GET /aggregate` - Retrieves aggregate stats for all posts (total likes, shares, and comments).

## Setup

1. Clone the repository.
2. Run `go run main.go` to start the server.
3. Use the `/add` endpoint to add posts.
4. Use the `/stats` endpoint to get stats for a specific post.
5. Use the `/aggregate` endpoint to get aggregate stats.

## Dependencies

- None (uses Go standard library).
