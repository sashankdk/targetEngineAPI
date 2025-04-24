#!/bin/bash

# Step 1: Build and run only the test container
sudo docker compose up --build --exit-code-from test test

# Step 2: Check if test container exited successfully
if [ $? -eq 0 ]; then
  echo "Tests passed, starting remaining services..."
  sudo docker compose up -d api redis postgres
else
  echo "Tests failed. Not starting app containers."
fi
