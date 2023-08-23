[2023-08-23 18:48:00] completed in 2 ms
your_db> select * from users where date_of_birth = '2009-05-31'
[2023-08-23 18:48:01] 500 rows retrieved starting from 1 in 794 ms (execution: 688 ms, fetching: 106 ms)
your_db> explain select * from users where date_of_birth = '2009-05-31'
[2023-08-23 18:48:09] 1 row retrieved starting from 1 in 32 ms (execution: 5 ms, fetching: 27 ms)
your_db> explain analyze select * from users where date_of_birth = '2009-05-31'
[2023-08-23 18:49:51] [42000][1064] (conn=4) You have an error in your SQL syntax; check the manual that corresponds to your MariaDB server version for the right syntax to use near 'analyze select * from users where date_of_birth = '2009-05-31'' at line 1
your_db> explain select * from users where date_of_birth = '2009-05-31'
[2023-08-23 18:50:55] 1 row retrieved starting from 1 in 27 ms (execution: 7 ms, fetching: 20 ms)
your_db> analyze select * from users where date_of_birth = '2009-05-31'
[2023-08-23 18:51:27] 1 row retrieved starting from 1 in 5 s 610 ms (execution: 5 s 583 ms, fetching: 27 ms)
your_db> analyze select * from users where date_of_birth = '2009-05-31'
[2023-08-23 18:51:31] 1 row retrieved starting from 1 in 4 s 624 ms (execution: 4 s 605 ms, fetching: 19 ms)
your_db> explain select * from users where date_of_birth = '2009-05-31'
[2023-08-23 18:56:40] 1 row retrieved starting from 1 in 29 ms (execution: 8 ms, fetching: 21 ms)
your_db> CREATE INDEX idx_date_of_birth ON users (date_of_birth) USING BTREE
[2023-08-23 18:57:31] completed in 33 s 25 ms
your_db> select * from users where date_of_birth = '2009-05-31'
[2023-08-23 19:04:47] 500 rows retrieved starting from 1 in 67 ms (execution: 40 ms, fetching: 27 ms)
your_db> explain select * from users where date_of_birth = '2009-05-31'
[2023-08-23 19:04:50] 1 row retrieved starting from 1 in 22 ms (execution: 4 ms, fetching: 18 ms)
your_db> analyze select * from users where date_of_birth = '2009-05-31'
[2023-08-23 19:05:14] 1 row retrieved starting from 1 in 151 ms (execution: 129 ms, fetching: 22 ms)
your_db> CREATE INDEX idx_date_of_birth ON users (date_of_birth) USING HASH
[2023-08-23 19:12:02] [42000][1061] (conn=4) Duplicate key name 'idx_date_of_birth'
your_db> CREATE INDEX idx_hash_date_of_birth ON users (date_of_birth) USING HASH
[2023-08-23 19:12:44] completed in 33 s 28 ms
your_db> SELECT * FROM users USE INDEX (idx_hash_date_of_birth) WHERE date_of_birth = '2000-01-01'
[2023-08-23 19:13:56] 500 rows retrieved starting from 1 in 43 ms (execution: 22 ms, fetching: 21 ms)
your_db> explain SELECT * FROM users USE INDEX (idx_hash_date_of_birth) WHERE date_of_birth = '2000-01-01'
[2023-08-23 19:14:02] 1 row retrieved starting from 1 in 21 ms (execution: 4 ms, fetching: 17 ms)
your_db> analyze SELECT * FROM users USE INDEX (idx_hash_date_of_birth) WHERE date_of_birth = '2000-01-01'
[2023-08-23 19:14:37] 1 row retrieved starting from 1 in 136 ms (execution: 125 ms, fetching: 11 ms)
your_db> SET GLOBAL innodb_flush_log_at_trx_commit = 0
[2023-08-23 19:17:52] [42000][1227] (conn=4) Access denied; you need (at least one of) the SUPER privilege(s) for this operation
your_db> explain insert into users (id, name, date_of_birth) values (default, 'somename', '2000-01-01')
[2023-08-23 19:24:03] 1 row retrieved starting from 1 in 27 ms (execution: 10 ms, fetching: 17 ms)


## innodb_flush_log_at_trx_commit=0

your_db> begin
[2023-08-23 19:53:45] completed in 4 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2000-01-01')
[2023-08-23 19:53:45] 1 row affected in 6 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2002-01-01')
[2023-08-23 19:53:45] 1 row affected in 4 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2003-01-01')
[2023-08-23 19:53:45] 1 row affected in 5 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2004-01-01')
[2023-08-23 19:53:45] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2013-01-01')
[2023-08-23 19:53:45] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2220-01-01')
[2023-08-23 19:53:45] 1 row affected in 4 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2220-01-01')
[2023-08-23 19:53:45] 1 row affected in 4 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2550-01-01')
[2023-08-23 19:53:45] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2660-01-01')
[2023-08-23 19:53:45] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2770-01-01')
[2023-08-23 19:53:45] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2880-01-01')
[2023-08-23 19:53:45] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2990-01-01')
[2023-08-23 19:53:45] 1 row affected in 5 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2120-01-01')
[2023-08-23 19:53:45] 1 row affected in 2 ms
your_db> COMMIT
[2023-08-23 19:53:45] completed in 2 ms


## innodb_flush_log_at_trx_commit=1
your_db> begin
[2023-08-23 19:54:13] completed in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2000-01-01')
[2023-08-23 19:54:13] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2002-01-01')
[2023-08-23 19:54:13] 1 row affected in 9 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2003-01-01')
[2023-08-23 19:54:13] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2004-01-01')
[2023-08-23 19:54:13] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2013-01-01')
[2023-08-23 19:54:13] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2220-01-01')
[2023-08-23 19:54:13] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2220-01-01')
[2023-08-23 19:54:13] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2550-01-01')
[2023-08-23 19:54:13] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2660-01-01')
[2023-08-23 19:54:13] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2770-01-01')
[2023-08-23 19:54:13] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2880-01-01')
[2023-08-23 19:54:13] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2990-01-01')
[2023-08-23 19:54:13] 1 row affected in 5 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2120-01-01')
[2023-08-23 19:54:13] 1 row affected in 2 ms
your_db> COMMIT
[2023-08-23 19:54:13] completed in 3 ms

## innodb_flush_log_at_trx_commit=2
your_db> begin
[2023-08-23 19:54:49] completed in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2000-01-01')
[2023-08-23 19:54:49] 1 row affected in 4 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2002-01-01')
[2023-08-23 19:54:49] 1 row affected in 4 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2003-01-01')
[2023-08-23 19:54:49] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2004-01-01')
[2023-08-23 19:54:49] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2013-01-01')
[2023-08-23 19:54:49] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2220-01-01')
[2023-08-23 19:54:49] 1 row affected in 2 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2220-01-01')
[2023-08-23 19:54:49] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2550-01-01')
[2023-08-23 19:54:49] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2660-01-01')
[2023-08-23 19:54:49] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2770-01-01')
[2023-08-23 19:54:49] 1 row affected in 9 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2880-01-01')
[2023-08-23 19:54:49] 1 row affected in 3 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2990-01-01')
[2023-08-23 19:54:49] 1 row affected in 4 ms
your_db> insert into users (id, name, date_of_birth) values (default, 'somename', '2120-01-01')
[2023-08-23 19:54:49] 1 row affected in 3 ms
your_db> COMMIT
[2023-08-23 19:54:49] completed in 3 ms
