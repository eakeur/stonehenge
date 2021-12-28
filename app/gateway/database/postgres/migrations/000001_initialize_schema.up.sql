begin;

-- create extension if not exists "uuid-ossp";

create table if not exists accounts (
     --id                         uuid            not null    default uuid_generate_v4()
     id                         varchar(36)     not null
    ,name                       varchar(255)    not null
    ,document                   varchar(14)     not null
    ,balance                    bigint          not null    default 0
    ,secret                     varchar(255)    not null
    ,updated_at                 timestamp       not null    default now()
    ,created_at                 timestamp       not null    default now()

    ,primary key (id)
    ,unique      (document)
);


create table if not exists transfers (
     id                         varchar(36)     not null
    ,account_origin_id          varchar(36)     not null
    ,account_destination_id     varchar(36)     not null
    ,amount                     bigint          not null
    ,effective_date             timestamp       not null    default now()
    ,updated_at                 timestamp       not null    default now()
    ,created_at                 timestamp       not null    default now()

    ,primary key (id)
    ,foreign key (account_origin_id)      references accounts (id)
    ,foreign key (account_destination_id) references accounts (id)
);


commit;