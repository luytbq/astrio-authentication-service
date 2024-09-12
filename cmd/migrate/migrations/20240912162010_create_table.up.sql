create table users (
    id serial primary key,
	email varchar(256) unique,
	password varchar(256),
	password_salt varchar(256),
	create_at timestamp
);

alter sequence users_id_seq restart with 123123;

create table login_sessions (
	id serial primary key,
	user_id integer references users(id),
	ip_address varchar(32),
	login_at timestamp
);
