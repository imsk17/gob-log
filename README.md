<div align="center">

<h1>gob-log</h1>

A Lightning Fast, Log File Themed Blog with VIM Aesthetics, written in Go.

[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/imsk17/Gbu-Agenda/graphs/commit-activity)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![Last Commit](https://img.shields.io/github/last-commit/imsk17/gob-log)](https://github.com/imsk17/gob-log/commits/main)
[![Stars](https://img.shields.io/github/stars/imsk17/gob-log?style=social)](https://github.com/imsk17/gob-log/stargazers)
</div>

# Features

The blogs are embedded into the binary using Go's `embed` package.

- Uses [Tailwind CSS](https://tailwindcss.com/) for Styling Needs.
- Uses [Fiber](https://github.com/gofiber/fiber) as Web Framework.

## Running Locally

* First, Make sure you have the latest versions of go, and yarn installed.

* Then, Run the command `yarn install` to install the node dependencies on your machine.

* Then, Create the stylesheet by entering `yarn release` into your terminal.

* Modify the blog according to your needs.

* Build the binary by `go build -o blog -ldflags "-s" main.go` .

* Run the binary on your machine. 

### Code style

This project uses [gofmt](https://golang.org/cmd/gofmt/), for Formatting.

## Contributions

If you've found an error in this blog, please file an issue.

Patches are encouraged and may be submitted by forking this project and
submitting a pull request.

## TODO
Some of the future goals in no particular order are - 
- [ ] Responsive Design
- [ ] More SEO Tags
- [ ] Use the http package to reduce dependency on fiber.
- [ ] Get Rid of Tailwind. (!important)