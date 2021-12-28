begin;

-- create extension if not exists "uuid-ossp";

create table if not exists accounts (
     id                         serial          not null
    ,external_id                uuid            not null    default uuid_generate_v4()
    ,name                       varchar(255)    not null
    ,document                   varchar(14)     not null
    ,balance                    bigint          not null    default 0
    ,secret                     varchar(255)    not null
    ,updated_at                 timestamp       not null    default now()
    ,created_at                 timestamp       not null    default now()

    ,primary key (id)
    ,unique      (document)
    ,unique      (external_id)
);


create table if not exists transfers (
     id                         serial          not null
    ,external_id                uuid            not null    default uuid_generate_v4()
    ,account_origin_id          int             not null
    ,account_destination_id     int             not null
    ,amount                     bigint          not null
    ,effective_date             timestamp       not null    default now()
    ,updated_at                 timestamp       not null    default now()
    ,created_at                 timestamp       not null    default now()

    ,primary key (id)
    ,foreign key (account_origin_id)      references accounts (id)
    ,foreign key (account_destination_id) references accounts (id)
);


commit;