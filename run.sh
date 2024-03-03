#!/bin/sh

npx tailwindcss -i ./static/styles/input.css -o ./static/styles/tailwind.css && \
go run main.go
