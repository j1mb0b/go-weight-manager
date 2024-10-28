#!/bin/bash
set -e

exec /usr/local/bin/docker-entrypoint.sh cassandra -f

echo "Waiting for Cassandra to start..."
until cqlsh -e "DESCRIBE KEYSPACES" > /dev/null 2>&1; do
  echo "Still waiting for Cassandra..."
  sleep 5
done

echo "Cassandra is up, executing CQL script..."
if cqlsh -f /docker-entrypoint-initdb.d/init.cql; then
  echo "CQL script executed successfully."
else
  echo "Failed to execute CQL script."
  exit 1
fi
