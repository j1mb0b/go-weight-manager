CREATE KEYSPACE weight_manager
WITH REPLICATION = {
  'class' : 'SimpleStrategy',
  'replication_factor' : 1
};