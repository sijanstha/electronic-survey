insert into user(first_name, last_name, email, hash_password, created_at, updated_at)
values ('Test', 'User', 'test@evs.com', '$2a$10$28cm0tPse3.cFS1TmV..NOtpI1FeFyUXqwbm1NxukmlFvNhRHE/Qi', now(), now()),
       ('Bob', 'Marley', 'bob@evs.com', '$2a$10$28cm0tPse3.cFS1TmV..NOtpI1FeFyUXqwbm1NxukmlFvNhRHE/Qi', now(), now()),
       ('John', 'Doe', 'john.doe@evs.com', '$2a$10$28cm0tPse3.cFS1TmV..NOtpI1FeFyUXqwbm1NxukmlFvNhRHE/Qi', now(),
        now());

insert into poll(title, description, starts_at, ends_at, state, timezone, created_at, updated_at)
values ('Poll title #1',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 1 day_hour), datetime),
        convert(date_add(now(), interval 1 day), datetime),
        'PREPARED',
        'Asia/Katmandu',
        now(),
        now()),
       ('Poll title #2',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 1 day_hour), datetime),
        convert(date_add(now(), interval 2 day), datetime),
        'PREPARED',
        'Asia/Katmandu',
        now(),
        now()),
       ('Poll title #3',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 2 day_hour), datetime),
        convert(date_add(now(), interval 1 day), datetime),
        'STARTED',
        'Asia/Katmandu',
        now(),
        now()),
       ('Poll title #4',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 30 day_minute), datetime),
        convert(date_add(now(), interval 1 day), datetime),
        'FINISHED',
        'Asia/Katmandu',
        now(),
        now()),
       ('Poll title #5',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 1 day_hour), datetime),
        convert(date_add(now(), interval 1 day), datetime),
        'PREPARED',
        'Asia/Katmandu',
        now(),
        now()),
       ('Poll title #6',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 1 day_hour), datetime),
        convert(date_add(now(), interval 1 day), datetime),
        'PREPARED',
        'Asia/Katmandu',
        now(),
        now()),
       ('Poll title #7',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 1 day_hour), datetime),
        convert(date_add(now(), interval 1 day), datetime),
        'PREPARED',
        'Asia/Katmandu',
        now(),
        now()),
       ('Poll title #8',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 1 day_hour), datetime),
        convert(date_add(now(), interval 1 day), datetime),
        'PREPARED',
        'Asia/Katmandu',
        now(),
        now()),
       ('Poll title #9',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 1 day_hour), datetime),
        convert(date_add(now(), interval 1 day), datetime),
        'PREPARED',
        'Asia/Katmandu',
        now(),
        now()),
       ('Poll title #10',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into el',
        convert(date_add(now(), interval 1 day_hour), datetime),
        convert(date_add(now(), interval 1 day), datetime),
        'PREPARED',
        'Asia/Katmandu',
        now(),
        now());

insert into poll_organizer(fk_poll_id, fk_organizer_id, primary_organizer, created_at, updated_at)
values ((select id from poll where title = 'Poll title #1'),
        (select id from user where email = 'test@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #2'),
        (select id from user where email = 'test@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #3'),
        (select id from user where email = 'test@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #4'),
        (select id from user where email = 'test@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #5'),
        (select id from user where email = 'test@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #1'),
        (select id from user where email = 'bob@evs.com'),
        false,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #6'),
        (select id from user where email = 'bob@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #7'),
        (select id from user where email = 'bob@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #8'),
        (select id from user where email = 'bob@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #9'),
        (select id from user where email = 'bob@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #10'),
        (select id from user where email = 'bob@evs.com'),
        true,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #8'),
        (select id from user where email = 'test@evs.com'),
        false,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #9'),
        (select id from user where email = 'test@evs.com'),
        false,
        now(),
        now()),
       ((select id from poll where title = 'Poll title #10'),
        (select id from user where email = 'test@evs.com'),
        false,
        now(),
        now())
;