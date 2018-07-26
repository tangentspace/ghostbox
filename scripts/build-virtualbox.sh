#!/usr/bin/env bash

# Change to repo root directory
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
cd -P "$( dirname "$SOURCE" )/.."

packer build --force packer/virtualbox.json && vagrant box remove ghostbox && vagrant box add --name ghostbox output/virtualbox-vagrant/ghostbox.box
