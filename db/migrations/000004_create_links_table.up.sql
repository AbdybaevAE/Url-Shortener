create table if not exists links(
    link_id serial not null primary key,
    link_value varchar(500) not null,
    key_value varchar(10) UNIQUE not null,
    created_at timestamp not null default now(),
    visited_at timestamp not null default now(),
    expired_at timestamp not null,
    visited boolean not null default false
);