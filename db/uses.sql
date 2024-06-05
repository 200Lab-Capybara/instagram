CREATE TABLE IF NOT EXISTS users(
    id varchar(36) PRIMARY KEY ,
    email varchar(50) NOT NULL UNIQUE,
    password varchar(255) NOT NULL ,
    salt varchar(50) NOT NULL ,
    first_name nvarchar(50) NOT NULL ,
    last_name nvarchar(50) NOT NULL ,
    role enum('user', 'admin') default 'user',
    created_at timestamp,
    updated_at timestamp
)