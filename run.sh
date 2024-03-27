#!/bin/zsh

# ##############################################################################

do_build() {
	go build -o bin/stacker
}

do_cover() {
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
}

do_test() {
	watch --color go test ./tests/...
}

do_vet() {
	watch --color go vet
}

do_fmt() {
	watch gofmt -d ***/*.go
}

usage() {
	echo "Usage statment"

	format="%4s %-4s : %s\n"
	printf "${format}" "Flag" "Arg" "Description"
	printf "${format}" "----" "----" "----"
	printf "${format}" "b" "----" "Build application"
	printf "${format}" "t" "----" "Run tests in a watch"
	printf "${format}" "v" "----" "Run vet in a watch"
	printf "${format}" "x" "msg" "Prints out a message"
}

# ##############################################################################

# Behavior for no parameters
if [ $# -eq 0 ]; then
	usage
	exit 1
fi

while getopts bcftvx: opt ; do
    case "${opt}" in
        b) do_build ;;
        c) do_cover ;;
        t) do_test ;;
        v) do_vet ;;
        f) do_fmt ;;
        x) echo "${OPTARG}" ;;
        *) usage ;;
    esac
done
