
-- Вывести уведомления клиента, по введённой карте
CREATE OR REPLACE FUNCTION CardNotification(cid INT)
RETURNS TABLE (
    nid INT,
    card_id INT,
    nmessage text,
    created_at TIMESTAMP,
    seen BOOLEAN
)

LANGUAGE PLPGSQL
AS $$
BEGIN
RETURN QUERY 
SELECT 
    notification.id,
    tcard.id,
    notification.nmessage,
    notification.created_at,
    notification.seen
FROM (((notification JOIN client ON notification.client_id = client.id)
    JOIN account ON client.id = account.client_id)
    JOIN (SELECT id, account_id FROM card WHERE card.id = cid) as tcard ON tcard.account_id = account.id); 
END;
$$;

-- INSERT INTO notification (client_id, nmessage, created_at, seen)
-- VALUES
-- (642, 'done 1000 рублей', '2021-02-02', true),
-- (641, 'done 2000 рублей', '2021-02-03', true),
-- (642, 'done 3000 рублей', '2021-02-04', false),
-- (641, 'got 1000 рублей', '2021-02-02', false),
-- (642, 'got 2000 рублей', '2021-02-03', false),
-- (641, 'got 3000 рублей', '2021-02-04', false);

SELECT client.id as clientid, card.id as cardid
FROM client JOIN account ON client.id = account.client_id JOIN card ON card.account_id = account.id
WHERE client.id = 641;

SELECT * FROM CardNotification(104);
