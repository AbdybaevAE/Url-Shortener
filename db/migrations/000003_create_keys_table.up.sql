create table if not exists keys(
    key_id serial not null primary key,
    key_value varchar(10),
    algo_id integer references algos(algo_id)
);