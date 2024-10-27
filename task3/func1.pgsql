-- Скалярная функция

-- Возвращает суммарный баланс человека на его счетах, без учёта кредитных

CREATE OR REPLACE FUNCTION TotalBalance(cid INT)
RETURNS NUMERIC(20, 2)
LANGUAGE PLPGSQL
AS $$
DECLARE
result NUMERIC(20, 2);
BEGIN
    SELECT SUM(balance)
    FROM account
    WHERE account.client_id = cid AND atype != 'credit';
    RETURN result;
END;
$$;

SELECT id, TotalBalance(id) FROM client;