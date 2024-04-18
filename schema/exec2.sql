select 
    u.id, 
    u.name, 
    u.surname, 
    u.username, 
    u.email, 
    m.id as roleid 
from t_users u 
join t_manager m ON m.user_id = u.id 
order by m.created_time DESC
limit 10
OFFSET 3