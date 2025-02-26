drop table if exists transactions;
drop table if exists bank_accounts;
drop table if exists broker_accounts;
drop table if exists banks;
drop table if exists brokers;
drop table if exists deposits;
drop table if exists coupons;
drop table if exists bonds;
drop type if exists currency;

create type currency as enum ('UAH', 'USD', 'EUR');

create table transactions
(
    id            text primary key,
    amount        int      not null,
    currency_type currency not null,
    created_at    timestamp without time zone default now()
);

create table brokers
(
    id         text primary key,
    name       text not null unique,
    created_at timestamp without time zone default now()
);

create table broker_accounts
(
    id         text primary key,
    broker_id  text references brokers (id) not null,
    created_at timestamp without time zone default now()
);

create table banks
(
    id         text primary key,
    name       text not null unique,
    created_at timestamp without time zone default now()
);

create table bank_accounts
(
    id            text primary key,
    name          text                       not null,
    amount        int                        not null,
    currency_type currency                   not null,
    bank_id       text references banks (id) not null,
    created_at    timestamp without time zone default now()
);

create table deposits
(
    id            text primary key,
    name          text,
    amount        int      not null,
    currency_type currency not null,
    start_date    timestamp without time zone,
    end_date      timestamp without time zone,
    created_at    timestamp without time zone default now()
);

create table bonds
(
    id            text primary key,
    name          text,
    isin varchar(12) not null,
    count         int            not null,
    buy_price     numeric(10, 2) not null,
    sell_price    numeric(10, 2) not null,
    currency_type currency       not null,
    start_date    date,
    end_date      date           not null,
    created_at    timestamp without time zone default now()
);

create table coupons
(
    id           text primary key,
    amount       numeric(10, 2)             not null,
    payment_date date                       not null,
    bond         text references bonds (id) not null,
    created_at   timestamp without time zone default now()
);
