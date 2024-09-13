COPY client FROM '/usr/local/postgres/ibank/client.csv' DELIMITER ',' HEADER;
COPY account FROM '/usr/local/postgres/ibank/account.csv' DELIMITER ',' HEADER;
COPY card FROM '/usr/local/postgres/ibank/card.csv' DELIMITER ',' HEADER;
COPY loan FROM '/usr/local/postgres/ibank/loan.csv' DELIMITER ',' HEADER;
COPY transaction FROM '/usr/local/postgres/ibank/transaction.csv' DELIMITER ',' HEADER;
COPY service FROM '/usr/local/postgres/ibank/service.csv' DELIMITER ',' HEADER;
COPY client_service FROM '/usr/local/postgres/ibank/clientservice.csv' DELIMITER ',' HEADER;
COPY app_user FROM '/usr/local/postgres/ibank/user.csv' DELIMITER ',' HEADER;
COPY notification FROM '/usr/local/postgres/ibank/notification.csv' DELIMITER ',' HEADER;

