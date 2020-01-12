# `point`
## A simple tool for creating HTML presentations

## Installation

`go get github.com/skuzzymiglet/point`

##  Usage

Two blank lines in your `md` file indicate a new slide

`point -i IN [-s STYLE]`

(IN being a markdown file)

Prints a HTML file to stdout by default.

It can then be opened in a browser. Use <kbd>Right</kbd> and <kbd>Left</kbd> to move through the slides

## TODO

- [ ] Integrated mode: embeds CSS, JS and images (as `base64`) in the file to make it truly portable
