#!/bin/bash

#!/bin/bash

GRAM_IMAGE=gsc-financial-agent-backend.gz
SSH_HOST=tee
echo "running scp $GRAM_IMAGE $SSH_HOST:/root/xyber-financial-agent/"
scp $GRAM_IMAGE "$SSH_HOST:/root/xyber-financial-agent/"