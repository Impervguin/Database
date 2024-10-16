-- Хранимая процедура

-- Акция: прощение долга случайному человеку

CREATE OR REPLACE PROCEDURE PromoRandom()
LANGUAGE PLPGSQL
AS $$
DECLARE LUCKY INT;
DECLARE LOAN_ID INT;
BEGIN
SELECT id
INTO LOAN_ID
FROM loan
WHERE lstatus = 'active'
ORDER BY random() LIMIT 1;



SELECT client_id
INTO LUCKY
FROM loan
WHERE id = LOAN_ID;

RAISE INFO 'Lucky guy is % with loan %', LUCKY, LOAN_ID;

UPDATE loan 
SET lstatus = 'defaulted'
WHERE id = LOAN_ID;
END;
$$;

SELECT COUNT(*) FROM loan WHERE lstatus = 'active';

CALL PromoRandom();

SELECT COUNT(*) FROM loan WHERE lstatus = 'active';

