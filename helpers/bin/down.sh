#!/bin/bash

if [ -z $1 ]; then
    echo usage: $0 resourcegroup
    exit 1
fi


az group delete --name "$1" --no-wait --yes

