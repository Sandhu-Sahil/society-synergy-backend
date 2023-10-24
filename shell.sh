#!/bin/bash

# check go installation and install if not installed
if ! [ -x "$(command -v go)" ]; then
  echo 'Error: go is not installed.' >&2
  echo 'Installing go...' >&2
  sudo apt-get install golang-go
fi

# install go packages
echo 'Installing go packages...' >&2
go install

# build the project
echo 'Building the project...' >&2
go build -o society-synergy-backend

# run the project
echo 'Running the project...' >&2
./society-synergy-backend

# remove the executable
echo 'Removing the executable...' >&2
rm society-synergy-backend
