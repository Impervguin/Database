ALTER TABLE client
    ADD CONSTRAINT client_pk PRIMARY KEY (id),
    ADD CONSTRAINT phone_uniq UNIQUE (phone_number),
    ADD CONSTRAINT email_uniq UNIQUE (email),
    ADD CONSTRAINT valid_birthday_created CHECK (dob  + interval '18 years' <= created_at),
    ADD CONSTRAINT valid_email CHECK (email ~ '^[\w\-\.]+@([\w\-]+\.)+[\w\-]+$'),
    ADD CONSTRAINT valid_phone CHECK (phone_number ~ '^[\+]?[0-9]?[-\s]?[0-9]{3}[-\s]?[0-9]{3}[-\s]?[0-9]{4,6}$'),
    ALTER COLUMN phone_number set NOT NULL,
    ALTER COLUMN dob set NOT NULL,
    ALTER COLUMN created_at set NOT NULL,
    ALTER COLUMN first_name set NOT NULL,
    ALTER COLUMN last_name set NOT NULL,
    ALTER COLUMN address set NOT NULL;

ALTER TABLE account
    ADD CONSTRAINT account_pk PRIMARY KEY (id),
    ADD CONSTRAINT fk_client_id FOREIGN KEY (client_id) REFERENCES client (id),
    ADD CONSTRAINT balance_positive CHECK (balance >= 0),
    ADD CONSTRAINT interest_positive CHECK (interest >= 0),
    -- ADD CONSTRAINT created_after_client CHECK (created_at >= ) -- Условие на то, чтобы счёт был создан после клиента
    ALTER COLUMN client_id set NOT NULL,
    ALTER COLUMN balance set NOT NULL,
    ALTER COLUMN created_at set NOT NULL,
    ALTER COLUMN astatus set NOT NULL,
    ALTER COLUMN atype set NOT NULL;

ALTER TABLE card
    ADD CONSTRAINT card_pk PRIMARY KEY (id),
    ADD CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES account (id),
    ADD CONSTRAINT card_number_uniq UNIQUE (cnumber),
    ADD CONSTRAINT cvv_valid CHECK (cvv ~ '^[0-9]{3}$'),
    ADD CONSTRAINT card_number_valid CHECK (cnumber ~ '^[0-9]{16}$'),
    ADD CONSTRAINT expire_after_create CHECK (created_at < expired_at),
    -- ADD CONSTRAINT created_after_account CHECK (created_at >= ) -- Условие на то, чтобы карта была создана после счёта
    ALTER COLUMN account_id set NOT NULL,
    ALTER COLUMN cnumber set NOT NULL,
    ALTER COLUMN cvv set NOT NULL,
    ALTER COLUMN expired_at set NOT NULL,
    ALTER COLUMN created_at set NOT NULL,
    ALTER COLUMN cstatus set NOT NULL;


ALTER TABLE transaction
    ADD CONSTRAINT transaction_pk PRIMARY KEY (id),
    ADD CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES account (id),
    -- ADD CONSTRAINT created_after_account CHECK (created_at >= ) -- Условие на то, чтобы транзакция была создана после счёта
    ADD CONSTRAINT balance_after_positive CHECK (balance_after >= 0),
    ALTER COLUMN account_id set NOT NULL,
    ALTER COLUMN ttype set NOT NULL,
    ALTER COLUMN amount set NOT NULL,
    ALTER COLUMN balance_after set NOT NULL,
    ALTER COLUMN done_at set NOT NULL;

ALTER TABLE loan
    ADD CONSTRAINT loan_pk PRIMARY KEY (id),
    ADD CONSTRAINT fk_client_id FOREIGN KEY (client_id) REFERENCES client (id),
    -- ADD CONSTRAINT started_after_client CHECK (created_at >= ) -- Условие на то, чтобы заём был создан после клиента
    ADD CONSTRAINT amount_positive CHECK (amount >= 0),
    ADD CONSTRAINT interest_positive CHECK (interest >= 0),
    ADD CONSTRAINT remaining_positive CHECK (remaining_amount >= 0),
    ADD CONSTRAINT monthly_payment_positive CHECK (monthly_payment >= 0),
    ADD CONSTRAINT start_end_valid CHECK (start_date <= end_date),
    ALTER COLUMN client_id set NOT NULL,
    ALTER COLUMN amount set NOT NULL,
    ALTER COLUMN remaining_amount set NOT NULL,
    ALTER COLUMN monthly_payment set NOT NULL,
    ALTER COLUMN interest set NOT NULL,
    ALTER COLUMN start_date set NOT NULL,
    ALTER COLUMN end_date set NOT NULL,
    ALTER COLUMN lstatus set NOT NULL;

ALTER TABLE service
    ADD CONSTRAINT service_pk PRIMARY KEY (id),
    ADD CONSTRAINT fee_positive CHECK (fee >= 0),
    ALTER COLUMN sname set NOT NULL,
    ALTER COLUMN fee set NOT NULL;

ALTER TABLE client_service
    ADD CONSTRAINT client_service_pk PRIMARY KEY (id),
    ADD CONSTRAINT fk_client_id FOREIGN KEY (client_id) REFERENCES client (id),
    ADD CONSTRAINT fk_service_id FOREIGN KEY (service_id) REFERENCES service (id);

ALTER TABLE app_user
    ADD CONSTRAINT app_user_pk PRIMARY KEY (id),
    ADD CONSTRAINT fk_client_id FOREIGN KEY (client_id) REFERENCES client (id),
    ADD CONSTRAINT username_uniq UNIQUE (username),
    ADD CONSTRAINT failed_attempts_positive CHECK (failed_attempts >= 0),
    ALTER COLUMN username set NOT NULL,
    ALTER COLUMN client_id set NOT NULL,
    ALTER COLUMN hashpassword set NOT NULL,
    ALTER COLUMN failed_attempts set NOT NULL,
    ALTER COLUMN ustatus set NOT NULL;

ALTER TABLE notification
    ADD CONSTRAINT notification_pk PRIMARY KEY (id),
    ADD CONSTRAINT fk_client_id FOREIGN KEY (client_id) REFERENCES client (id),
    ALTER COLUMN client_id set NOT NULL,
    ALTER COLUMN nmessage set NOT NULL,
    ALTER COLUMN created_at set NOT NULL,
    ALTER COLUMN seen set NOT NULL,
    ALTER COLUMN seen set DEFAULT FALSE;
    



    



