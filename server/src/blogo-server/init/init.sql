DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS posts;

CREATE TABLE authors (
  id integer AUTO_INCREMENT,
  name varchar(255),
  mail varchar(255),
  password varchar(255),
  github varchar(255),
  primary key(id)
);

INSERT INTO authors (name, mail, password, github)
VALUES
  ('po3rin', 'abctail30@gmail.com', '1234567890', 'po3rin');

CREATE TABLE posts (
  id integer AUTO_INCREMENT,
  title varchar(255),
  content varchar(255),
  created_at datetime,
  updated_at datetime,
  primary key(id)
);

INSERT INTO posts (title, content, created_at)
VALUES
  ('test', 'test-content', now());