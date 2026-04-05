#!/bin/bash

# Project folder
WATCHED_DIR="./"

while true; do
  # Fayl o'zgarganda trigger bo'ladi
  inotifywait -r -e modify,create,delete $WATCHED_DIR
  git add -A
  git commit -m "Auto commit: $(date '+%Y-%m-%d %H:%M:%S')"
  git push
done
