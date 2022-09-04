#!/bin/bash

# -v for response headers
RESP=$(curl -v "http://localhost:8080/heart")
echo "$RESP"