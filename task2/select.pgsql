-- --- 1. Инструкция SELECT, использующая предикат сравнения.
-- --- Найти все невыплаченные займы, оставшаяся сумма которых превышает 1 000 000, а ежемесячный платёж больше 10 000
-- SELECT id, remaining_amount, monthly_payment, lstatus
-- FROM loan
-- WHERE lstatus = 'active'
-- AND monthly_payment >= 10000
-- AND remaining_amount >= 1000000
-- ORDER BY remaining_amount;


-- --- 2. Инструкция SELECT, использующая предикат BETWEEN.
-- --- Найти все транзакция сделанные в мае 2022 года.
-- SELECT id, done_at
-- FROM transaction
-- WHERE done_at BETWEEN '2022-05-01' AND '2022-05-31'
-- ORDER BY done_at;

-- --- 3. Инструкция SELECT, использующая предикат LIKE.
-- --- Найти аккаунты всех пользователей, у которых почта оканчивается на .com
-- SELECT app_user.id AS user_id, client.email
-- FROM (app_user JOIN client ON app_user.client_id = client.id)
-- WHERE email LIKE '%.com';

-- --- 4. Инструкция SELECT, использующая предикат IN с вложенным подзапросом.
-- --- Найти имена всех клиентов банка у которых в июне 2022 были транзакции снятия или перевода
-- SELECT client.id, client.last_name, client.first_name
-- FROM client
-- WHERE client.id IN (
--     SELECT account.client_id
--     FROM account JOIN transaction ON transaction.account_id = account.id
--     WHERE transaction.ttype = 'withdraw' OR (transaction.ttype = 'transfer' AND transaction.amount < 0) 
--     AND transaction.done_at BETWEEN '2022-06-01' AND '2022-06-30'
-- );

-- --- 5. Инструкция SELECT, использующая предикат EXISTS с вложенным подзапросом.
-- --- Найти клиентов которые не подписаны ни на один сервис
-- SELECT client1.id, client1.last_name, client1.first_name
-- FROM client AS client1
-- WHERE NOT EXISTS (
--     SELECT client_service.service_id
--     FROM client_service
--     WHERE client_service.client_id = client1.id
-- );

-- --- 6. Инструкция SELECT, использующая предикат сравнения с квантором.
-- --- Найти клиентов у которых ежемесечная плата по займу выше, чем остаток по всем счетам раздельно
-- SELECT client.id, client.last_name, client.first_name, loan.monthly_payment
-- FROM client JOIN loan ON client.id = loan.client_id
-- WHERE loan.monthly_payment > ALL (
--     SELECT account.balance
--     FROM account
--     WHERE account.client_id = client.id
-- );

-- --- 7. Инструкция SELECT, использующая агрегатные функции в выражениях столбцов. 
-- --- Найти количество счётов каждого вида и посчитать средний остаток по счёту на них
-- SELECT atype, acnt, avgbalance
-- FROM (
--     SELECT account.atype, COUNT(*) as acnt, AVG(account.balance) as avgbalance
--     FROM account
--     GROUP BY account.atype
--     );

-- --- 8.Инструкция  SELECT, использующая скалярные подзапросы в выражениях столбцов.
-- --- Для каждого сервиса найти количество клиентов использующих его, а также суммарный доход по сервису
-- SELECT id, (
--     SELECT COUNT(*)
--     FROM client_service
--     WHERE service_id = service.id
-- ) as scnt, fee,
-- (
--     SELECT COUNT(*) * service.fee
--     FROM client_service
--     WHERE service_id = service.id
-- ) as total_fee
-- FROM service;

-- 9. Инструкция SELECT, использующая простое выражение CASE. 
-- Выбрать богатеньких, средненьких и бедных клиентов, по суммарному балансу всех счетов, кроме кредитных
-- SELECT id, sumbalance, (
--     CASE
--         WHEN sumbalance < 100000 THEN 'smol bich'
--         WHEN sumbalance > 10000000 THEN 'rich bich(potential papik)'
--         ELSE 'average bich'
--     END
-- ) AS my_comment
-- FROM (SELECT id, (
--     SELECT SUM(balance)
--     FROM account
--     WHERE account.client_id = client.id
--     ) as sumbalance
--     FROM client
--     WHERE (
--         SELECT SUM(balance)
--         FROM account
--         WHERE account.client_id = client.id
--     ) IS NOT NULL
-- )


-- 10. Инструкция SELECT, использующая поисковое выражение CASE. 
-- Выбрать популярные, используемые и не используемые сервисы
-- SELECT id, clientcnt, (
--     CASE
--         WHEN clientcnt = 0 THEN 'unused'
--         WHEN clientcnt < 20 THEN 'used'
--         ELSE 'popular'
--     END
-- ) as metric
-- FROM (
--     SELECT id, (
--         SELECT COUNT(*)
--         FROM client_service
--         WHERE client_service.service_id = service.id
--     ) as clientcnt
--     FROM service
-- )

-- 11. Создание новой временной локальной таблицы из результирующего набора данных инструкции SELECT. 
--- Для каждого сервиса найти количество клиентов использующих его, а также суммарный доход по сервису
-- SELECT id, (
--     SELECT COUNT(*)
--     FROM client_service
--     WHERE service_id = service.id
-- ) as scnt, fee,
-- (
--     SELECT COUNT(*) * service.fee
--     FROM client_service
--     WHERE service_id = service.id
-- ) as total_fee
-- INTO tmp_services
-- FROM service;
-- DROP TABLE tmp_services;

-- 12. Инструкция SELECT, использующая вложенные коррелированные подзапросы в качестве производных таблиц в предложении FROM.
-- Получить всех клиентов, у которых суммарная месячная выплата больше остатка по не кредитным счетам
-- SELECT id, phone_number, sum_payment, sum_balance
-- FROM (client JOIN (
--     SELECT client_id, SUM(monthly_payment) sum_payment
--     FROM loan
--     WHERE loan.lstatus = 'active'
--     GROUP BY client_id
-- ) AS l ON client.id = l.client_id) as lc JOIN (
--     SELECT client_id, SUM(balance) sum_balance
--     FROM account
--     WHERE account.atype <> 'credit'
--     GROUP BY client_id
-- ) AS a ON a.client_id = lc.client_id
-- WHERE sum_balance < sum_payment;

-- 13. Инструкция SELECT, использующая вложенные подзапросы с уровнем вложенности 3. 
-- Получить всех клиентов, у которых есть хотя бы одна заблокированная карта.
-- SELECT id, phone_number
-- FROM client
-- WHERE id IN (
--     SELECT client_id
--     FROM account
--     WHERE id IN (
--         SELECT account_id
--         FROM card
--         WHERE card.cstatus = 'blocked'
--     )
-- )


-- 14. Инструкция  SELECT, консолидирующая данные с помощью предложения GROUP BY, но без предложения HAVING. 
-- SELECT a.client_id, ttype, COUNT(*)
-- FROM account a JOIN (
--     SELECT account_id, ttype, COUNT(*) cnt
--     FROM transaction
--     GROUP BY account_id, ttype
--     ORDER BY account_id, ttype
-- ) AS acc_type ON acc_type.account_id = a.id
-- GROUP BY client_id, ttype
-- ORDER BY client_id, ttype;

-- 15. Инструкция  SELECT, консолидирующая данные с помощью предложения GROUP BY и предложения HAVING.
-- Получить всех клиентов, у которых суммарная месячная выплата больше остатка по не кредитным счетам
-- SELECT client.id, SUM(balance) sum_balance
-- FROM client JOIN account ON client.id = account.client_id
-- WHERE atype <> 'credit'
-- GROUP BY client.id
-- HAVING SUM(balance) < (
--     SELECT SUM(monthly_payment)
--     FROM loan
--     WHERE loan.client_id = client.id AND loan.lstatus = 'active'
-- )


-- 16. Однострочная инструкция  INSERT, выполняющая вставку в таблицу одной строки значений.
-- INSERT INTO client (id, first_name, last_name, dob, email, phone_number, address, created_at)
-- VALUES (1001, 'Dmitriy', 'Shakhnovich', '12-01-2004', 'something@ya.ru', '+79854321234', 'Russia, Yarik', now());

-- 17. Многострочная инструкция INSERT, выполняющая вставку в таблицу результирующего набора данных вложенного подзапроса. 
-- INSERT INTO client_service (client_id, service_id)
-- SELECT client.id, 10
--     FROM client JOIN (
--         SELECT client_id, SUM(balance)
--         FROM account
--         GROUP BY client_id
--         ) as accs ON client.id = accs.client_id;

-- -- 18. Простая инструкция UPDATE. 
-- UPDATE service
-- SET fee = service.fee * 1.5
-- WHERE service.id = 12;

-- SELECT id, fee FROM service WHERE service.id = 12;

-- 19. Инструкция UPDATE со скалярным подзапросом в предложении SET. 
-- UPDATE account
-- SET balance = balance - (
--         SELECT SUM(monthly_payment)
--         FROM loan
--         WHERE loan.client_id = account.client_id AND loan.lstatus = 'active'
--     )
-- WHERE id = ANY(SELECT id FROM account WHERE client_id = 555);

-- 20. Простая инструкция DELETE.
-- DELETE FROM loan WHERE lstatus = 'closed'

-- 21. Инструкция DELETE с вложенным коррелированным подзапросом в предложении WHERE. 
-- DELETE FROM client_service
-- WHERE client_id IN (
--     SELECT client.id
--     FROM client JOIN account ON account.client_id = client.id
--     WHERE account.atype <> 'credit'
--     GROUP BY client.id
--     HAVING SUM(balance) > (SELECT fee FROM service WHERE service.id = client_service.service_id)
-- )
-- RETURNING client_service.id;

-- 22. Инструкция SELECT, использующая простое обобщенное табличное выражение 
-- Выбрать популярные, используемые и не используемые сервисы
-- WITH tmp (id, clientcnt) AS (
--     SELECT client_id, COUNT(*)
--     FROM client_service
--     GROUP BY client_id
-- )
-- SELECT id, clientcnt, (
--     CASE
--         WHEN clientcnt = 0 THEN 'unused'
--         WHEN clientcnt < 20 THEN 'used'
--         ELSE 'popular'
--     END
-- ) as metric
-- FROM tmp;


-- 23. Инструкция  SELECT, использующая рекурсивное обобщенное табличное выражение.
-- DROP TABLE IF EXISTS service_review;
-- CREATE TABLE IF NOT EXISTS service_review (
--     id SERIAL PRIMARY KEY,
--     service_id INTEGER REFERENCES service(id),
--     grade INTEGER CHECK (grade BETWEEN 1 AND 5),
--     review text,
--     on_review INTEGER DEFAULT NULL
-- );


-- INSERT INTO service_review (service_id, grade, review, on_review)
-- VALUES (5, 5, 'Отличный сервис!', NULL),
-- (5, 5, 'Согласен', 1),
-- (5, 1, 'Не согласен', 2);

-- WITH RECURSIVE direct_review (id, on_review, service_id, grade, review, on_level) AS (
--     SELECT sr.id, on_review, service_id, grade, review, 0 on_level
--     FROM service_review as sr
--     WHERE sr.on_review IS NULL
--     UNION ALL
--     SELECT sr.id, sr.on_review, sr.service_id, sr.grade, sr.review, on_level + 1
--     FROM service_review as sr JOIN direct_review as dr ON sr.on_review = dr.id
-- )
-- SELECT id, service_id, grade, review, on_review, on_level
-- FROM direct_review;

-- 24. Оконные функции. Использование конструкций MIN/MAX/AVG OVER()
-- Для каждого вида счёта добавить столбцы средних, минимальных и максимальных процентов.
-- SELECT id, atype, interest, 
--     AVG(interest) OVER(PARTITION BY atype) as avginterest, 
--     MAX(interest) OVER(PARTITION BY atype) as maxinterest, 
--     MIN(interest) OVER(PARTITION BY atype) as mininterest
-- FROM account

-- 25. Оконные фнкции для устранения дублей 
-- Придумать запрос, в результате которого в данных появляются полные дубли. 
-- Устранить дублирующиеся строки с использованием функции ROW_NUMBER().
-- INSERT INTO service (id, sname, sdescription, fee)
-- SELECT id + 1000, sname, sdescription, fee
-- FROM service;

-- SELECT * FROM service;
-- SELECT 
--     id,
--     sdescription,
--     sname,
--     fee
-- FROM (
--     SELECT 
--         ROW_NUMBER() OVER (PARTITION BY sdescription, sname, fee) AS cnt,
--         id,
--         sname,
--         sdescription,
--         fee 
--     FROM service
-- )
-- WHERE cnt = 1;