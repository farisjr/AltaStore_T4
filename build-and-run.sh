#!/bin/bash


go build
sudo setcap 'cap_net_bind_service=+ep' "$(realpath project)"
source .env
./project

