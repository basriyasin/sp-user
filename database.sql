/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */


-- Considering the simplicity of the project architecture,
-- authentication, profile, and user phone number will be
-- put into one table to maintain simplicity.
-- and for larger projects, it is recommended to separate 
-- authentication, user profile, and user phone into 
-- separate tables to increase flexibility in handling 
-- future scenarios, such as login with multiple 
-- authorization methods, a single user with multiple 
-- phone numbers, etc.
create table if not exists profile (
    id          serial,
    name        varchar(60) not null,
    password    varchar(60) not null,
    phone       varchar(14) unique not null,
    login_count integer default 0,
    created_at  timestamp default current_timestamp,
    updated_at  timestamp
);
create index on profile (id);
create index on profile (phone,password);
create index on profile (created_at);