#!/bin/bash

# ./run_all.sh data/example_1.txt

# Check if a filename is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <filename>"
    exit 1
fi

# Define temporary files to store output
NAIVE_OUTPUT="naive_output.txt"
SWISS_OUTPUT="swiss_output.txt"
WF_OUTPUT="wf_output.txt"

# Run naive/brute script
echo "Running naive/brute script..."
go run . -fileName=$1 --useNaive=True --caseSensitive=False > $NAIVE_OUTPUT
echo "Finished running naive/brute script."

echo "Sleeping ..."  && sleep 5

# Run SwissTable script
echo "Running SwissTable script..."
go run . -fileName=$1 --useNaive=False --caseSensitive=False > $SWISS_OUTPUT
echo "Finished running SwissTable script."

echo "Sleeping ..."  && sleep 5

# Run word_frequency script
echo "Running word_frequency script..."
bash word_frequency.sh $1 > $WF_OUTPUT
echo "Finished running word_frequency script."

# Process outputs into a table format
echo -e "\nResults in Table Format:\n"
paste $NAIVE_OUTPUT $SWISS_OUTPUT $WF_OUTPUT | column -t

# Cleanup temporary files
rm -f $NAIVE_OUTPUT $SWISS_OUTPUT $WF_OUTPUT
