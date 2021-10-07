#!/bin/bash

## Resources

# Qeue (SQS)
awslocal sqs create-queue --queue-name consumer

# DynamoDB
awslocal dynamodb create-table \
  --table-name users \
  --attribute-definitions AttributeName=id,AttributeType=S \
  --key-schema AttributeName=id,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=10

# Function (Lambda)
awslocal lambda create-function \
  --function-name user_gen_password \
  --runtime go1.x \
  --memory-size 128 \
  --zip-file fileb://main.zip \
  --region us-east-1 \
  --handler bin/main \
  --role arn:aws:iam::000000000000:role/irrelevant \
  --timeout 10 \
  --environment "Variables={SERVER_ENVIRONMENT=development,DEFAULT_REGION=us-east-1,DYNAMODB_AWS_ENDPOINT=http://172.16.0.11:4566,TABLE_NAME=users}"

# Event Trigger
awslocal lambda create-event-source-mapping \
    --function-name user_gen_password \
    --batch-size 1 \
    --event-source-arn arn:aws:sqs:us-east-1:000000000000:consumer
