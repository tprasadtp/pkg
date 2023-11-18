#!/bin/bash

set -eo pipefail

printf "Set git user.email\n"
git config user.email 151254910+remote-code-sync[bot]@users.noreply.github.com

printf "Set git user.name\n"
git config user.name remote-code-sync[bot]