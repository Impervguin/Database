-- Многооператорная функция

-- Собрать статистику по клиенту: месячный доход по счетам, месячная выплата по долгам, сервисам, остаток на счетах

-- DROP FUNCTION ClientMonthly;

CREATE OR REPLACE FUNCTION ClientMonthly(cid INT)
RETURNS TABLE (clientid int, lf TEXT, phone VARCHAR(20), montly_income NUMERIC(20, 2), monthly_loan NUMERIC(20, 2), monthly_service NUMERIC(20, 2), total_balance NUMERIC(20, 2))
LANGUAGE PLPGSQL
AS $$
DECLARE montly_income NUMERIC(20, 2);
DECLARE monthly_loan NUMERIC(20, 2);
DECLARE monthly_service NUMERIC(20, 2);
DECLARE total_balance NUMERIC(20, 2);
BEGIN
    SELECT SUM(balance)
    INTO total_balance
    FROM account
    WHERE account.client_id = cid AND account.atype <> 'credit' AND account.astatus = 'active';

    SELECT SUM(loan.monthly_payment)
    INTO monthly_loan
    FROM loan
    WHERE loan.client_id = cid AND loan.lstatus = 'active';

    SELECT SUM(service.fee)
    INTO monthly_service
    FROM service JOIN (SELECT service_id FROM client_service WHERE client_id = cid) as ts ON ts.service_id = service.id;

    SELECT SUM(balance * interest / 100)
    INTO montly_income
    FROM account
    WHERE account.client_id = cid AND account.atype <> 'credit' AND account.astatus = 'active';

    RETURN QUERY
    SELECT client.id,
    client.last_name || ' ' || client.first_name,
    client.phone_number,
    montly_income,
    monthly_loan,
    monthly_service,
    total_balance
    FROM client
    WHERE client.id = cid;
END;
$$;

SELECT * FROM ClientMonthly(1);
