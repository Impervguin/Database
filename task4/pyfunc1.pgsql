-- Скалярная функция CLR

-- Функция возвращает id более богатого клиента из 2-х
CREATE OR REPLACE FUNCTION Richier (cl1 INT, cl2 INT)
  RETURNS INT
  LANGUAGE plpython3u
AS $$
quer = "SELECT SUM(balance) as sbalance FROM account WHERE account.client_id = $1 AND atype != 'credit';"
prep = plpy.prepare(quer, ["INT"])
clsum1 = prep.execute([cl1])
clsum2 = prep.execute([cl2])
if clsum1[0]['sbalance'] > clsum2[0]['sbalance']:
  return cl1
return cl2
$$;

SELECT SUM(balance) as sbalance, account.client_id as cid
FROM account
WHERE account.client_id = 24 OR account.client_id = 29 AND atype != 'credit'
GROUP BY account.client_id;

SELECT SUM(balance) as sbalance, account.client_id as cid
FROM account
WHERE account.client_id = Richier(24, 29)
GROUP BY account.client_id;








