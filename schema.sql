create database if not exists testms;

use testms;

create table if not exists ports
(
	name varchar(5) null,
	city varchar(100) not null,
	country varchar(100) not null,
	lat float not null,
	lng float not null
);

create unique index ports_name_uindex
	on ports (name);

alter table ports
	add constraint ports_pk
		primary key (name);
