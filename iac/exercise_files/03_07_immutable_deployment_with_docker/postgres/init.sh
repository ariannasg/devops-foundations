#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE assetdb;
    CREATE DATABASE quartz;
    GRANT ALL PRIVILEGES ON DATABASE assetdb TO test;
    GRANT ALL PRIVILEGES ON DATABASE quartz TO test;
EOSQL
