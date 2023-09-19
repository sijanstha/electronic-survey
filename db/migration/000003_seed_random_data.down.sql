delete
from poll_organizer
where fk_poll_id in (select id
                     from poll
                     where title in ('Poll title #1',
                                     'Poll title #2',
                                     'Poll title #3',
                                     'Poll title #4',
                                     'Poll title #5',
                                     'Poll title #6',
                                     'Poll title #7',
                                     'Poll title #8',
                                     'Poll title #9',
                                     'Poll title #10'));

delete
from poll
where title in ('Poll title #1',
                'Poll title #2',
                'Poll title #3',
                'Poll title #4',
                'Poll title #5',
                'Poll title #6',
                'Poll title #7',
                'Poll title #8',
                'Poll title #9',
                'Poll title #10');

delete
from user
where email in ('test@evs.com', 'bob@evs.com', 'john.doe@evs.com');