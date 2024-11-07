DROP TABLE IF EXISTS client_info;
CREATE TABLE IF NOT EXISTS client_info (
  id SERIAL PRIMARY KEY,
	client_id  INT,
  data JSONB
);

INSERT INTO client_info (client_id, data)
VALUES
(1, '{"mom":"big boy", "achievements":[], "education":"primary", "iq":72}'),
(2, '{"mom":"best son", "achievements":["big brain conference 2023", "best pancakes with mom 2022"], "education":"doctorate", "iq":148}'),
(3, '{"achievements":["russian_bear 2020", "british bulldog 2021"], "education":"master", "iq":110}');

SELECT * FROM client_info;

-- 1.Извлечь XML/JSON фрагмент из XML/JSON документа
SELECT 
  first_name,
  last_name,
  data->'achievements' as achievements
FROM client_info JOIN client ON client.id = client_info.client_id;

-- 2. Извлечь значения конкретных узлов или атрибутов XML/JSON документа

SELECT 
  first_name,
  last_name,
  data->'iq' as iq
FROM client_info JOIN client ON client.id = client_info.client_id;

-- 3. Выполнить проверку существования узла или атрибута

SELECT 
  first_name,
  last_name,
  data->>'mom' as mom_comment
FROM client_info JOIN client ON client.id = client_info.client_id
WHERE data->>'mom' IS NOT NULL;

-- 4. Изменить XML/JSON документ

UPDATE client_info
SET data = jsonb_set(data, '{mom}', '"not proud of him"', TRUE)
WHERE data->>'mom' IS NULL;

SELECT 
  first_name,
  last_name,
  data->>'mom' as mom_comment
FROM client_info JOIN client ON client.id = client_info.client_id
WHERE data->>'mom' IS NOT NULL;

-- 5. Разделить XML/JSON документ на несколько строк по узлам

WITH tmp AS (
  SELECT data->'achievements' as arr FROM client_info
)
SELECT jsonb_array_elements_text(tmp.arr) FROM tmp;








