DROP DATABASE IF EXISTS charactersdb;
CREATE DATABASE charactersdb;
USE charactersdb;

CREATE TABLE characters (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(50),
  level varchar(50),
  gender varchar(50),
  race varchar(50),
  class varchar(50),
  faction varchar(50),
  PRIMARY KEY(id)
);

INSERT INTO characters VALUES(1, 'Leeroy Jenkins', '70', 'Male', 'Humanoid', 'Paladin', 'Alliance');
INSERT INTO characters VALUES(2, 'Sylvanas Windrunner', '??(Boss)', 'Female', 'Forsaken', 'Ranger', 'Horde');