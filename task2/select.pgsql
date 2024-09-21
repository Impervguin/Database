--- 1. Инструкция SELECT, использующая предикат сравнения.
--- Найти все невыплаченные займы, оставшаяся сумма которых превышает 1 000 000, а ежемесячный платёж больше 10 000
SELECT id, remaining_amount, monthly_payment, lstatus
FROM loan
WHERE lstatus = 'active'
AND monthly_payment >= 10000
AND remaining_amount >= 1000000
ORDER BY remaining_amount;


--- 2. Инструкция SELECT, использующая предикат BETWEEN.
--- Найти все транзакция сделанные в мае 2022 года.
SELECT id, done_at
FROM transaction
WHERE done_at BETWEEN '2022-05-01' AND '2022-05-31'
ORDER BY done_at;

--- 3. Инструкция SELECT, использующая предикат LIKE.
--- Найти аккаунты всех пользователей, у которых почта оканчивается на .com
SELECT app_user.id AS user_id, client.email
FROM (app_user JOIN client ON app_user.client_id = client.id)
WHERE email LIKE '%.com';

--- 4. Инструкция SELECT, использующая предикат IN с вложенным подзапросом.
--- Найти имена всех клиентов банка у которых в июне 2022 были транзакции снятия или перевода
SELECT client.id, client.last_name, client.first_name
FROM client
WHERE client.id IN (
    SELECT account.client_id
    FROM account JOIN transaction ON transaction.account_id = account.id
    WHERE transaction.ttype = 'withdraw' OR (transaction.ttype = 'transfer' AND transaction.amount < 0) 
    AND transaction.done_at BETWEEN '2022-06-01' AND '2022-06-30'
);

--- 5. Инструкция SELECT, использующая предикат EXISTS с вложенным подзапросом.
--- Найти клиентов которые не подписаны ни на один сервис
SELECT client1.id, client1.last_name, client1.first_name
FROM client AS client1
WHERE NOT EXISTS (
    SELECT client_service.service_id
    FROM client_service
    WHERE client_service.client_id = client1.id
);

--- 6. Инструкция SELECT, использующая предикат сравнения с квантором.
--- Найти клиентов у которых ежемесечная плата по займу выше, чем остаток по всем счетам раздельно
SELECT client.id, client.last_name, client.first_name, loan.monthly_payment
FROM client JOIN loan ON client.id = loan.client_id
WHERE loan.monthly_payment > ALL (
    SELECT account.balance
    FROM account
    WHERE account.client_id = client.id
);

