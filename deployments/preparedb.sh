#!/bin/bash

source postgres.env

ln -s /usr/lib/postgresql/16/bin/initdb /bin/initdb
ln -s /usr/lib/postgresql/16/bin/postgres /bin/postgres
ln -s /usr/lib/postgresql/16/bin/pg_ctl /bin/pg_ctl

chown postgres /usr/local/postgres/data
sudo -u postgres initdb /usr/local/postgres/data

echo "listen_addresses = '*'" >> /usr/local/postgres/data/postgresql.conf
echo "port = 5432" >> /usr/local/postgres/data/postgresql.conf
echo "host all $POSTGRES_USER 0.0.0.0/0 md5" >> /usr/local/postgres/data/pg_hba.conf

sudo -u postgres pg_ctl -D /usr/local/postgres/data start

psql -U postgres -c "CREATE ROLE $POSTGRES_USER WITH LOGIN PASSWORD '$POSTGRES_PASSWORD'"
psql -U postgres -c "ALTER ROLE $POSTGRES_USER WITH SUPERUSER "
psql -U postgres -c "CREATE DATABASE $POSTGRES_DB WITH OWNER = '$POSTGRES_USER'"