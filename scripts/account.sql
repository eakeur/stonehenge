create table accounts (
     id             varchar(255)    not null
    ,name           varchar(255)    not null
    ,cpf            varchar(14)     not null
    ,balance        bigint          not null    default 0
    ,secret         varchar(255)    not null
    ,created_at     datetime        not null

    ,primary key (id)
    ,unique(cpf)
)