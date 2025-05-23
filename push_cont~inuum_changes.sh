#!/bin/bash
set -e

# Ensure we are in the correct directory
cd "$(dirname "$0")/.."

# Add all changes
git add .

# Commit with a standard message
git commit -m "Automated commit from ChatGPT via commit helper"

# Push using the provided token
git push https://$GITHUB_TOKEN@github.com/AndrewWhiteside/continuum-core.git HEAD:main
