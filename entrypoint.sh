#! /usr/bin/env sh

# args=("$@")
mkdir "trufflehog"
chmod -R 777 trufflehog
/usr/bin/trufflehog $* > trufflehog/trufflehog_output.json
