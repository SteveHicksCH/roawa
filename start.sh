#!/bin/bash
#
# Start script for roawa

# Start script for local CHS (within Vagrant) 

BINARY_NAME=${GOPATH}/bin/roawa

source ~/.chs_env/global_env

exec ${BINARY_NAME} 
