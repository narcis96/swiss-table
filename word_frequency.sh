#!/bin/bash

# Check if a filename is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <filename>"
    exit 1
fi

# Ensure the file exists and is readable
if [ ! -f "$1" ] || [ ! -r "$1" ]; then
    echo "Error: Cannot read file '$1'"
    exit 2
fi

# Process the file: extract words, count frequency, sort, and show the top 20
cat "$1" | tr -cs '[:alnum:]' '\n' | tr '[:upper:]' '[:lower:]' | sort | uniq -c | sort -nr | head -20
# cat $1 | tr " .,();{}[]" "\n" | sort | grep -v "^$" | uniq -c | sort -nr | head -20