#!/bin/bash

echo -e "\n+++++ Starting deployment +++++\n"

tfswitch 1.0.0

rm -rf ./bin

echo "+++++ build go packages +++++"

cd source/
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o main.go
cd ..

echo "+++++ ipaddress module +++++"
cd infrastructure/
terraform init -input=false
terraform apply -input=false -auto-approve
cd ../

echo -e "\n+++++ Deployment done +++++\n"