#!/bin/sh

if ! [ -f "assets/css/tailwind.css" ]; then
    npx tailwindcss -i "assets/css/input.css" -o "assets/css/tailwind.css"
fi

(cd src && go run main.go)
