create schema if not exists ls;

-- таблицы

create extension if not exists "uuid-ossp";

create table if not exists ls.song
(
    id            uuid primary key not null default uuid_generate_v4(),
    _group        text             not null,
    _song         text             not null,
    _release_date date             not null,
    _text         text             not null,
    _link         text             not null
);

-- роли бд

create role administrator;

grant usage on schema ls to administrator;
grant all privileges on all tables in schema ls to administrator;

create user admin_user with password 'admin';

grant administrator to admin_user;