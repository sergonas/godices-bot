create table user_mapping (
    user_mapping_id serial not null,
    telegram_username varchar(255) not null,
    subsys_username varchar(255) not null,
    primary key(user_mapping_id)
);

create index user_mapping_telegram_idx on user_mapping(telegram_username);

create index user_mapping_subsys_idx on user_mapping(subsys_username);