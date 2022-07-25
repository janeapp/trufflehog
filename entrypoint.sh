#! /usr/bin/env bash

args=("$@")
/usr/bin/trufflehog ${args[@]} > trufflehog_output.json
