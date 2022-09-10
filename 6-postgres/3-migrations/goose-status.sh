#!/bin/bash

goose -dir ./migrations \
  postgres "user=akirinyuk password=pwd host=localhost database=migrations_3 port=5432 sslmode=disable" \
  up