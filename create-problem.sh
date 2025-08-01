#!/bin/bash

# Script to create a new data structure problem using Claude CLI
# Usage: ./create-problem.sh <difficulty> <data-structure>
# Example: ./create-problem.sh intermediate linked-lists

# Check if correct number of arguments provided
if [ $# -ne 2 ]; then
    echo "Usage: $0 <difficulty> <data-structure>"
    echo ""
    echo "Difficulty levels: beginner, intermediate, advanced, senior"
    echo "Data structures: arrays, linked-lists, stacks, queues, binary-trees, etc."
    echo ""
    echo "Example: $0 intermediate linked-lists"
    exit 1
fi

DIFFICULTY=$1
DATA_STRUCTURE=$2

# Validate difficulty level
case $DIFFICULTY in
    beginner|intermediate|advanced|senior)
        ;;
    *)
        echo "Error: Invalid difficulty level '$DIFFICULTY'"
        echo "Valid options: beginner, intermediate, advanced, senior"
        exit 1
        ;;
esac

# Construct the prompt
PROMPT="Do not create a todo list, use steps.md verbatim. Create a new $DIFFICULTY level problem for $DATA_STRUCTURE."


echo "Creating $DIFFICULTY level problem for $DATA_STRUCTURE..."
echo "Prompt: $PROMPT"
echo ""

# Call Claude CLI with the constructed prompt
claude "$PROMPT"
