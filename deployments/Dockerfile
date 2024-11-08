FROM ubuntu:noble
RUN apt update && apt install -y postgresql sudo postgresql-plpython3

RUN mkdir /usr/local/postgres
RUN mkdir /usr/local/postgres/data

COPY ./preparedb.sh preparedb.sh
COPY ./postgres.env postgres.env
RUN chmod 777 preparedb.sh

RUN ./preparedb.sh