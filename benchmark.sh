#! /bin/bash

go build
cd examples
echo "pandoc beamer:"
hyperfine --show-output "pandoc -i example.md -o out.pdf"
echo "point:"
hyperfine --show-output "go run ../main.go -i example.md -o out.html"
