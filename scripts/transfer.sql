create table transfers (
     id                         varchar(255)    not null
    ,account_origin_id          varchar(255)    not null
    ,account_destination_id     varchar(255)    not null
    ,amount                     bigint          not null
    ,created_at                 datetime        not null

    ,primary key(id)
    ,foreign key (account_origin_id) references accounts (id)
    ,foreign key (account_destination_id) references accounts (id)
)