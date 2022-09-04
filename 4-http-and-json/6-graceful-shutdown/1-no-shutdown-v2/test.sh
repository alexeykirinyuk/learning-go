#!/bin/bash

while true
do
    curl "http://localhost:8080/slow" || sleep 1
done