# GhostBox

[![CircleCI](https://circleci.com/gh/tangentspace/ghostbox.svg?style=svg)](https://circleci.com/gh/tangentspace/ghostbox)

GhostBox is a virtual appliance that allows users to easily deploy and manage
the [Ghost open source publishing platform](https://ghost.org/). It consists of
* Virtual machine images for VirtualBox, KVM, Amazon EC2, and Google Compute Engine.
* (TODO) The `gbutil` CLI utility to deploy and manage GhostBox virtual machines.

## Developing GhostBox

### Prerequisites

Building and testing GhostBox requires installing a few tools on your system.
* [Packer](https://www.packer.io/intro/getting-started/install.html) is required
  to build VM images for different platforms.
* [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
  is required to build VM images.
* [Go](https://golang.org/dl/) is required to build the `gbutil` CLI management
  tool.
* [Terraform](https://www.terraform.io/downloads.html) is required to provision
  and test GhostBox virtual machines on Amazon EC2 or Google Compute Engine
  cloud platforms.
* [Vagrant](https://www.vagrantup.com/downloads.html) is required to provision
  and test GhostBox virtual machines on VirtualBox or Linux KVM hypervisors.
