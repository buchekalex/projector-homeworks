explain select * from users where date_of_birth = '2009-05-31';

CREATE INDEX idx_hash_date_of_birth ON users (date_of_birth) USING HASH;

analyze select * from users where date_of_birth = '2009-05-31';

analyze SELECT * FROM users USE INDEX (idx_hash_date_of_birth) WHERE date_of_birth = '2000-01-01';

SET GLOBAL innodb_flush_log_at_trx_commit = 0;

explain insert into users (id, name, date_of_birth) values (default, 'somename', '2000-01-01');

show variables like 'innodb_flush_log_at_trx_commit';


## Test insert queries

begin;
insert into users (id, name, date_of_birth) values (default, 'somename', '2000-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2002-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2003-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2004-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2013-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2220-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2220-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2550-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2660-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2770-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2880-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2990-01-01');
insert into users (id, name, date_of_birth) values (default, 'somename', '2120-01-01');
COMMIT;