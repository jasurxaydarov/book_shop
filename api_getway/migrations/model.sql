CREATE DATABASE book_shop;


-- CREATE TABLE users (
--     user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     username VARCHAR(50) UNIQUE NOT NULL,
--     email VARCHAR(100) UNIQUE NOT NULL,
--     password VARCHAR(255) NOT NULL,
--     full_name VARCHAR(100),
--     is_admin BOOLEAN DEFAULT FALSE,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP


-- );


CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    user_role varchar(24)   DEFAULT "viwer",
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP


);





ALTER TABLE users
DROP COLUMN user_role, -- Remove the is_admin column
ADD COLUMN user_role VARCHAR(24) DEFAULT 'user'; -- Add the user_role column with a default value