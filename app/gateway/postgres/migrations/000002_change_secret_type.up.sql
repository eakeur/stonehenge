begin;

alter table accounts alter column secret type varchar (60);

commit;