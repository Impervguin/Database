COPY (SELECT array_to_json(array_agg(row_to_json(tc))) FROM transaction as tc)
TO '/usr/local/postgres/ibank/transaction.json';

SELECT row_to_json(tc) FROM transaction as tc;

DELETE FROM transaction;

SELECT * FROM transaction;


CREATE TEMP TABLE transactionsjson(data json) ON COMMIT DROP;

COPY transactionsjson FROM '/usr/local/postgres/ibank/transaction.json';

INSERT INTO transaction
SELECT (json_populate_record(null::transaction, json_array_elements(data))).* FROM transactionsjson;

SELECT * FROM transaction
ORDER BY id;














