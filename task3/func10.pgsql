-- Триггер INSTEAD OF

CREATE OR REPLACE FUNCTION insert_account()
RETURNS TRIGGER
LANGUAGE PLPGSQL
AS $$
BEGIN
    IF (SELECT id FROM client WHERE id = NEW.client_id) IS NULL THEN
        RETURN NULL;
    ELSIF NEW.created_at < (SELECT created_at FROM client WHERE id = NEW.client_id) THEN
        RETURN NULL;
    ELSE
        INSERT INTO account (id, client_id, balance, interest, created_at, atype, astatus)
        VALUES (
            NEW.id,
            NEW.client_id,
            NEW.balance,
            NEW.interest,
            NEW.created_at,
            NEW.atype,
            NEW.astatus
        );
        RETURN NEW;
    END IF;
END;
$$;

DROP VIEW IF EXISTS account_view;
CREATE VIEW  account_view
AS SELECT * FROM account;

-- DROP TRIGGER instead_insert_account;

CREATE OR REPLACE TRIGGER instead_insert_account
INSTEAD OF INSERT ON account_view
FOR EACH ROW
EXECUTE FUNCTION insert_account();

INSERT INTO account_view (id, client_id, balance, interest, created_at, atype, astatus)
VALUES
(1008, 253, 10000, 10, (SELECT created_at + interval '1 day' FROM client WHERE id= 253), 'checking', 'active');

INSERT INTO account_view (id, client_id, balance, interest, created_at, atype, astatus)
VALUES
(1009, 253, 10000, 10, (SELECT created_at - interval '1 day' FROM client WHERE id= 253), 'checking', 'active');

SELECT * FROM account_view WHERE client_id = 253;

DELETE FROM account WHERE id = 1008;