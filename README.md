# go-math-quiz
Math quiz CLI. You give it a list of questions/answers in a CSV, and see how well you do

# Requirements
* [Golang](https://golang.org/) v1.10
* [Glide](https://glide.sh/)

# Getting Started
After you have Golang and Glide installed, run the CLI command to install the dependencies:

`glide install`

# Development
## Testing
* Test Coverage: `go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html && open cover.html`
* Test Suite: `go test`

# Usage
* Run `go run quiz.go`
* If the project is built: `./go-math-quiz`

# Building
* Run `go build`
