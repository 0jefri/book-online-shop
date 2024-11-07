create database book_online_shop;

CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(100),
    password VARCHAR(100)
);