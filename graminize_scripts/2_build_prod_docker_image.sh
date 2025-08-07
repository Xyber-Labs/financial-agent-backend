#!/bin/bash

# Build production docker image from Dockerfile.prod in the repo's root.
# Used to create configuration to be used in production graminized image.
docker build -t financial-agent-backend -f ./Dockerfile.prod .