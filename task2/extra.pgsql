DROP SCHEMA IF EXISTS lab_02_extra CASCADE;
CREATE SCHEMA lab_02_extra;

CREATE TABLE IF NOT EXISTS lab_02_extra.table1(
    id int,
    var1 text,
    valid_from date,
    valid_to date
);

CREATE TABLE IF NOT EXISTS lab_02_extra.table2(
    id int,
    var2 text,
    valid_from date,
    valid_to date
);

INSERT INTO lab_02_extra.table1 (id, var1, valid_from, valid_to) VALUES
    (1, 'A', '2022-01-04', '2022-01-15'),
    (1, 'B', '2022-01-16', '2022-01-30'),
    (1, 'C', '2022-01-31', '2022-02-28'),

    (2, 'A', '2022-01-01', '2022-01-31'),
    (2, 'G', '2022-02-01', '2022-02-13'),
    (2, 'H', '2022-02-14', '2022-03-02'),
    (2, 'E', '2022-03-03', '2022-03-31');

INSERT INTO lab_02_extra.table2 (id, var2, valid_from, valid_to) VALUES
    (1, 'A', '2022-01-01', '2022-01-10'),
    (1, 'B', '2022-01-11', '2022-01-20'),
    (1, 'C', '2022-01-21', '2022-02-28'),
    
    (2, 'J', '2022-01-01', '2022-02-28'),
    (2, 'K', '2022-03-01', '2022-03-31'),
    
    (3, 'J', '2022-01-01', '2022-02-28');


SELECT coalesce(t1.id, t2.id) as id, t1.var1, t2.var2, 
       GREATEST(t1.valid_from, t2.valid_from) AS valid_from, 
       LEAST(t1.valid_to, t2.valid_to) AS valid_to
FROM lab_02_extra.table1 t1
     full JOIN lab_02_extra.table2 t2 ON t1.id = t2.id
WHERE GREATEST(t1.valid_from, t2.valid_from) < LEAST(t1.valid_to, t2.valid_to)
ORDER BY id, valid_from