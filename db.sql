-- create database if not exists

CREATE DATABASE moji_chat;

-- use the database
\c dbname moji_chat;

-- create gender type
CREATE TYPE Gender AS ENUM ('Girl', 'Boy', 'Unknown');

-- create user status
CREATE TYPE UserStatus AS ENUM ('Normal', 'Ban', 'NotActivated')

-- create user table if not exists
CREATE TABLE sys_user (
  uid BIGSERIAL PRIMARY KEY,
  user_code VARCHAR(50) NOT NULL,
  avatar VARCHAR(512) NOT NULL,
  password VARCHAR(512) NOT NULL,
  nickname VARCHAR(50) NOT NULL,
  gender Gender NOT NULL DEFAULT 'Boy',
  birth DATE,
  phone_num VARCHAR(15),
  email VARCHAR(255),
  status UserStatus NOT NULL DEFAULT 'NotActivated',
  create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE sys_user IS 'user table';

COMMENT ON COLUMN sys_user.uid IS 'system generate of user id';
COMMENT ON COLUMN sys_user.user_code IS 'user defined id';
COMMENT ON COLUMN sys_user.avatar IS 'user portrait';
COMMENT ON COLUMN sys_user.password IS 'encrypted password';
COMMENT ON COLUMN sys_user.nickname IS 'display name';
COMMENT ON COLUMN sys_user.gender IS 'user gender, Boy is default';
COMMENT ON COLUMN sys_user.birth IS 'user birthday';
COMMENT ON COLUMN sys_user.phone_num IS 'the current bound mobile phone number';
COMMENT ON COLUMN sys_user.email IS 'the current bound email';
COMMENT ON COLUMN sys_user.status IS 'current status of the account. default is not activated in the register';
COMMENT ON COLUMN sys_user.create_at IS 'time of create';

insert into sys_user(user_code, avatar, password, nickname, gender, birth, email, status, create_at)
values ('100000', 'https://0.gravatar.com/avatar/bb935245d9ce4b3868849718160085f8bc5e62c0231177a24cca99bc99e27fe9?size=256', '123456', 'Clover You', 'Boy', '2002-04-13', 'cloveryou02@gmail.com', 'Normal', '2024-07-31 09:58:19')
