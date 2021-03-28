create table if not exists algorithms (
    algorithm_id serial not null primary key,
    algorithm_name varchar(100) not null unique,
    algorithm_number_id integer references numbers(number_id),
    algorithm_metadata varchar(1000) not null
);