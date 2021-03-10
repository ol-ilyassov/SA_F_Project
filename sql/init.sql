CREATE TABLE articles(
  id serial not null unique, 
  authorid int not null,
  title varchar(100) not null, 
  content text not null, 
  created timestamp not null, 
  expires timestamp not null
);

CREATE TABLE users (
  id SERIAL NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  created timestamp NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

ALTER TABLE articles ADD CONSTRAINT fk_authorid 
FOREIGN KEY (authorid) REFERENCES users (id);

CREATE INDEX idx_articles_created ON articles(created);
