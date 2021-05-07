#!/bin/bash

# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin"

# version of air
#air -v

# initialize the .air.toml configuration file to the current directory
air init

# Run app in Hot-reload mode
# air
