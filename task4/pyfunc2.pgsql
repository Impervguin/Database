-- Агрегатная функция CLR
CREATE OR REPLACE FUNCTION _Sortedinsert (vals NUMERIC(20, 2)[], new_val NUMERIC(20, 2))
  RETURNS NUMERIC(20, 2)[]
  LANGUAGE plpython3u
AS $$
left, right = 0, len(vals)
    
while left < right:
  mid = (left + right) // 2
  if vals[mid] < new_val:
    left = mid + 1
  else:
    right = mid
if (len(vals) == left):
	vals.append(new_val)
else:
	vals.insert(left, new_val)
return vals
$$;

-- Функция ищет медиану Numeric(20, 2)
CREATE OR REPLACE FUNCTION _Median (vals NUMERIC(20, 2)[])
  RETURNS NUMERIC(20, 2)
  LANGUAGE plpython3u
AS $$
n = len(vals)
if n == 0:
	return 0
if n % 2 == 0:
	return (vals[n // 2] + vals[n // 2 - 1]) / 2
return vals[n // 2]
$$;

CREATE OR REPLACE AGGREGATE Median (
  SFUNC = _Sortedinsert,
  BASETYPE = NUMERIC(20, 2),
  STYPE = NUMERIC(20, 2)[],
  INITCOND = '{}',
  FINALFUNC = _Median
);

SELECT client_id, Median(service.fee), AVG(service.fee), SUM(service.fee), COUNT(service.fee)
FROM client_service JOIN service ON service.id = service_id
GROUP BY client_id
ORDER BY client_id;

