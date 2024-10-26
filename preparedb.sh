#!/bin/bash

ln -s /usr/lib/postgresql/16/bin/initdb /bin/initdb
ln -s /usr/lib/postgresql/16/bin/postgres /bin/postgres
ln -s /usr/lib/postgresql/16/bin/pg_ctl /bin/pg_ctl

chown postgres /usr/local/postgres/data
sudo -u postgres initdb /usr/local/postgres/data

echo "listen_addresses = '*'" >> /usr/local/postgres/data/postgresql.conf
echo "port = 5432" >> /usr/local/postgres/data/postgresql.conf
echo "host all impi 0.0.0.0/0 md5" >> /usr/local/postgres/data/pg_hba.conf

sudo -u postgres pg_ctl -D /usr/local/postgres/data start

psql -U postgres -c "CREATE ROLE impi WITH LOGIN PASSWORD 'imp'"
psql -U postgres -c "ALTER ROLE impi WITH SUPERUSER "
psql -U postgres -c "CREATE DATABASE ibank WITH OWNER = 'impi'"