# Project Setup

## run below command to install go mod

`go mod init github.com/product-management`

## install gorilla/mux router

`go get github.com/gorilla/mux`

## install a mysql package for go

`go get github.com/go-sql-driver/mysql v1.6.0`

## Setup database

1. Install mysql
2. Run below scripts

```sql
create database inventory;

use inventory;

create table products(id int NOT NULL AUTO_INCREMENT, name varchar(255) not null, quantity int, price float(10,7), primary key(id));

insert into products values(1, "chair", 100, 200.00);
```
