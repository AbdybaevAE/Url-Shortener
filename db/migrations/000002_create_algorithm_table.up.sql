create table if not exists algos (
    algo_id serial not null primary key,
    algo_strategy varchar(100) not null unique,
    number_id integer references numbers(number_id),
    increment_value integer not null default 500,
    dict varchar(100) not null
);