#!/bin/bash

# Check if urls.txt file doesn't exist
if [ ! -f ./urls.txt ]; then
  # Use LC_ALL=C to handle non-ASCII characters
  export LC_ALL=C

  for i in $(seq 1 10000); do
      # Generate random user data
      id=$RANDOM
      name=$(cat /dev/urandom | tr -dc 'a-zA-Z' | fold -w 10 | head -n 1)
      email="${name}@example.com"

      # Save email to file
      echo "http://localhost:8088/app/users/create/${email}" >> ./urls.txt
  done
else
  echo "urls.txt file exists. Starting siege..."
fi

siege -f ./urls.txt -c 100 -j -v -r 1000
