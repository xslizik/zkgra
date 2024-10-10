#!/bin/bash

iterations=100
total_time=0

if [ -z "$1" ]; then
    echo "Usage: $0 <path_to_script>"
    exit 1
fi

script="$1"
go build "$script"

for ((i=1; i<=iterations; i++)); do
    elapsed_time=$( { time ./main; } 2>&1 | grep real | awk '{print $2}' )
    
    if [[ $elapsed_time == *m* ]]; then
        minutes=$(echo $elapsed_time | cut -d'm' -f1)
        seconds=$(echo $elapsed_time | cut -d'm' -f2 | sed 's/s//')
        elapsed_time=$(echo "$minutes * 60 + $seconds" | bc)
    else
        elapsed_time=$(echo $elapsed_time | sed 's/s//')
    fi
    
    total_time=$(echo "$total_time + $elapsed_time" | bc)
done

average_time=$(echo "$total_time / $iterations" | bc -l)

printf "Average real time over %d runs: %.6f seconds\n" "$iterations" "$average_time"
