#! /usr/bin/env bash

args=("$@")
/usr/bin/trufflehog ${args[@]} --json > trufflehog_output.json
