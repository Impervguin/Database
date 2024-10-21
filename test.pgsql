SELECT *
FROM client JOIN 
(SELECT client_id, COUNT(*) AS cnt FROM client_service Group BY client_id) AS tmp ON tmp.client_id = client.id
WHERE tmp.cnt = MAX()