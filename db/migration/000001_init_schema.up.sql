create table if not exists user
(
    id            int          not null auto_increment primary key,
    first_name    varchar(50)  not null,
    last_name     varchar(50)  not null,
    email         varchar(100) not null,
    role          varchar(50)  not null default "ROLE_ORGANIZER",
    hash_password varchar(200) not null,
    created_at    timestamp    not null,
    updated_at    timestamp    not null
);

create table if not exists poll
(
    id          int          not null auto_increment primary key,
    title       varchar(100) not null unique,
    description longtext,
    state       varchar(50)  not null default "PREPARED",
    starts_at   datetime     not null,
    ends_at     datetime     not null,
    timezone    varchar(100) not null,
    created_at  timestamp    not null,
    updated_at  timestamp    not null
);

create table if not exists poll_organizer
(
    id                int       not null auto_increment primary key,
    fk_poll_id        int       not null,
    fk_organizer_id   int       not null,
    primary_organizer bit(1) default 0,
    created_at        timestamp not null,
    updated_at        timestamp not null
);


alter table poll_organizer
    add foreign key (fk_poll_id) references poll (id);
alter table poll_organizer
    add foreign key (fk_organizer_id) references user (id);
create unique index id_user_email on user (email);
create index idx_poll_state on poll (state);
create index idx_poll_start_end_date on poll (starts_at, ends_at);