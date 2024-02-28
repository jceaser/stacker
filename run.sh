#!/bin/zsh

# ##############################################################################

do_build() {
	go build -o bin/stacker
}

do_test() {
	go test
}

usage() {
	echo "Usage statment"

	format="%4s %-4s : %s\n"
	printf "${format}" "Flag" "Arg" "Description"
	printf "${format}" "----" "----" "----"
	printf "${format}" "b" "----" "Build application"
	printf "${format}" "2" "----" "Run tests"
	printf "${format}" "x" "msg" "Prints out a message"
}

# ##############################################################################

# Behavior for no parameters
if [ $# -eq 0 ]; then
	usage
	exit 1
fi

while getopts btx: opt ; do
    case "${opt}" in
        b) do_build ;;
        t) do_test ;;
        x) echo "${OPTARG}" ;;
        *) usage ;;
    esac
done
