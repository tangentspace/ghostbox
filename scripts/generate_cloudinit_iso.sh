#!/usr/bin/env bash

# Change to repo root directory
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
cd -P "$( dirname "$SOURCE" )/.."

cloud-localds --verbose packer/ubuntu-cloud-init.iso packer/virtualbox-ubuntu-cloud-user-data.txt
