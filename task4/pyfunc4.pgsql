-- Хранимая процедура

--- Досрочное погашение долга с пересчётом
CREATE OR REPLACE PROCEDURE SetPassword (cid INT, passwd TEXT)
  LANGUAGE plpython3u
AS $$
import hashlib
salt = "plpython3u"
hash = hashlib.md5((passwd + salt).encode())
hashString = hash.hexdigest()
quer = "UPDATE app_user SET hashpassword=$1 WHERE client_id = $2"
plpy.prepare(quer, ["text", "int"]).execute([hashString, cid])
$$;

SELECT *
FROM app_user
WHERE client_id = 5;

CALL SetPassword(5, 'dsdw');

SELECT *
FROM app_user
WHERE client_id = 5;
