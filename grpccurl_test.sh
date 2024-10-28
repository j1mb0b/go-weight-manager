#!/bin/bash

# Check if an argument is provided
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <hostname>"
  exit 1
fi

HOSTNAME=$1

# List services
if ! grpcurl -plaintext $HOSTNAME:50051 list wm_grpc.WMService; then
  echo "Error: could not list functions"
  exit 1
fi

# Add an entry
echo "Adding a weight entry..."

# Prepare the JSON payload for the weight entry
ENTRY_PAYLOAD='{
  "uid": 1,
  "date": "2024-09-01",
  "weight": 75
}'

# Make the gRPC call to add the entry
ADD_RESPONSE=$(echo "$ENTRY_PAYLOAD" | grpcurl -plaintext -d @ $HOSTNAME:50051 wm_grpc.WMService.AddEntry)

echo "Add Entry Response: $ADD_RESPONSE"

# Now, retrieve entries to verify the addition
echo "Retrieving entries..."

GET_RESPONSE=$(grpcurl -plaintext -d '{"uid": 1}' $HOSTNAME:50051 wm_grpc.WMService.GetEntries)

echo "Get Entries Response: $GET_RESPONSE"
