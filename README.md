# `point`
## A simple tool for creating HTML presentations

[![Maintainability](https://api.codeclimate.com/v1/badges/1cede0dc4f659ebb2c3c/maintainability)](https://codeclimate.com/github/skuzzymiglet/point/maintainability)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b238e72b91fb42c2986eee097bc06947)](https://www.codacy.com/manual/skuzzymiglet/point?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=skuzzymiglet/point&amp;utm_campaign=Badge_Grade)
![Go](https://github.com/skuzzymiglet/point/workflows/Go/badge.svg)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fskuzzymiglet%2Fpoint.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fskuzzymiglet%2Fpoint?ref=badge_shield)
## Installation

`go get github.com/skuzzymiglet/point`

## Why?

+ Lightweight and fast
+ No special software needed to display, just a browser
+ Easy CSS styling

## Usage

Every title (`#`/`<h1>`) starts a new slide

`point -i IN [-s STYLE] [-o OUT]`

(IN being a markdown file)

Prints a HTML file to stdout by default.

It can then be opened in a browser, anywhere - there are no file dependencies since everything is embedded. Use <kbd>Right</kbd> and <kbd>Left</kbd> to move through the slides

## TODO

  - [X] Embed CSS
  - [X] Embed JS
  - [X] Embed images as `base64`
  - [ ] Center slide
  - [ ] Show slide number
  - [ ] More styles
  - [ ] Embed web images
  - [ ] Automatic moving through slides
  - [ ] Add tests
  - [ ] Benchmark against `pandoc`-generated beamer
  - [ ] Hosted webapp

## Issues

+ ~~Slides are parsed, changed and rendered individually - each one has a full HTML tree~~ **Fixed in [2b3f10831b](https://github.com/skuzzymiglet/point/commit/2b3f10831bbe38ea49f61e3daed4286bed71d191)**


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fskuzzymiglet%2Fpoint.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fskuzzymiglet%2Fpoint?ref=badge_large)