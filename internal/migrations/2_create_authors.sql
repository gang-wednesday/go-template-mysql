-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE  authors  (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username varchar(100) UNIQUE,
    email varchar(100) UNIQUE,
    name TEXT ,
    active BOOLEAN,
    author_address TEXT,
    last_login TIMESTAMP,
    last_password_change TIMESTAMP,
    token TEXT,
    role_id int,
    created_at TIMESTAMP,
	updated_at TIMESTAMP DEFAULT NOW(),
	deleted_at TIMESTAMP, 
    CONSTRAINT fk__role_users FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- +migrate Down
DROP TABLE authors;