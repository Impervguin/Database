LANGUAGE PLPGSQL
AS $$
DECLARE tabl information_schema.tables;
DECLARE cur_schema CURSOR FOR SELECT * FROM information_schema.tables WHERE table_schema = scheme_name;
BEGIN
  OPEN cur_schema;
  LOOP
    FETCH cur_schema INTO tabl;
    EXIT WHEN NOT FOUND;
    EXECUTE 'CREATE TABLE ' || scheme_name || '.' || tabl.table_catalog || '_' || tabl.table_name || '_' || to_char(now(), 'YYYYMMDD');
  END LOOP;
  CLOSE cur_schema;
END;
$$;

CALL BACKUP_SCHEME('public');