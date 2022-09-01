#!/bin/bash

JSON=$(curl -s "https://gorest.co.in/public/v1/users?name=Agrata")

echo "$JSON" | jq

echo "$JSON" | jq '.data[] | {email: .email, name: .name}'