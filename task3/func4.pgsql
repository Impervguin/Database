-- Рекурсивная функция
-- DROP FUNCTION RecursiveReviewTree;

DROP TABLE IF EXISTS service_review;
CREATE TABLE IF NOT EXISTS service_review (
    id SERIAL PRIMARY KEY,
    service_id INTEGER REFERENCES service(id),
    grade INTEGER CHECK (grade BETWEEN 1 AND 5),
    review text,
    on_review INTEGER DEFAULT NULL
);


INSERT INTO service_review (service_id, grade, review, on_review)
VALUES (5, 5, 'Отличный сервис!', NULL),
(5, 5, 'Согласен', 1),
(5, 1, 'Не согласен', 2);

CREATE OR REPLACE FUNCTION RecursiveReviewTree(current_level INT, review_id INT)
RETURNS TABLE (rid INT, serviceid INT, rgrade INT, rreview TEXT, onreview INT, onlevel INT)
LANGUAGE PLPGSQL
AS $$
BEGIN
RETURN QUERY
WITH RECURSIVE direct_review (id, on_review, service_id, grade, review, on_level) AS (
    SELECT sr.id, on_review, service_id, grade, review, current_level on_level
    FROM service_review as sr
    WHERE sr.id = review_id
    UNION ALL
    SELECT sr.id, sr.on_review, sr.service_id, sr.grade, sr.review, on_level + 1
    FROM service_review as sr JOIN direct_review as dr ON sr.on_review = dr.id
) 
SELECT id, service_id, grade, review, on_review, on_level
FROM direct_review;
END;
$$;

SELECT * FROM RecursiveReviewTree(0, 1);