FROM ubuntu:noble
RUN apt update && apt install -y postgresql sudo

RUN mkdir /usr/local/postgres
RUN mkdir /usr/local/postgres/data


RUN sudo ln -s /usr/lib/postgresql/16/bin/initdb /bin/initdb
RUN sudo ln -s /usr/lib/postgresql/16/bin/postgres /bin/postgres
RUN sudo ln -s /usr/lib/postgresql/16/bin/pg_ctl /bin/pg_ctl

RUN chown postgres /usr/local/postgres/data
RUN sudo -u postgres initdb /usr/local/postgres/data
RUN echo "listen_addresses = '*'" >> /usr/local/postgres/data/postgresql.conf
RUN echo "port = 5432" >> /usr/local/postgres/data/postgresql.conf
RUN echo "host all impi 172.17.0.0/16 md5" >> /usr/local/postgres/data/pg_hba.conf
RUN echo "host all impi 172.18.0.0/16 md5" >> /usr/local/postgres/data/pg_hba.conf
RUN echo "host all impi 172.20.0.1/16 md5" >> /usr/local/postgres/data/pg_hba.conf
RUN sudo -u postgres pg_ctl -D /usr/local/postgres/data start && psql -U postgres -c "CREATE ROLE impi WITH LOGIN PASSWORD 'imp'"
RUN sudo -u postgres pg_ctl -D /usr/local/postgres/data start && psql -U postgres -c "ALTER ROLE impi WITH SUPERUSER "
RUN sudo -u postgres pg_ctl -D /usr/local/postgres/data start && psql -U postgres -c "CREATE DATABASE ibank WITH OWNER = 'impi'"
# COPY ./addhost.sh /addhost.sh
# RUN chmod 777 /addhost.sh