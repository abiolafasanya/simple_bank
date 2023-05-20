#!/bin/bash

# Add the file to the staging area
git add .

# Show the status of the repository
git status

# Prompt the user for a commit message
read -p "Enter a commit message: " message

# Create a commit with the given message
git commit -m "$message"

# Push the commit to the remote branch
git push upstream
