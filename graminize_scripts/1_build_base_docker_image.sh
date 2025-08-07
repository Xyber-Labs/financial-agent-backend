#!/bin/bash

# Build base docker image from Dockerfile in the repo's root
docker build -t financial-agent-backend-base -f ../Dockerfile ..