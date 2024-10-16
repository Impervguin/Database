-- Рекурсивная хранимая процедура

CREATE OR REPLACE PROCEDURE InterestUntilRich(account_id INT, rich NUMERIC(20, 2))
LANGUAGE PLPGSQL
AS $$
BEGIN
IF (SELECT balance FROM account WHERE account.id = account_id) < rich THEN
    UPDATE account SET balance = balance + balance * interest / 100 WHERE account.id = account_id;
    CALL InterestUntilRich(account_id, rich);
END IF;
END;
$$;

UPDATE account SET balance = 254 WHERE account.id = 0;
SELECT * FROM account WHERE account.id = 0;
CALL InterestUntilRich(0, 25000);
SELECT * FROM account WHERE account.id = 0;

