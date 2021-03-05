#!/bin/bash
for i in {0..255}; do
  printf "\e[48;5;%sm%-4s\e[0m" "$i" "$i"
done
echo ""
