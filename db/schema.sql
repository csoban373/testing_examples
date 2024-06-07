DROP TABLE IF EXISTS people;
CREATE TABLE people(
    id                int primary key,
    name          varchar(50)    default '' not null
);