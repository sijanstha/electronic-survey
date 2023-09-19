create table if not exists participant_list
(
    id            int          not null auto_increment primary key,
    name          varchar(100) not null,
    emails        json not null,
    fk_organizer_id   int       not null,
    created_at    timestamp    not null,
    updated_at    timestamp    not null
);

alter table participant_list
    add foreign key (fk_organizer_id) references user (id);