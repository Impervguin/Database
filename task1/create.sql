CREATE TABLE IF NOT EXISTS client (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(50),
        last_name VARCHAR(50),
        dob DATE,
        email VARCHAR(100),
        phone_number VARCHAR(20),
        address VARCHAR(255),
        created_at DATE
);

CREATE TYPE account_type AS ENUM('savings', 'checking', 'fd', 'credit');
CREATE TYPE account_status AS ENUM('active', 'inactive', 'closed');

CREATE TABLE IF NOT EXISTS account (
        id SERIAL PRIMARY KEY,
        client_id INTEGER,
        balance NUMERIC(20, 2),
        interest NUMERIC(6, 2),
        created_at DATE,
        atype account_type,
        astatus account_status
);

CREATE TYPE card_status AS ENUM('active', 'blocked', 'expired');

CREATE TABLE IF NOT EXISTS card (
        id SERIAL PRIMARY KEY,
        account_id INTEGER,
        cnumber VARCHAR(16),
        cvv VARCHAR(3),
        created_at DATE,
        expired_at DATE,
        cstatus card_status
);

CREATE TYPE transaction_type AS ENUM('deposit', 'withdraw', 'transfer');

CREATE TABLE IF NOT EXISTS transaction (
        id SERIAL PRIMARY KEY,
        account_id INTEGER,
        ttype transaction_type,
        amount NUMERIC(20, 2),
        done_at TIMESTAMP,
        balance_after NUMERIC(20, 2),
        system_description text,
        client_description text
);

CREATE TYPE loan_status AS ENUM('active', 'closed', 'defaulted');

CREATE TABLE IF NOT EXISTS loan (
    
        id SERIAL PRIMARY KEY,
        client_id INTEGER,
        amount NUMERIC(20, 2),
        interest NUMERIC(6, 2),
        remaining_amount NUMERIC(20, 2),
        monthly_payment NUMERIC(20, 2),
        start_date DATE,
        end_date DATE,
        lstatus loan_status,
        ldescription text
);

CREATE TABLE IF NOT EXISTS service (
    
        id SERIAL PRIMARY KEY,
        sname VARCHAR(50),
        sdescription text,
        fee NUMERIC(6, 2)
    
);

CREATE TABLE IF NOT EXISTS client_service (
    
        id SERIAL PRIMARY KEY,
        client_id INTEGER,
        service_id INTEGER
    
);

CREATE TYPE user_status AS ENUM('active', 'blocked');


CREATE TABLE IF NOT EXISTS app_user (
        id SERIAL PRIMARY KEY,
        client_id INTEGER,
        username VARCHAR(50),
        hashpassword VARCHAR(255),
        last_login TIMESTAMP,
        failed_attempts INTEGER,
        ustatus user_status 
);

CREATE TABLE IF NOT EXISTS notification (
        id SERIAL PRIMARY KEY,
        client_id INTEGER,
        nmessage text,
        created_at TIMESTAMP,
        seen BOOLEAN
);