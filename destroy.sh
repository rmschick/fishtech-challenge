#!/bin/bash

echo -e "----- Starting deletion -----\n"

echo -e "Deleting with Terraform...\n"

cd infrastructure
terraform destroy -input=false -auto-approve
cd ../

echo -e "Done...\n"
