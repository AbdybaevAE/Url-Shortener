create table if not exists links(
    link_id serial not null primary key,
    link_value varchar(500) not null,
    key_value varchar(10) not null,
    created_at timestamp not null default now(),
    visited_at timestamp,
    expired_at timestamp not null
);