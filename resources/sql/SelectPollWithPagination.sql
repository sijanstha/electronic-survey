select p.id                                   as id,
       p.title                                as title,
       p.description                          as description,
       p.starts_at                            as startsAt,
       p.ends_at                              as endsAt,
       p.state                                as state,
       p.created_at                           as createdAt,
       p.updated_at                           as updatedAt,
       concat(u.first_name, ' ', u.last_name) as primaryOrganizerFullName,
       u.email                                as primaryOrganizerEmail
from poll p
         left join poll_organizer po on p.id = po.fk_poll_id
         left join user u on u.id = po.fk_organizer_id
where po.primary_organizer = true and %v
ORDER BY %v %v limit %v OFFSET %v;