select json_object(
               'id', res.id,
               'title', res.title,
               'description', res.description,
               'state', res.state,
               'timezone', res.timezone,
               'startsAt', date_format(res.starts_at, '%%Y-%%m-%%dT%%H:%%i:%%S.000Z'),
               'endsAt', date_format(res.ends_at, '%%Y-%%m-%%dT%%H:%%i:%%S.000Z'),
               'createdAt', date_format(res.created_at, '%%Y-%%m-%%dT%%H:%%i:%%S.000Z'),
               'updatedAt', date_format(res.updated_at, '%%Y-%%m-%%dT%%H:%%i:%%S.000Z'),
               'pollOrganizers', json_arrayagg(
                       json_object(
                               'email', res.email,
                               'fullName', res.name,
                               'id', res.organizerId,
                               'primaryOrganizer', res.primaryOrganizer
                           )
                   )
           )
from (select p.id,
             p.title,
             description,
             state,
             timezone,
             starts_at,
             ends_at,
             p.created_at,
             p.updated_at,
             u.email,
             concat(u.first_name, ' ', u.last_name) as name,
             u.id                                   as organizerId,
             if(po.primary_organizer, true, false)  as primaryOrganizer
      from poll p
               left join poll_organizer po on p.id = po.fk_poll_id
               left join user u on po.fk_organizer_id = u.id
      where %s) as res