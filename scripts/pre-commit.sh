#!/bin/bash

echo "Running pre-commit hook. It worked if it prints OK."

echo "Removing compiled files"
make clean

echo "OK"
exit 0
