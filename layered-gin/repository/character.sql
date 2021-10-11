DROP DATABASE IF EXISTS db_characters;
CREATE DATABASE db_characters;
USE db_characters;

CREATE TABLE characters (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  title varchar(255) NOT NULL,
  gender varchar(100) NOT NULL,
  race varchar(100) NOT NULL,
  class varchar(100) NOT NULL,
  affiliation varchar(100) NOT NULL,
  status varchar(50) NOT NULL,
  PRIMARY KEY(id)
)