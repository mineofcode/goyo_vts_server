#!/bin/bash

if docker stop gpst; then docker rm gpst; fi

if docker rmi $(docker images --filter "dangling=true" -q --no-trunc); then :; fi