#!/bin/bash
set -e

echo "Waiting for Cassandra to start..."
until cqlsh -e "DESCRIBE KEYSPACES" > /dev/null 2>&1; do
  sleep 5
done

echo "Cassandra is up, executing CQL script..."
cqlsh -f /docker-entrypoint-initdb.d/init.cql
