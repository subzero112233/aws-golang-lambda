#!/bin/bash

set -e

S3_ARTIFACT_BUCKET=$1
FUNCTION_NAME="aws-golang-lambda-apig-authorizer"

export STACK_NAME=$FUNCTION_NAME

# prepare
mkdir -p dist/
cp lambda_handler dist/
cp template.yaml dist/

# package
sam package --metadata function=$FUNCTION_NAME --s3-prefix $FUNCTION_NAME --template-file dist/template.yaml --s3-bucket $S3_ARTIFACT_BUCKET --output-template-file dist/packaged.yaml

# deploy
sam deploy --template-file dist/packaged.yaml --stack-name "$STACK_NAME" --capabilities CAPABILITY_IAM

exit 0
