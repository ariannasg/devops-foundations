#!/usr/bin/env bash

export AWS_REGION=us-east-1

aws --region $AWS_REGION cloudformation validate-template --template-body file://./ec2.yml
aws --region $AWS_REGION cloudformation create-stack --stack-name myec2demo --parameters file://./ec2-parameters.yml --template-body file://./ec2.yml
