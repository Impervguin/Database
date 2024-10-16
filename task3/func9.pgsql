-- Триггер AFTER

CREATE OR REPLACE FUNCTION delete_transaction()
RETURNS TRIGGER
LANGUAGE PLPGSQL
AS $$
BEGIN
    IF OLD.ttype = 'withdraw' THEN
        UPDATE account SET balance = balance + OLD.amount WHERE account.id = OLD.account_id;
    ELSE
        UPDATE account SET balance = balance - OLD.amount WHERE account.id = OLD.account_id;
    END IF;
    RETURN NULL;
END;
$$;

DROP TRIGGER IF EXISTS after_delete_transaction;

CREATE TRIGGER after_delete_transaction
AFTER DELETE ON transaction
FOR EACH ROW
EXECUTE FUNCTION delete_transaction();

SELECT * FROM account
WHERE id = 253;

DELETE FROM transaction WHERE id = 1;

SELECT * FROM account
WHERE id = 253;