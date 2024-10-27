CREATE OR REPLACE FUNCTION UpcaseName ()
  RETURNS TRIGGER
  LANGUAGE plpython3u
AS $$
TD["new"]["first_name"] = TD["new"]["first_name"].capitalize()
TD["new"]["last_name"] = TD["new"]["last_name"].capitalize()
return "MODIFY"
$$;

DROP TRIGGER IF EXISTS before_insert_client ON client;

CREATE TRIGGER before_insert_client
BEFORE INSERT ON client
FOR EACH ROW
EXECUTE FUNCTION UpcaseName ();

INSERT INTO client (id, first_name, last_name, dob, email, phone_number, address, created_at)
VALUES (1008, 'dmitriy', 'shakhnovich', '12-01-2004', 'somedsethindgd2d@ya.ru', '+75754651234', 'Russia, Yarik', now());

SELECT *
FROM client
WHERE id = 1008;