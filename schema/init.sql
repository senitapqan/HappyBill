CREATE TABLE IF NOT EXISTS t_users (
    id serial Primary key,
    username varchar(50) unique not null,
    password varchar(50) not null, 
    email varchar(255) not null, 
    name varchar(50) not null, 
    surname varchar(50) not null
);

CREATE TABLE IF NOT EXISTS t_roles (
   id serial Primary key, 
   role_name varchar(50) not null
);

CREATE TABLE IF NOT EXISTS t_users_roles (
    user_id int not null,
    role_id int not null, 
    Primary key(user_id, role_id), 
    FOREIGN key (user_id) REFERENCES t_users(id),
    FOREIGN key (role_id) REFERENCES t_roles(id)
);

CREATE TABLE IF NOT EXISTS t_client (
   id serial Primary key, 
   user_id int not null,
   FOREIGN KEY (user_id) 
      REFERENCES t_users(id)
);

CREATE TABLE IF NOT EXISTS t_manager (
    id serial Primary key, 
    user_id int not null,
    FOREIGN KEY (user_id) 
      REFERENCES t_users(id)
);

CREATE TABLE IF NOT EXISTS t_admin (
    id serial Primary key, 
    user_id int not null,
    FOREIGN KEY (user_id) 
      REFERENCES t_users(id)
);

CREATE TABLE IF NOT EXISTS t_region (
    id serial Primary key, 
    name varchar(50) not null
);

CREATE TABLE IF NOT EXISTS t_location (
    id serial Primary key, 
    name varchar(50) not null,
    link varchar(255), 
    regionId int not null, 
    FOREIGN KEY (regionId) REFERENCES t_region(id) 
);

CREATE TABLE IF NOT EXISTS t_product (
    id serial Primary key, 
    width int, 
    height int,
    display_type varchar(50) not null, 
    locationId int not null, 
    price int,
    FOREIGN KEY (locationId) REFERENCES t_location(id)
);

CREATE TABLE IF NOT EXISTS t_order (
    id serial Primary key, 
    status varchar(50) not null,
    orderTime Date not null, 
    deadline Date not null, 
    startDate Date not null, 
    endDate Date not null, 
    productId int not null, 
    clientId int not null, 
    managerId int not null, 
    FOREIGN KEY (productId) REFERENCES t_product(id),
    FOREIGN KEY (clientId) REFERENCES t_client(id),
    FOREIGN KEY (managerId) REFERENCES t_manager(id)
);