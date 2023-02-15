drop table if exists users cascade;
CREATE TABLE users ( 
    username varchar(30) primary key ,
    email varchar(255) not null,
    password varchar(30) not null
);

drop table if exists tokens cascade;
CREATE TABLE tokens (
 username varchar(30) primary key references users,
 token varchar(300) not null
);
drop table if exists domains cascade;
create table domains (
  username varchar(30) references users,
  domain varchar(200) not null unique
);

drop table if exists views cascade;
create table views (
domain varchar(200) references domains(domain),
view_count int ,
view_date varchar(20)
);

