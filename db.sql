-- create database if not exists

CREATE DATABASE IF NOT EXISTS moji_chat;

-- use the database
\c dbname moji_chat;

-- create gender type
CREATE TYPE Gender AS ENUM ('Girl', 'Boy', 'Unknown');

-- create user status
CREATE TYPE UserStatus AS ENUM ('Normal', 'Ban', 'NotActivated')

-- create user table if not exists
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_tables WHERE schemaname = 'public' AND tablename = 'sys_user') THEN
    CREATE TABLE sys_user (
      uid BIGINT PRIMARY KEY,
      user_code VARCHAR(50) NOT NULL,
      avatar VARCHAR(512) NOT NULL,
      password VARCHAR(50) NOT NULL,
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
    
  END IF;
END $$;

---墨迹ID， 


CREATE TABLE weather (
  city VARCHAR(80),
  temp_lo INT,
  temp_hi INT,
  prcp REAL,
  date DATE
);

CREATE TABLE cities (
  name VARCHAR(80),
  location point
);

INSERT INTO weather VALUES ('San Francisco', 46, 50, 0.25, '1994-11-27');

INSERT INTO cities VALUES ('San Francisco', '(-194.0, 53.0)');

INSERT INTO weather (date, city, temp_hi, temp_lo)
    VALUES('1994-11-29', 'Hayward', 54, 37);

SELECT * from weather;

SELECT avg(temp_lo) from weather;

SELECT city, count(*) FILTER (WHERE temp_lo < 45), max(temp_lo) from weather GROUP BY city;
