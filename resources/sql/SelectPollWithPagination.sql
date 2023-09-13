select p.id,
       title,
       description,
       startsAt,
       endsAt,
       state,
       p.created_at as createdAt,
       p.updated_at as updatedAt,
       primaryOrganizerFullName,
       primaryOrganizerEmail
from (select poll.id                                as id,
             poll.title                             as title,
             poll.description                       as description,
             poll.starts_at                         as startsAt,
             poll.ends_at                           as endsAt,
             poll.state                             as state,
             poll.created_at                        as created_at,
             poll.updated_at                        as updated_at,
             concat(u.first_name, ' ', u.last_name) as primaryOrganizerFullName,
             u.email                                as primaryOrganizerEmail
      from poll poll
               left join poll_organizer po2 on poll.id = po2.fk_poll_id
               left join user u on u.id = po2.fk_organizer_id
      where primary_organizer = true) p
         left join poll_organizer po on po.fk_poll_id = p.id
where %v
ORDER BY %v %v limit %v OFFSET %v;