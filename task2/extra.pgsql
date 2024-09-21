DROP SCHEMA lab_02_extra CASCADE;
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
    (1, 'B', '2022-01-15', '2022-01-30'),
    (1, 'C', '2022-01-31', '2022-02-28');

INSERT INTO lab_02_extra.table2 (id, var2, valid_from_dttm, valid_to_dttm) VALUES
    (1, 'A', '2022-01-01', '2022-01-10'),
    (1, 'B', '2022-01-11', '2022-01-20'),
    (1, 'C', '2022-01-21', '2022-02-28');


SELECT
FROM (lab_02_extra.table1 JOIN lab_02_extra.table2 ON lab_02_extra.table1.id = lab_02_extra.table2.id) AS t;