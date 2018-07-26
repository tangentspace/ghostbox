#!/usr/bin/env bash

# Change to repo root directory
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
cd -P "$( dirname "$SOURCE" )/.."

ssh-keygen -t rsa -N "" -f config/gbadmin.key
