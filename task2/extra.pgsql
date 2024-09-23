DROP SCHEMA IF EXISTS lab_02_extra CASCADE;
CREATE SCHEMA lab_02_extra;

CREATE TABLE IF NOT EXISTS lab_02_extra.table1(
    id int,
    var1 text,
    valid_from_dttm date,
    valid_to_dttm date
);

CREATE TABLE IF NOT EXISTS lab_02_extra.table2(
    id int,
    var2 text,
    valid_from_dttm date,
    valid_to_dttm date
);

INSERT INTO lab_02_extra.table1 (id, var1, valid_from_dttm, valid_to_dttm) VALUES
    (1, 'A', '2022-01-01', '2022-01-15'),
    (1, 'B', '2022-01-16', '2022-01-30'),
    (1, 'C', '2022-01-31', '2022-02-28'),
    (2, 'A', '2022-01-01', '2022-01-31'),
    (2, 'G', '2022-02-01', '2022-02-13'),
    (2, 'H', '2022-02-14', '2022-03-02'),
    (2, 'E', '2022-03-03', '2022-03-31');

INSERT INTO lab_02_extra.table2 (id, var2, valid_from_dttm, valid_to_dttm) VALUES
    (1, 'A', '2022-01-01', '2022-01-10'),
    (1, 'B', '2022-01-11', '2022-01-20'),
    (1, 'C', '2022-01-21', '2022-02-28'),
    (2, 'J', '2022-01-01', '2022-02-28'),
    (2, 'K', '2022-03-01', '2022-03-31');


SELECT t.id1, t.var1, t.var2, GREATEST(t.valid_from1, t.valid_from2) AS valid_from, LEAST(t.valid_to1, t.valid_to2) AS valid_to
FROM (
    (SELECT id AS id1, var1, valid_from_dttm AS valid_from1, valid_to_dttm AS valid_to1 FROM lab_02_extra.table1) AS t1
    JOIN 
    (SELECT id AS id2, var2, valid_from_dttm AS valid_from2, valid_to_dttm AS valid_to2 FROM lab_02_extra.table2) AS t2
    ON t1.id1 = t2.id2
    ) AS t
WHERE GREATEST(t.valid_from1, t.valid_from2) < LEAST(t.valid_to1, t.valid_to2)
ORDER BY t.id1, valid_from