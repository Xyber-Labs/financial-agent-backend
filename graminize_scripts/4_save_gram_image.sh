#!/bin/bash

GRAM_IMAGE_NAME=gsc-financial-agent-backend
GRAM_IMAGE_OUTPUT_FILE=gsc-financial-agent-backend.gz
echo "Saving docker image $GRAM_IMAGE_NAME to $GRAM_IMAGE_OUTPUT_FILE"
docker image save $GRAM_IMAGE_NAME | gzip -9 - > $GRAM_IMAGE_OUTPUT_FILE

echo "Saved $GRAM_IMAGE_NAME to $GRAM_IMAGE_OUTPUT_FILE"
