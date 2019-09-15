
-- +migrate Up
CREATE TABLE user (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(16) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE tag (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(32) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE quest (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(32) NOT NULL,
  capacity BIGINT NOT NULL,
  member_descriotion TEXT,
  quest_description TEXT,
  period DATETIME NOT NULL,
  reward TEXT,
  isFinished BIT NOT NULL,
  created_date DATETIME NOT NULL,
  updated_date DATETIME NOT NULL,
  producer_id BIGINT NOT NULL, 
  tag_id BIGINT, 
  PRIMARY KEY (id),
  FOREIGN KEY (producer_id) REFERENCES user(id),
  FOREIGN KEY (tag_id) REFERENCES tag(id)
);

CREATE TABLE application (
  id BIGINT NOT NULL AUTO_INCREMENT,
  status BIGINT NOT NULL,
  applicant_id BIGINT NOT NULL,
  quest_id BIGINT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (applicant_id) REFERENCES user(id),
  FOREIGN KEY (quest_id) REFERENCES quest(id)
);

CREATE TABLE message (
  id BIGINT NOT NULL AUTO_INCREMENT,
  body TEXT NOT NULL,
  send_time DATETIME NOT NULL,
  sender_id BIGINT NOT NULL,
  quest_id BIGINT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (sender_id) REFERENCES user(id),
  FOREIGN KEY (quest_id) REFERENCES quest(id)
);

-- +migrate Down
DROP TABLE IF EXISTS message;
DROP TABLE IF EXISTS application;
DROP TABLE IF EXISTS quest;
DROP TABLE IF EXISTS tag;
DROP TABLE IF EXISTS user;