create database if not exists testms;

use testms;

create table ports
(
  id      varchar(5)   not null,
  name    varchar(30)  not null,
  city    varchar(100) not null,
  country varchar(100) not null,
  lat     float        not null,
  lng     float        not null,
  constraint ports_id_uindex
    unique (id)
);
