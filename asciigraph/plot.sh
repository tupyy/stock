#!/bin/bash

while true; do
    sleep 2
    curl -s localhost:18080/stock?label=$1 | jq '.values[0].value'
done | asciigraph -r -h 10 -w 100 -c "$2"
