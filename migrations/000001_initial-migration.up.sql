create table account (
	id Serial primary key,
	document varchar(20) not null,
	created_at timestamp default current_timestamp,
	deleted_at timestamp default null
);

create table operation_type (
	id Serial primary key,
	description varchar(50) not null unique,
	created_at timestamp default current_timestamp,
	deleted_at timestamp default null
);

create table transaction (
	id Serial primary key,
	account_id int not null,
	operation_type_id int not null,
	amount decimal(10,5),
	created_at timestamp default current_timestamp,
	deleted_at timestamp default null,
	constraint fk_account foreign key(account_id) references account(id) on delete set null,
	constraint fk_operation_type foreign key(operation_type_id) references operation_type(id) on delete set null
);
