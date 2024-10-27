DROP FUNCTION IF EXISTS GetServiceStatistics;
DROP TYPE IF EXISTS service_statistics;
CREATE TYPE service_statistics AS (
  sid int,
  sname   text,
  sdescription  text,
  fee NUMERIC(20, 2),
  total_fee NUMERIC(20, 2),
  users INT
);

CREATE OR REPLACE FUNCTION GetServiceStatistics (sid INT)
  RETURNS service_statistics
  LANGUAGE plpython3u
AS $$
class serv_stat:
  def __init__(self, sid):
    self.sid = sid
    quar = '''
    SELECT service.id, sname, sdescription, fee, COUNT(*) users
    FROM service JOIN client_service  ON service.id = client_service.service_id
    WHERE service.id = $1
    GROUP BY service.id, sname, sdescription, fee;
    '''
    res = plpy.prepare(quar, ["int"]).execute([sid])
    if res.nrows() == 0:
    	return
    rw = res[0]
    self.sname = rw['sname']
    self.sdescription = rw['sdescription']
    self.fee = rw['fee']
    self.users = rw['users']
    self.total_fee = self.users*self.fee
return serv_stat(sid)
$$;

SELECT * FROM GetServiceStatistics(1);


