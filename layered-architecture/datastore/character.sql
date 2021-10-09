DROP DATABASE IF EXISTS character;
CREATE DATABASE character;
USE  character;

CREATE TABLE characters(
  id int NOT NULL AUTO_INCRMENT,
  name varchar(50),
  level varchar(50),
  gender varchar(50),
  race varchar(50),
  class varchar(50),
  faction varchar(50),
  PRIMARY KEY(id)
);

INSERT INTO characters VALUES(1, 'Leeroy Jenkins', "70", "Male", "Humanoid", "Paladin", "Alliance")
INSERT INTO characters VALUES(2, 'Sylvanas Windrunner', "??(Boss)", "Female", "Forsaken", "Ranger", "Horde")