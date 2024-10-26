FROM ubuntu:noble
RUN apt update && apt install -y postgresql sudo postgresql-plpython3

RUN mkdir /usr/local/postgres
RUN mkdir /usr/local/postgres/data

COPY ./preparedb.sh preparedb.sh
RUN chmod 777 preparedb.sh
COPY ./ibank/ /usr/local/postgres/ibank

RUN ./preparedb.sh