#! /usr/bin/env sh

# args=("$@")
/usr/bin/trufflehog $* --json > trufflehog/trufflehog_output.json
