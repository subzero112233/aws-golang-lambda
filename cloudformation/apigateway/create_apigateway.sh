#!/bin/bash

# create the stack
aws cloudformation deploy --template-file apig.yaml --stack-name aws-lambda-go-apigateway --capabilities CAPABILITY_IAM

exit_status=$?
if [ $exit_status -eq 1 ]; then
    echo "failed to create api gateway"
    exit 1
fi

echo "\n"

# output the endpoint
aws cloudformation describe-stacks --stack-name aws-lambda-go-apigateway --query "Stacks[0].Outputs[?OutputKey=='Endpoint'].OutputValue" --output text
