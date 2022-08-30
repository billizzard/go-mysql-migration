-- +migrate Up
-- MIGRATION EXAMPLE
START TRANSACTION;

CREATE TABLE user
(
    id       int auto_increment,
    email     varchar (100)  default null,
    role     varchar (10)  not null,
    created_at     datetime  not null,
    uid     varchar(50) not null,
    CONSTRAINT link_pk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX user_uid_uindex on user (uid);
CREATE UNIQUE INDEX user_email_uindex on user (email);

COMMIT;
