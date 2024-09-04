# Check if an argument is provided
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <hostname>"
  exit 1
fi

# List services
if ! grpcurl -plaintext $1:50051 list wm_grpc.WMService; then
  echo "Error: could not list functions"
  exit 1
fi

# Example: Describe a service
# wm_grpc.WeightManager/AddEntry
# wm_grpc.WeightManager/GetEntries
# wm_grpc.WeightManager/UpdateEntry
# wm_grpc.WeightManager/DeleteEntry
grpcurl -plaintext $1:50051 describe wm_grpc.WMService.AnalyzeWeight

echo "{}" | grpcurl -plaintext -d @ $1:50051 wm_grpc.WMService.AnalyzeWeight
