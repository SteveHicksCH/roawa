#!/bin/bash
#
# Start script for roawa in local CHS (within Vagrant) 

BINARY_NAME=${GOPATH}/bin/roawa

source ~/.chs_env/global_env

exec ${BINARY_NAME} 
