-- +migrate Up
-- MIGRATION EXAMPLE
CREATE TABLE log
(
    id       int auto_increment,
    user_id     varchar (100)  default null,
    code     varchar (30)  default null,
    log     varchar (1000)  default null,
    created_at     datetime  not null,
    CONSTRAINT link_pk PRIMARY KEY (id)
);
