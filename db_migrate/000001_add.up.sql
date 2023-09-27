CREATE TABLE users(
id serial not null unique,
nname  varchar(250) not null,
username varchar(250) not null unique,
password_hash varchar(250) not null
);

CREATE TABLE lists(
   id serial not null unique,
   title varchar(250) not null,
   ddescription varchar(250)

);

CREATE TABLE users_lists
(
 id serial not null unique,
 id_users int references users(id) on delete cascade not null,
 id_lists int references lists(id) on delete cascade not null

);

CREATE TABLE todo_items
(
id serial not null unique,
title varchar(250) not null,
ddescription varchar(250),
done boolean 
);

CREATE TABLE lists_items
(
    id serial not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    lists_id int references lists (id) on delete cascade not null

);