#!/bin/bash

# Enshure the execution directory is the same as this script
cd "$(dirname "$0")"

# Run all go files 
ls *.go | xargs -L1 -I{} sh -c 'printf "Runing \e[1;93m{}\e[0m: "; go run {}'
