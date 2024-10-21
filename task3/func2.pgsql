-- Подставляемая табличная функция.

-- Вывести имена, телефоны и данные карточки клиентов, у которых они заблокированы
CREATE OR REPLACE FUNCTION BlockedCards(cid INT)
RETURNS TABLE (lf TEXT, phone VARCHAR(20), cnumber VARCHAR(16), cstatus card_status)
LANGUAGE PLPGSQL
AS $$
BEGIN
RETURN QUERY 
SELECT client.last_name || ' ' || client.first_name,
    client.phone_number,
    tcard.cnumber,
    tcard.cstatus
FROM (client JOIN account ON client.id = account.client_id) JOIN 
    (SELECT * FROM card WHERE card.cstatus = 'blocked') as tcard ON tcard.account_id = account.id;
END;
$$;

SELECT * FROM BlockedCards();