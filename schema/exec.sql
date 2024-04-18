/*INSERT INTO t_order (orderTime, status, deadLine, startDate, endDate, productId, clientId, managerId, price) 
VALUES ('2024-04-16', 1, '2024-04-23', '2024-04-20', '2024-04-25', 5, 3, 5, 50000)*/

SELECT
    ord.id,
    userc.name as client_name,
    userc.username as client_username,
    userm.name as manager_name,
    userm.username as manager_username,
    ord.price
FROM t_order ord
INNER JOIN t_client client ON client.id = ord.clientId  
INNER JOIN t_manager manager ON manager.id = ord.managerId
INNER JOIN t_users userc ON userc.id = client.user_id
INNER JOIN t_users userm ON userm.id = manager.user_id